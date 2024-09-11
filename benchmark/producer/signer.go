package producer

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxSignRequest struct {
	Tx   *types.Transaction
	Data interface{}
}

type TxSignResponse struct {
	Err      error
	TxHash   common.Hash
	Request  *TxSignRequest
	SignedTx *types.Transaction
}

type Signer struct {
	signerAddress common.Address
	Auth          *bind.TransactOpts
	EvmClient     *ethclient.Client
}

func NewSigner(
	evmClient *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	chainId *big.Int,
) (*Signer, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainId)
	if err != nil {
		return nil, err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	auth.NoSend = true

	return &Signer{
		Auth:          auth,
		signerAddress: crypto.PubkeyToAddress(*publicKeyECDSA),
		EvmClient:     evmClient,
	}, nil
}

func (s *Signer) Run(requests <-chan *TxSignRequest) <-chan *TxSignResponse {
	responses := make(chan *TxSignResponse)

	// receive tx requests, sign & broadcast them.
	// Responses are sent once the tx is added to the pending tx pool.
	// To see result, use TransactionReceipt after tx has been included in a block.
	go func() {
		for {
			// wait for incoming request
			req := <-requests

			signedTx, err := s.Auth.Signer(s.signerAddress, req.Tx)
			if err != nil {
				err = errors.Wrap(err, "failed to sign transaction")
				responses <- &TxSignResponse{
					Request: req,
					Err:     err,
				}
			} else {
				println(s.signerAddress.Hex()+" -> tx signed: ", signedTx.Hash().String(), dumpTx(req.Tx))
				responses <- &TxSignResponse{
					Request:  req,
					SignedTx: signedTx,
					TxHash:   signedTx.Hash(),
					Err:      nil,
				}
			}
		}
	}()

	return responses
}

func (s *Signer) Address() common.Address {
	return s.signerAddress
}
