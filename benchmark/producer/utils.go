package producer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
	"golang.org/x/exp/rand"
)

// func generateEthSecp256k1PrivateKeyByUint32(n uint32) *ecdsa.PrivateKey {
// 	// privKeyInt := new(big.Int).SetUint64(uint64(n))
// 	curve := crypto.S256()
// 	byteSlice := make([]byte, 32)
// 	binary.LittleEndian.PutUint32(byteSlice, n)
// 	privateKey, err := ecdsa.GenerateKey(curve, bytes.NewReader(byteSlice))
// 	if err != nil {
// 		panic(errors.Wrapf(err, "failed to generate private key, n: %d, slice : %v", n, byteSlice))
// 	}

// 	return privateKey
// }

func toChecksumAddress(address string) string {
	address = strings.ToLower(strings.TrimPrefix(address, "0x"))

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(address))
	hash := hasher.Sum(nil)
	hashHex := hex.EncodeToString(hash)

	checksummedAddress := "0x"
	for i, c := range address {
		if c >= '0' && c <= '9' {
			checksummedAddress += string(c)
		} else {
			if hashHex[i] >= '8' {
				checksummedAddress += strings.ToUpper(string(c))
			} else {
				checksummedAddress += string(c)
			}
		}
	}

	return checksummedAddress
}

func loadPrivateKey(privateKeyBytes []byte) *ecdsa.PrivateKey {
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func shuffle(slice []uint32) []uint32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

type TransactionOutput struct {
	Nonce    uint64   `json:"nonce"`
	To       string   `json:"to"`
	Value    *big.Int `json:"value"`
	GasLimit uint64   `json:"gas_limit"`
	GasPrice *big.Int `json:"gas_price"`
	Data     []byte   `json:"data"`
}

func dumpTx(tx *types.Transaction) string {
	txOutput := TransactionOutput{
		Nonce:    tx.Nonce(),
		To:       tx.To().Hex(),
		Value:    tx.Value(),
		GasLimit: tx.Gas(),
		GasPrice: tx.GasPrice(),
		Data:     tx.Data(),
	}

	// 将结构体序列化为 JSON
	jsonData, err := json.MarshalIndent(txOutput, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal transaction to JSON: %v", err)
	}

	return string(jsonData)
}

func toBigInt(amount any) *big.Int {
	if amount == nil {
		return big.NewInt(0)
	}
	var val *big.Int
	switch amount.(type) {
	case int:
		val = big.NewInt(int64(amount.(int)))
	case int32:
		val = big.NewInt(int64(amount.(int32)))
	case int64:
		val = big.NewInt(amount.(int64))
	case string:
		var ok bool
		val, ok = new(big.Int).SetString(amount.(string), 0)
		if !ok {
			panic(fmt.Sprintf("invalid amount string: %s", amount.(string)))
		}
	case *big.Int:
		val = amount.(*big.Int)
	case float64:
		val = decimal.NewFromFloat(amount.(float64)).BigInt()
	default:
		panic(fmt.Sprintf("invalid amount type: %T", amount))
	}

	return val
}
