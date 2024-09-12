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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type erc20GeneratorImlp struct {
	client     *ethclient.Client
	chainId    *big.Int
	signer     types.Signer
	accountMap *AccountMap
	taskPool   chan *task
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

	am, err := NewAccountMap(ethClient, numAccounts, faucetPrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	return &erc20GeneratorImlp{
		client:     ethClient,
		chainId:    chainID,
		signer:     types.NewEIP155Signer(chainID),
		accountMap: am,
		poolSize:   numAccounts,
		taskPool:   make(chan *task, 100),
		txPool:     make(chan *types.Transaction, numAccounts),
	}, nil
}

func (g *erc20GeneratorImlp) WarmUp() error {
	// make transfer from faucet account to other accounts
	taskList := make([]*task, 0, g.accountMap.total)

	auth, instance, err := deployContractErc20Test(g.client, g.accountMap.GetFaucetAccount())
	if err != nil {
		return errors.Wrapf(err, "failed to deploy contract")
	}
	auth.NoSend = true
	g.auth = auth
	g.instance = instance

	for i := 0; i < int(g.accountMap.total); i++ {
		taskList = append(taskList, &task{
			fromAccount: g.accountMap.faucetAcct,
			toAccout:    g.accountMap.GetAccount(uint32(i)),
			value:       toBigInt(initialTransferVal),
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

func (g *erc20GeneratorImlp) generateTransaction(t *task) (*types.Transaction, error) {
	nonce := t.fromAccount.GetAndIncrementNonce()

	g.auth.Nonce = big.NewInt(int64(nonce))

	tx, err := g.instance.Transfer(t.fromAccount.Signer.Auth, t.toAccout.Address, t.value)
	if err != nil {
		return nil, err
	}

	t.fromAccount.ReqChan <- &TxSignRequest{
		Tx: tx,
	}

	res := <-t.fromAccount.ResChan

	return res.SignedTx, nil
}

func (g *erc20GeneratorImlp) GenerateTransfer() <-chan *types.Transaction {
	go func() {
		for {
			for i := uint32(0); i < g.accountMap.total; i++ {
				g.taskPool <- &task{
					fromAccount: g.accountMap.GetFaucetAccount(),
					toAccout:    g.accountMap.GetAccount(i),
					value:       toBigInt(defaultTransferVal),
				}
			}
		}
	}()

	go func() {
		for {
			t := <-g.taskPool
			tx, err := g.generateTransaction(t)
			if err != nil {
				log.Fatal("generate transaction error: ", err.Error())
			}
			g.txPool <- tx
		}
	}()

	return g.txPool
}
