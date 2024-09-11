package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	defaultTransferGasLimit      = uint64(22000)
	defaultErc20TransferGasLimit = uint64(310000)
	initialTransferVal           = 100000000000000
	defaultTransferVal           = 10000000000
)

type task struct {
	fromAccount *Account
	toAccout    *Account
	value       *big.Int
}

type Generator interface {
	WarmUp() error
	GenerateTransfer() <-chan *types.Transaction
}

type DeployedErc20 struct {
	Address string
}
