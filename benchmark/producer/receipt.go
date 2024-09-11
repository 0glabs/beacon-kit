package producer

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrBroadcastTimeout = errors.New("timed out waiting for tx to be committed to block")
	ErrTxFailed         = errors.New("transaction was committed but failed. likely an execution revert by contract code")
)

func WaitForTxReceipt(client *ethclient.Client, txHash common.Hash, timeout time.Duration) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error
	outOfTime := time.After(timeout)
	for {
		select {
		case <-outOfTime:
			err = ErrBroadcastTimeout
		default:
			receipt, err = client.TransactionReceipt(context.Background(), txHash)
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			if receipt.Status == 0 {
				err = ErrTxFailed
			}
		}
		break
	}
	return receipt, err
}
