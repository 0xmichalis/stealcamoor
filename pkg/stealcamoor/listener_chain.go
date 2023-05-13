package stealcamoor

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/event"

	"github.com/0xmichalis/stealcamoor/pkg/abis"
)

var (
	zero = big.NewInt(0)
)

func (sc *Stealcamoor) isStolenFromCreator(event *abis.StealcamStolen) bool {
	for _, creator := range sc.creators {
		// Needs to be a creator we care about. Also currently only interested
		// in getting the first mint but this may be expanded in the future to
		// be more configurable.
		if event.From.String() == creator.String() && event.Value.Cmp(zero) == 0 {
			return true
		}
	}
	return false
}

func (sc *Stealcamoor) startChainListener() {
	log.Print("Starting on-chain listener...")

	steals := make(chan *abis.StealcamStolen)
	sub := event.Resubscribe(2*time.Second, func(ctx context.Context) (event.Subscription, error) {
		sub, err := sc.stealcamContract.StealcamFilterer.WatchStolen(nil, steals)
		if err != nil {
			log.Fatalf("Failed to subscribe to Stolen events: %v", err)
		}
		return sub, nil
	})

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				log.Printf("Got subscription error: %v", err)
			} else {
				log.Print("Got a nil subsciption error?! Not sure why this can happen.")
				time.Sleep(10 * time.Second)
			}

		case stolen := <-steals:
			sc.handleStolenEvent(stolen)
		}
	}
}

func (sc *Stealcamoor) handleStolenEvent(event *abis.StealcamStolen) {
	id := event.Id.Uint64()
	log.Print("Handling Stolen event with id ", id)

	if !sc.isStolenFromCreator(event) {
		// Skip if not a mint
		return
	}

	sc.tryMint(event.From, []uint64{id})
}
