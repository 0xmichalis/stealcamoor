package stealcamoor

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func (sc *Stealcamoor) sendEmail(msgFmt string, ids []uint64, creator common.Address) {
	sc.emailCacheLock.Lock()
	defer sc.emailCacheLock.Unlock()

	for _, id := range ids {
		if sc.emailCache[id] {
			// Skip if already notified (via the API listener)
			continue
		}

		msg := fmt.Sprintf(msgFmt, id, creator, id)

		log.Print(msg)

		if err := sc.emailClient.Send([]string{sc.to}, msg); err != nil {
			log.Print("Failed to send email: ", err)
		} else {
			sc.emailCache[id] = true
		}
	}
}
