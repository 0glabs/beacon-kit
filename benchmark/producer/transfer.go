package producer

// 1, prepare accounts
// 2, make transfer from faucet account to other accounts
// 3, make transfer between accounts

import (
	"context"
	"math/big"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	defaultTransferGasLimit = uint64(21000)
	initialTransferVal      = 1000000000000000000
	defaultTransferVal      = 100000000000
)

type task struct {
	fromAccount *Account
	toAccout    *Account
	value       int64
}

type Generator interface {
	WarmUp() error
	GenerateGeneralTransfer(numTransfers int) []*types.Transaction
}

type generatorImlp struct {
	client     *ethclient.Client
	chainId    *big.Int
	signer     types.Signer
	accountMap *AccountMap
	txPool     chan *types.Transaction
	poolSize   uint32
}

func NewGenerator(numAccounts uint32, faucetPrivateKey string, ethClient *ethclient.Client) (Generator, error) {
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return &generatorImlp{
		client:     ethClient,
		chainId:    chainID,
		signer:     types.NewEIP155Signer(chainID),
		accountMap: NewAccountMap(numAccounts, faucetPrivateKey),
		poolSize:   numAccounts,
		txPool:     make(chan *types.Transaction, numAccounts),
	}, nil
}

func (g *generatorImlp) WarmUp() error {
	// make transfer from faucet account to other accounts
	taskList := make([]*task, 0, g.accountMap.total)

	ctx := context.Background()
	faucetAcct := g.accountMap.GetFaucetAccount()
	{
		nonce, err := g.client.PendingNonceAt(ctx, common.HexToAddress(faucetAcct.Address))
		if err != nil {
			return err
		}
		faucetAcct.Nonce = nonce
	}

	for i := 0; i < int(g.accountMap.total); i++ {
		thisAccount := g.accountMap.GetAccount(uint32(i))
		nonce, err := g.client.PendingNonceAt(ctx, common.HexToAddress(thisAccount.Address))
		if err != nil {
			return err
		}
		thisAccount.Nonce = nonce
	}

	for i := 0; i < int(g.accountMap.total); i++ {
		taskList = append(taskList, &task{
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

func (g *generatorImlp) generateTransaction(t *task) (*types.Transaction, error) {
	ctx := context.Background()

	nonce := t.fromAccount.GetAndIncrementNonce()

	gasLimit := defaultTransferGasLimit
	gasPrice, err := g.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(t.toAccout.Address), big.NewInt(t.value), gasLimit, gasPrice, nil)

	signedTx, err := types.SignTx(tx, g.signer, loadPrivateKey(t.fromAccount.PrivateKey))
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (g *generatorImlp) GenerateGeneralTransfer(numTransfers int) []*types.Transaction {
	go func() {
		acctIdxList := make([]uint32, 0, g.accountMap.total)

		for i := 0; i < int(g.accountMap.total); i++ {
			acctIdxList = append(acctIdxList, uint32(i))
		}

		acctIdxList = shuffle(acctIdxList)

		pairs := make([][2]uint32, 0, g.accountMap.total)
		usedFrom := make(map[uint32]bool)
		usedTo := make(map[uint32]bool)

		for i := 0; i < len(acctIdxList); i++ {
			for j := 0; j < len(acctIdxList); j++ {
				if acctIdxList[i] != acctIdxList[j] && !usedFrom[acctIdxList[i]] && !usedTo[acctIdxList[j]] {
					pairs = append(pairs, [2]uint32{acctIdxList[i], acctIdxList[j]})
					usedFrom[acctIdxList[i]] = true
					usedTo[acctIdxList[j]] = true
					break
				}
			}
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

func (g *generatorImlp) generateTransfer(paired [][2]uint32, value int64) {
	taskList := make([]*task, 0, len(paired))
	for i := 0; i < len(paired); i++ {
		taskList = append(taskList, &task{
			fromAccount: g.accountMap.GetAccount(paired[i][0]),
			toAccout:    g.accountMap.GetAccount(paired[i][1]),
			value:       value,
		})
	}

	swg := NewSizedWaitGroup(runtime.NumCPU())
	for i := range taskList {
		swg.Add()
		go func(t *task) {
			defer swg.Done()
			tx, err := g.generateTransaction(t)
			if err != nil {
				return
			}
			g.txPool <- tx
		}(taskList[i])
	}

	swg.Wait()
}
