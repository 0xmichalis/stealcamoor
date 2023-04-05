package stealcamoor

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	noOpts = new(bind.CallOpts)
	zero   = big.NewInt(0)
)

func (sc *Stealcamoor) startChainListener() {
	log.Println("Starting on-chain listener...")

	headers := make(chan *types.Header)
	sub, err := sc.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to headers: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Got subscription error: %v", err)

		case header := <-headers:
			log.Printf("Processing block %d", header.Number.Uint64())

		}
	}
}
