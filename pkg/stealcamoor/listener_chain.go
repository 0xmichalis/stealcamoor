package stealcamoor

import (
	"fmt"
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
			sc.handleStolenEvent(stolen)
		}
	}
}

func (sc *Stealcamoor) handleStolenEvent(event *abis.StealcamStolen) {
	id := event.Id.Uint64()

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

	msg := fmt.Sprintf(`Newly minted memory id %s for %s!

	Steal at https://www.stealcam.com/memories/%d`, id, event.From, id)

	log.Print(msg)

	if err := sc.emailClient.Send([]string{sc.to}, msg); err != nil {
		log.Print("Failed to send email:", err)
	} else {
		sc.emailCache[id] = true
	}
}
