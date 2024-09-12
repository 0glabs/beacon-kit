package producer

import (
	"math/big"

	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/geth-primitives/pkg/ethclient"
)

type AccountMap struct {
	total      uint32
	accounts   []*Account
	faucetAcct *Account
}

func NewAccountMap(client *ethclient.Client, total uint32, faucetPrivateKey string, chainId *big.Int) (*AccountMap, error) {
	am := &AccountMap{
		total:    total,
		accounts: make([]*Account, 0, total),
	}
	var err error
	am.faucetAcct, err = CreateFaucetAccount(client, faucetPrivateKey, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create faucet account")
	}

	for i := uint32(0); i < total; i++ {
		newAccnt, err := NewAccount(client, chainId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create account")
		}
		am.accounts = append(am.accounts, newAccnt)
	}

	return am, nil
}

func (am AccountMap) GetAccount(index uint32) *Account {
	if index < am.total {
		return am.accounts[index]
	}
	return nil
}

func (am AccountMap) GetAccountCount() uint32 {
	return am.total
}

func (am AccountMap) GetFaucetAccount() *Account {
	return am.faucetAcct
}
