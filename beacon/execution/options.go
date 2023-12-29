// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package execution

import (
	"cosmossdk.io/log"
	"github.com/itsdevbear/bolaris/beacon/execution/engine"
	"github.com/itsdevbear/bolaris/types/config"
)

type Option func(*EngineNotifier) error

// WithLogger is an option to set the logger for the Eth1Client.
func WithBeaconConfig(beaconCfg *config.Beacon) Option {
	return func(s *EngineNotifier) error {
		s.beaconCfg = beaconCfg
		return nil
	}
}

// WithLogger is an option to set the logger for the Eth1Client.
func WithLogger(logger log.Logger) Option {
	return func(s *EngineNotifier) error {
		s.logger = logger
		return nil
	}
}

// WithForkChoiceStoreProvider is an option to set the ForkChoiceStoreProvider
// for the EngineNotifier.
func WithForkChoiceStoreProvider(fcsp forkchoiceStoreProvider) Option {
	return func(s *EngineNotifier) error {
		s.fcsp = fcsp
		return nil
	}
}

// WithEngineCaller is an option to set the Caller for the EngineNotifier.
func WithEngineCaller(ec engine.Caller) Option {
	return func(s *EngineNotifier) error {
		s.engine = ec
		return nil
	}
}
