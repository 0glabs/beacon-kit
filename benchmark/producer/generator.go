package producer

import (
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	defaultTransferGasLimit      = uint64(22000)
	defaultErc20TransferGasLimit = uint64(310000)
	initialTransferVal           = 1000000000000000000
	defaultTransferVal           = 100000000000
)

type task struct {
	fromAccount *Account
	toAccout    *Account
	value       int64
}

type Generator interface {
	WarmUp() error
	GenerateTransfer(numTransfers int) []*types.Transaction
}
