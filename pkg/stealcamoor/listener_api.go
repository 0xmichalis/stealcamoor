package stealcamoor

import (
	"fmt"
	"log"
	"time"

	stealcamutils "github.com/0xmichalis/stealcamoor/pkg/stealcamapi/utils"
)

func (sc *Stealcamoor) startApiListener() {
	log.Println("Starting API listener...")

	sc.checkCreators()

	for {
		select {
		case <-time.Tick(sc.apiRequestInterval):
			log.Println("Checking creators...")
			sc.checkCreators()
		}
	}
}

func (sc *Stealcamoor) checkCreators() {
	for _, creator := range sc.creators {
		// Get all memories for a creator
		memories, err := sc.apiClient.GetMemories(creator)
		if err != nil {
			log.Println(err)
			continue
		}

		// Filter out all but unminted memories
		unminted := stealcamutils.FilterUnmintedMemories(memories)
		if len(unminted) == 0 {
			log.Printf("No unminted memories for %s\n", creator)
			continue
		}

		// Notify on unminted memories
		for _, m := range unminted {
			if sc.emailCache[m.ID] {
				// Skip if found in the cache
				continue
			}

			msg := fmt.Sprintf(`Unminted memory id %s for %s!

			Mint at https://www.stealcam.com/memories/%d`, m.ID, creator, m.ID)

			log.Print(msg)

			if err := sc.emailClient.Send([]string{sc.to}, msg); err != nil {
				log.Print("Failed to send email:", err)
			} else {
				sc.emailCache[m.ID] = true
			}
		}
	}
}
