package stealcamoor

import (
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
		memories, err := sc.apiClient.GetMemories(creator)
		if err != nil {
			log.Println(err)
			continue
		}
		unminted := stealcamutils.FilterUnmintedMemories(memories)
		if len(unminted) == 0 {
			log.Printf("No unminted memories for %s\n", creator)
			continue
		}
		for _, id := range unminted {
			log.Println("Unminted memory id %s for %s\n", id, creator)
		}
	}
}
