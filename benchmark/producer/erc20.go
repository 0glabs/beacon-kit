package producer

// 1, prepare accounts
// 2, make transfer from faucet account to other accounts
// 3, make transfer between accounts

import (
	"context"
	"log"
	"math/big"

	"github.com/berachain/beacon-kit/benchmark/producer/contract"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type erc20Task struct {
	fromAccount *Account
	toAccout    *Account
	value       int64
}

type erc20GeneratorImlp struct {
	client     *ethclient.Client
	chainId    *big.Int
	signer     types.Signer
	accountMap *AccountMap
	txPool     chan *types.Transaction
	poolSize   uint32
	auth       *bind.TransactOpts
	instance   *contract.Erc20test
}

func NewErc20Generator(numAccounts uint32, faucetPrivateKey string, ethClient *ethclient.Client) (Generator, error) {
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &erc20GeneratorImlp{
		client:     ethClient,
		chainId:    chainID,
		signer:     types.NewEIP155Signer(chainID),
		accountMap: NewAccountMap(numAccounts, faucetPrivateKey, chainID),
		poolSize:   numAccounts,
		txPool:     make(chan *types.Transaction, numAccounts),
	}, nil
}

func (g *erc20GeneratorImlp) WarmUp() error {
	// make transfer from faucet account to other accounts
	taskList := make([]*erc20Task, 0, g.accountMap.total)

	ctx := context.Background()
	faucetAcct := g.accountMap.GetFaucetAccount()
	{
		nonce, err := g.client.PendingNonceAt(ctx, common.HexToAddress(faucetAcct.Address))
		if err != nil {
			return errors.Wrap(err, "failed to get pending nonce")
		}
		faucetAcct.Nonce = nonce
	}

	for i := 0; i < int(g.accountMap.total); i++ {
		thisAccount := g.accountMap.GetAccount(uint32(i))
		nonce, err := g.client.PendingNonceAt(ctx, common.HexToAddress(thisAccount.Address))
		if err != nil {
			return errors.Wrap(err, "failed to get pending nonce")
		}
		thisAccount.Nonce = nonce
	}

	auth, instance, err := deployContractErc20Test(g.client, g.accountMap.GetFaucetAccount())
	if err != nil {
		return errors.Wrapf(err, "failed to deploy contract")
	}
	auth.NoSend = true
	g.auth = auth
	g.instance = instance

	for i := 0; i < int(g.accountMap.total); i++ {
		taskList = append(taskList, &erc20Task{
			fromAccount: faucetAcct,
			toAccout:    g.accountMap.GetAccount(uint32(i)),
			value:       initialTransferVal,
		})
	}

	for i := range taskList {
		tx, err := g.generateTransaction(taskList[i])
		if err != nil {
			return err
		}
		g.txPool <- tx
	}
	return nil
}

func (g *erc20GeneratorImlp) generateTransaction(t *erc20Task) (*types.Transaction, error) {
	nonce := t.fromAccount.GetAndIncrementNonce()

	g.auth.Nonce = big.NewInt(int64(nonce))

	tx, err := g.instance.Transfer(g.auth, common.HexToAddress(t.toAccout.Address), big.NewInt(t.value))
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, g.signer, loadPrivateKey(t.fromAccount.PrivateKey))
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (g *erc20GeneratorImlp) GenerateTransfer(numTransfers int) []*types.Transaction {
	go func() {
		pairs := make([][2]*Account, 0, g.accountMap.total)

		for i := uint32(0); i < g.accountMap.total; i++ {
			pairs = append(pairs, [2]*Account{g.accountMap.GetFaucetAccount(), g.accountMap.GetAccount(i)})
		}

		g.generateTransfer(pairs, defaultTransferVal)
	}()

	ret := make([]*types.Transaction, 0, numTransfers)

	for tx := range g.txPool {
		ret = append(ret, tx)
		if len(ret) == numTransfers {
			break
		}
	}

	return ret
}

func (g *erc20GeneratorImlp) generateTransfer(paired [][2]*Account, value int64) {
	taskList := make([]*erc20Task, 0, len(paired))
	for i := 0; i < len(paired); i++ {
		taskList = append(taskList, &erc20Task{
			fromAccount: paired[i][0],
			toAccout:    paired[i][1],
			value:       value,
		})
	}

	for i := range taskList {
		tx, err := g.generateTransaction(taskList[i])
		if err != nil {
			log.Fatalf("Failed to generate transactions: %v", err)
			return
		}
		g.txPool <- tx
	}
}
