package producer

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	Index      uint32
	Nonce      uint64
	ChainId    *big.Int
	Address    string
	Checksum   string
	PrivateKey []byte
	IsFaucet   bool
}

func NewAccount(index uint32, ChainId *big.Int) Account {
	pk, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	return Account{
		Index:      index,
		Nonce:      0,
		Address:    addr,
		Checksum:   toChecksumAddress(addr),
		PrivateKey: crypto.FromECDSA(pk),
		ChainId:    ChainId,
	}
}

func CreateFaucetAccount(privateKey string, ChainId *big.Int) *Account {
	pks := strings.TrimPrefix(privateKey, "0x")
	pkBytes, _ := hex.DecodeString(pks)
	pk := loadPrivateKey(pkBytes)
	addr := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	return &Account{
		IsFaucet:   true,
		Index:      0,
		Nonce:      0,
		Address:    addr,
		Checksum:   toChecksumAddress(addr),
		PrivateKey: crypto.FromECDSA(pk),
		ChainId:    ChainId,
	}
}

func (a *Account) GetAndIncrementNonce() uint64 {
	now := a.Nonce
	a.Nonce += 1
	return now
}
