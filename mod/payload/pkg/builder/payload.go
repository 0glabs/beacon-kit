// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package builder

import (
	"context"
	"time"

	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

// RequestPayload builds a payload for the given slot and
// returns the payload ID.
func (pb *PayloadBuilder[
	BeaconStateT, ExecutionPayloadT, ExecutionPayloadHeaderT,
]) RequestPayloadAsync(
	ctx context.Context,
	st BeaconStateT,
	slot math.Slot,
	timestamp uint64,
	parentBlockRoot primitives.Root,
	headEth1BlockHash common.ExecutionHash,
	finalEth1BlockHash common.ExecutionHash,
	forceUpdate bool,
) (*engineprimitives.PayloadID, error) {
	if !pb.Enabled() {
		return nil, ErrPayloadBuilderDisabled
	}

	// Attempt to see if we previously fired off a payload built for
	// this particular slot and parent block root.
	prevBuiltPayloadID, found := pb.pc.Get(slot, parentBlockRoot)

	// If we find a payload and we are not in force mode, we can skip
	// building the payload.
	if found && !forceUpdate {
		pb.logger.Warn(
			"skipping payload build and forkchoice update; already built for slot",
			"for_slot", slot,
			"parent_block_root", parentBlockRoot,
		)
		return &prevBuiltPayloadID, nil
	}

	// Assemble the payload attributes.
	attrs, err := pb.getPayloadAttribute(st, slot, timestamp, parentBlockRoot)
	if err != nil {
		return nil, errors.Newf("%w error when getting payload attributes", err)
	}

	// Submit the forkchoice update to the execution client.
	var payloadID *engineprimitives.PayloadID
	payloadID, _, err = pb.ee.NotifyForkchoiceUpdate(
		ctx, &engineprimitives.ForkchoiceUpdateRequest{
			State: &engineprimitives.ForkchoiceStateV1{
				HeadBlockHash:      headEth1BlockHash,  // some 2954 of some description
				SafeBlockHash:      finalEth1BlockHash, // 2953
				FinalizedBlockHash: finalEth1BlockHash, // 2953
			},
			PayloadAttributes: attrs,
			ForkVersion:       pb.chainSpec.ActiveForkVersionForSlot(slot),
		},
	)

	// If we got a nil response when we were expecting a payload ID and we
	// have something to fall back to we can use the previous payload ID.
	if payloadID == nil && prevBuiltPayloadID != (engineprimitives.PayloadID{}) {
		args := []any{
			"for_slot", slot,
			"parent_block_root", parentBlockRoot,
			"prev_payload_id", prevBuiltPayloadID,
			"force_update", forceUpdate,
		}

		if err != nil {
			args = append(args, "error", err)
		}

		pb.logger.Warn(
			"payload build produced nil payload ID - "+
				"falling back to previously built payload", args,
		)

		return &prevBuiltPayloadID, nil
	} else if err != nil {
		return nil, err
	} else if payloadID == nil {
		return nil, ErrNilPayloadID
	}

	pb.logger.Info(
		"bob the builder; can we forkchoice update it?;"+
			" bob the builder; yes we can 🚧",
		"head_eth1_hash",
		headEth1BlockHash,
		"for_slot",
		slot,
		"parent_block_root",
		parentBlockRoot,
		"payload_id",
		payloadID,
	)

	pb.pc.Set(slot, parentBlockRoot, *payloadID)
	return payloadID, nil
}

// RequestPayload request a payload for the given slot and
// blocks until the payload is delivered.
func (pb *PayloadBuilder[
	BeaconStateT, ExecutionPayloadT, ExecutionPayloadHeaderT,
]) RequestPayloadSync(
	ctx context.Context,
	st BeaconStateT,
	slot math.Slot,
	timestamp uint64,
	parentBlockRoot primitives.Root,
	parentEth1Hash common.ExecutionHash,
	finalBlockHash common.ExecutionHash,
	forceUpdate bool,
) (engineprimitives.BuiltExecutionPayloadEnv[ExecutionPayloadT], error) {
	if !pb.Enabled() {
		return nil, ErrPayloadBuilderDisabled
	}
	// Build the payload and wait for the execution client to
	// return the payload ID.
	payloadID, err := pb.RequestPayloadAsync(
		ctx,
		st,
		slot,
		timestamp,
		parentBlockRoot,
		parentEth1Hash,
		finalBlockHash,
		forceUpdate,
	)
	if err != nil {
		return nil, err
	} else if payloadID == nil {
		return nil, ErrNilPayloadID
	}

	// Wait for the payload to be delivered to the execution client.
	pb.logger.Info(
		"waiting for local payload to be delivered to execution client",
		"for_slot", slot, "timeout", pb.cfg.PayloadTimeout.String(),
	)
	select {
	case <-time.After(pb.cfg.PayloadTimeout):
		// We want to trigger delivery of the payload to the execution client
		// before the timestamp expires.
		break
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Get the payload from the execution client.
	return pb.ee.GetPayload(
		ctx,
		&engineprimitives.GetPayloadRequest{
			PayloadID:   *payloadID,
			ForkVersion: pb.chainSpec.ActiveForkVersionForSlot(slot),
		},
	)
}

// RetrieveOrBuildPayload attempts to pull a previously built payload
// by reading a payloadID from the builder's cache. If it fails to
// retrieve a payload, it will build a new payload and wait for the
// execution client to return the payload.
func (pb *PayloadBuilder[
	BeaconStateT, ExecutionPayloadT, ExecutionPayloadHeaderT,
]) RetrievePayload(
	ctx context.Context,
	slot math.Slot,
	parentBlockRoot primitives.Root,
) (engineprimitives.BuiltExecutionPayloadEnv[ExecutionPayloadT], error) {
	if !pb.Enabled() {
		return nil, ErrPayloadBuilderDisabled
	}

	// Attempt to see if we previously fired off a payload built for
	// this particular slot and parent block root.
	payloadID, found := pb.pc.Get(slot, parentBlockRoot)
	if !found {
		return nil, ErrPayloadIDNotFound
	}

	envelope, err := pb.ee.GetPayload(
		ctx,
		&engineprimitives.GetPayloadRequest{
			PayloadID:   payloadID,
			ForkVersion: pb.chainSpec.ActiveForkVersionForSlot(slot),
		},
	)
	if err != nil {
		return nil, err
	} else if envelope == nil {
		return nil, ErrNilPayloadEnvelope
	}

	overrideBuilder := envelope.ShouldOverrideBuilder()
	args := []any{
		"for_slot", slot,
		"override_builder", overrideBuilder,
	}

	payload := envelope.GetExecutionPayload()
	if !payload.IsNil() {
		args = append(args,
			"payload_block_hash", payload.GetBlockHash(),
			"parent_hash", payload.GetParentHash(),
		)
	}

	blobsBundle := envelope.GetBlobsBundle()
	if blobsBundle != nil {
		args = append(args, "num_blobs", len(blobsBundle.GetBlobs()))
	}

	pb.logger.Info("payload retrieved from local builder 🏗️ ", args...)

	// If the payload was built by a different builder, something is
	// wrong the EL<>CL setup.
	if payload.GetFeeRecipient() != pb.cfg.SuggestedFeeRecipient {
		pb.logger.Warn(
			"payload fee recipient does not match suggested fee recipient - "+
				"please check both your CL and EL configuration",
			"payload_fee_recipient", payload.GetFeeRecipient(),
			"suggested_fee_recipient", pb.cfg.SuggestedFeeRecipient,
		)
	}
	return envelope, err
}

// RequestPayload builds a payload for the given slot and
// returns the payload ID.
//
// TODO: This should be moved onto a "sync service"
// of some kind.
func (pb *PayloadBuilder[
	BeaconStateT, ExecutionPayloadT, ExecutionPayloadHeaderT,
]) SendForceHeadFCU(
	ctx context.Context,
	st BeaconStateT,
	slot math.Slot,
) error {
	lph, err := st.GetLatestExecutionPayloadHeader()
	if err != nil {
		return err
	}

	pb.logger.Info(
		"sending startup forkchoice update to execution client 🚀 ",
		"head_eth1_hash", lph.GetBlockHash(),
		"safe_eth1_hash", lph.GetParentHash(),
		"finalized_eth1_hash", lph.GetParentHash(),
		"for_slot", slot,
	)

	// Submit the forkchoice update to the execution client.
	_, _, err = pb.ee.NotifyForkchoiceUpdate(
		ctx, &engineprimitives.ForkchoiceUpdateRequest{
			State: &engineprimitives.ForkchoiceStateV1{
				HeadBlockHash:      lph.GetBlockHash(),
				SafeBlockHash:      lph.GetParentHash(),
				FinalizedBlockHash: lph.GetParentHash(),
			},
			PayloadAttributes: nil,
			ForkVersion:       pb.chainSpec.ActiveForkVersionForSlot(slot),
		},
	)
	return err
}
