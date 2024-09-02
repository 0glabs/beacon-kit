package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/berachain/beacon-kit/benchmark/producer"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	accountCount     = 20000
	workloadPreBatch = accountCount
)

var generator producer.Generator

func main() {
	rpcUrl, _ := getParameters()

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	generator, err = producer.NewGenerator(accountCount, "0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306", client)
	if err != nil {
		log.Fatalf("Failed to create the generator: %v", err)
	}

	err = generator.WarmUp()
	if err != nil {
		log.Fatalf("Failed to warm up the generator: %v", err)
	}
	// complete the transfer from facuet to new generated accounts
	sendWorkload(client, getWorkload(workloadPreBatch))
	// sleep to wait the initial transfer to be completed
	time.Sleep(20 * time.Second)

	startBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the start block: %v", err)
	}

	sendWorkload(client, getWorkload(workloadPreBatch))
	sendWorkload(client, getWorkload(workloadPreBatch))

	endBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the end block: %v", err)
	}

	startHeight := startBlock.Number()
	startTime := startBlock.Time()

	endHeight := endBlock.Number()
	endTime := endBlock.Time()

	if endTime <= startTime {
		log.Fatal("More transactions are needed")
	}

	totalTxCount := 0

	// Iterate over the blocks from startHeight to the endHeight
	for blockNumber := startHeight.Add(startHeight, big.NewInt(1)); blockNumber.Cmp(endHeight) <= 0; blockNumber.Add(blockNumber, big.NewInt(1)) {
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatalf("Failed to fetch block: %v", err)
		}

		totalTxCount += len(block.Transactions())
	}

	elapsedSeconds := endTime - startTime
	tps := float64(totalTxCount) / float64(elapsedSeconds)

	fmt.Printf("Total transactions counted in %d seconds is %d\n", elapsedSeconds, totalTxCount)
	fmt.Printf("The TPS of the chain is %.2f\n", tps)
}

func getParameters() (string, int) {
	// handle command line flags
	rpcUrl := flag.String("rpc-url", "http://127.0.0.1:8545", "RPC url of the chain")
	count := flag.Int("count", 10000, "The number of transactions to be sent")
	flag.Parse()

	if *count > 1000000 {
		log.Fatal("Too many transactions to be generated and sent")
	}

	return *rpcUrl, *count
}

func sendWorkload(client *ethclient.Client, workload [](*types.Transaction)) {
	for _, tx := range workload {
		err := client.SendTransaction(context.Background(), tx)
		if err != nil {
			log.Fatal("Failed to send transactions")
		}
	}
}

func getWorkload(n int) [](*types.Transaction) {
	// load all to-be-sent transactions
	return generator.GenerateGeneralTransfer(n)
}
