package stealcamoor

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/0xmichalis/stealcamoor/pkg/abis"
)

var (
	noOpts = new(bind.CallOpts)
	zero   = big.NewInt(0)
)

func (sc *Stealcamoor) isStolenFromCreator(event *abis.StealcamStolen) bool {
	for _, creator := range sc.creators {
		if event.From.String() == creator.String() && event.Value.Cmp(zero) == 0 {
			return true
		}
	}
	return false
}

func (sc *Stealcamoor) startChainListener() {
	log.Println("Starting on-chain listener...")

	steals := make(chan *abis.StealcamStolen)
	sub, err := sc.stealcamContract.StealcamFilterer.WatchStolen(nil, steals)
	if err != nil {
		log.Fatalf("Failed to subscribe to Stolen events: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Got subscription error: %v", err)

		case stolen := <-steals:
			if sc.isStolenFromCreator(stolen) {

			}
		}
	}
}
