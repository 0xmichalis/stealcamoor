package stealcamoor

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"

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

	sc.emailCacheLock.Lock()
	defer sc.emailCacheLock.Unlock()

	if sc.emailCache[id] {
		// Skip if already notified (via the API listener)
		return
	}

	msg := fmt.Sprintf(`Newly minted memory id %d for %s!

	Steal at https://www.stealcam.com/memories/%d`, id, event.From, id)

	log.Print(msg)

	if err := sc.emailClient.Send([]string{sc.to}, msg); err != nil {
		log.Print("Failed to send email: ", err)
	} else {
		sc.emailCache[id] = true
	}
}
