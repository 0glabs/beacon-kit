package producer

// 1, prepare accounts
// 2, make transfer from faucet account to other accounts
// 3, make transfer between accounts

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type transferGeneratorImlp struct {
	client     *ethclient.Client
	chainId    *big.Int
	signer     types.Signer
	accountMap *AccountMap
	taskPool   chan *task
	txPool     chan *types.Transaction
	poolSize   uint32
}

func NewTransferGenerator(numAccounts uint32, faucetPrivateKey string, ethClient *ethclient.Client) (Generator, error) {
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	am, err := NewAccountMap(ethClient, numAccounts, faucetPrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	return &transferGeneratorImlp{
		client:     ethClient,
		chainId:    chainID,
		signer:     types.NewEIP155Signer(chainID),
		accountMap: am,
		poolSize:   numAccounts,
		taskPool:   make(chan *task, 100),
		txPool:     make(chan *types.Transaction, numAccounts),
	}, nil
}

func (g *transferGeneratorImlp) WarmUp() error {
	// make transfer from faucet account to other accounts
	taskList := make([]*task, 0, g.accountMap.total)
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

func (g *transferGeneratorImlp) generateTransaction(t *task) (*types.Transaction, error) {
	ctx := context.Background()

	nonce := t.fromAccount.GetAndIncrementNonce()

	gasLimit := defaultTransferGasLimit
	gasPrice, err := g.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// tx := types.NewTransaction(nonce, t.toAccout.Address, t.value, gasLimit, big.NewInt(0), nil)
	tx := types.NewTransaction(nonce, t.toAccout.Address, t.value, gasLimit, gasPrice, nil)
	t.fromAccount.ReqChan <- &TxSignRequest{
		Nonce: nonce,
		Tx:    tx,
	}

	res := <-t.fromAccount.ResChan

	return res.SignedTx, nil
}

func (g *transferGeneratorImlp) GenerateTransfer() <-chan *types.Transaction {
	go func() {
		for {
			acctIdxList := make([]uint32, 0, g.accountMap.total)

			for i := 0; i < int(g.accountMap.total); i++ {
				acctIdxList = append(acctIdxList, uint32(i))
			}

			acctIdxList = shuffle(acctIdxList)

			usedFrom := make(map[uint32]bool)
			usedTo := make(map[uint32]bool)

			for i := 0; i < len(acctIdxList); i++ {
				for j := 0; j < len(acctIdxList); j++ {
					if acctIdxList[i] != acctIdxList[j] && !usedFrom[acctIdxList[i]] && !usedTo[acctIdxList[j]] {
						usedFrom[acctIdxList[i]] = true
						usedTo[acctIdxList[j]] = true

						g.taskPool <- &task{
							fromAccount: g.accountMap.GetAccount(acctIdxList[i]),
							toAccout:    g.accountMap.GetAccount(acctIdxList[j]),
							value:       toBigInt(defaultTransferVal),
						}
						break
					}
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
