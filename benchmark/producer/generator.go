package producer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	defaultTransferGasLimit      = uint64(22000)
	defaultErc20TransferGasLimit = uint64(310000)
	initialTransferVal           = 1e16
	defaultTransferVal           = 1e6
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
