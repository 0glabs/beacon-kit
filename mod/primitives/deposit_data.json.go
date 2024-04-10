// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package primitives

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*depositDataJSONMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (d DepositData) MarshalJSON() ([]byte, error) {
	type DepositData struct {
		Pubkey      hexutil.Bytes `json:"pubkey" ssz-max:"48"`
		Credentials hexutil.Bytes `json:"credentials" ssz-size:"32"`
		Amount      Gwei          `json:"amount"`
		Signature   hexutil.Bytes `json:"signature" ssz-max:"96"`
	}
	var enc DepositData
	enc.Pubkey = d.Pubkey[:]
	enc.Credentials = d.Credentials[:]
	enc.Amount = d.Amount
	enc.Signature = d.Signature[:]
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (d *DepositData) UnmarshalJSON(input []byte) error {
	type DepositData struct {
		Pubkey      *hexutil.Bytes `json:"pubkey" ssz-max:"48"`
		Credentials *hexutil.Bytes `json:"credentials" ssz-size:"32"`
		Amount      *Gwei          `json:"amount"`
		Signature   *hexutil.Bytes `json:"signature" ssz-max:"96"`
	}
	var dec DepositData
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Pubkey != nil {
		if len(*dec.Pubkey) != len(d.Pubkey) {
			return errors.New("field 'pubkey' has wrong length, need 48 items")
		}
		copy(d.Pubkey[:], *dec.Pubkey)
	}
	if dec.Credentials != nil {
		if len(*dec.Credentials) != len(d.Credentials) {
			return errors.New("field 'credentials' has wrong length, need 32 items")
		}
		copy(d.Credentials[:], *dec.Credentials)
	}
	if dec.Amount != nil {
		d.Amount = *dec.Amount
	}
	if dec.Signature != nil {
		if len(*dec.Signature) != len(d.Signature) {
			return errors.New("field 'signature' has wrong length, need 96 items")
		}
		copy(d.Signature[:], *dec.Signature)
	}
	return nil
}
