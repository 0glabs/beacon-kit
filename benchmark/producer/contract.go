package producer

import (
	"context"
	"math/big"

	"github.com/berachain/beacon-kit/benchmark/producer/contract"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployContractErc20Test(evmClient *ethclient.Client, account *Account) (*bind.TransactOpts, *contract.Erc20test, error) {
	privateKey := loadPrivateKey(account.PrivateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, account.ChainId)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to create transactor with private key %s", account.PrivateKey)
	}

	auth.Nonce = big.NewInt(int64(account.GetAndIncrementNonce()))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(defaultErc20TransferGasLimit)

	gasPrice, err := evmClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get suggest gas price")
	}

	auth.GasPrice = gasPrice

	_, _, bondContract, err := contract.DeployErc20test(auth, evmClient)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to deploy contract")
	}

	// receipt, err := bind.WaitMined(context.Background(), evmClient, tx)
	// if err != nil {
	// 	return nil, nil, errors.Wrapf(err, "wait for contract deployment, %s", tx.Hash().Hex())
	// }

	// if receipt.Status != 1 {
	// 	return nil, nil, errors.Wrapf(err, "contract deployment failed, %s", tx.Hash().Hex())
	// }

	return auth, bondContract, nil
}
