package producer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	Nonce      uint64
	ChainId    *big.Int
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
	IsFaucet   bool

	Signer  *Signer
	ReqChan chan<- *TxSignRequest
	ResChan <-chan *TxSignResponse
}

func NewAccount(client *ethclient.Client, chainId *big.Int) (*Account, error) {
	pk, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pending nonce")
	}

	signer, err := NewSigner(client, pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signer")
	}

	reqChan := make(chan *TxSignRequest)
	resChan := signer.Run(reqChan)

	return &Account{
		Nonce:      nonce,
		Address:    addr,
		PrivateKey: pk,
		ChainId:    chainId,
		Signer:     signer,
		ReqChan:    reqChan,
		ResChan:    resChan,
	}, nil
}

func CreateFaucetAccount(client *ethclient.Client, privateKey string, chainId *big.Int) (*Account, error) {
	pks := strings.TrimPrefix(privateKey, "0x")
	pkBytes, _ := hex.DecodeString(pks)
	pk := loadPrivateKey(pkBytes)
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pending nonce")
	}

	signer, err := NewSigner(client, pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create signer")
	}

	reqChan := make(chan *TxSignRequest)
	resChan := signer.Run(reqChan)

	return &Account{
		IsFaucet:   true,
		Nonce:      nonce,
		Address:    addr,
		PrivateKey: pk,
		ChainId:    chainId,
		Signer:     signer,
		ReqChan:    reqChan,
		ResChan:    resChan,
	}, nil
}

func (a *Account) GetAndIncrementNonce() uint64 {
	now := a.Nonce
	a.Nonce += 1
	return now
}
