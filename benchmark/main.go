package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/berachain/beacon-kit/benchmark/producer"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	accountCount     = 5
	senderCount      = 4
	workloadPreBatch = accountCount
	txCount          = 100
)

type Sender struct {
	index     int
	generator producer.Generator
	client    *ethclient.Client
}

func (s *Sender) Send(txCount int) {
	ch := s.generator.GenerateTransfer()
	ctx := context.Background()
	cnt := 0
	for t := range ch {
		err := s.client.SendTransaction(ctx, t)
		if err != nil {
			log.Fatal("[", s.index, "] Failed to send transactions ", t.Hash().String(), " error: ", err.Error())
		} else {
			println("[", s.index, "] Sent transaction", t.Hash().String())
		}
		cnt++
		if cnt == txCount {
			break
		}
	}
}

var senders []*Sender

func main() {
	rpcUrl, _ := getParameters()

	senders = make([]*Sender, 0, senderCount)
	for i := range senderCount {
		client, err := ethclient.Dial(rpcUrl)
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}

		generator, err := producer.NewTransferGenerator(accountCount, "0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306", client)
		// generator, err = producer.NewErc20Generator(accountCount, "0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306", client)
		if err != nil {
			log.Fatalf("[%d]Failed to create the generator: %v", i, err.Error())
		}

		senders = append(senders, &Sender{i, generator, client})

		err = generator.WarmUp()
		if err != nil {
			log.Fatalf("[%d]Failed to warm up the generator: %v", i, err.Error())
		}

		tx := make([](*types.Transaction), 0, accountCount)
		c := generator.GenerateTransfer()

		for t := range c {
			tx = append(tx, t)
			if len(tx) == accountCount {
				break
			}
		}
		sendWorkload(client, tx)
	}

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	startBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the start block: %v", err)
	}

	swg := producer.NewSizedWaitGroup(senderCount)
	for i := range senders {
		swg.Add()
		go func(s *Sender) {
			defer swg.Done()
			s.Send(txCount)
		}(senders[i])
	}

	swg.Wait()

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
			log.Fatal("Failed to send transactions ", tx.Hash().String(), " error: ", err.Error())
		} else {
			println("Sent transaction", tx.Hash().String())
		}
	}
}
