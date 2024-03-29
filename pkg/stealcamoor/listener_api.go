package stealcamoor

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"

	stealcamutils "github.com/0xmichalis/stealcamoor/pkg/client/stealcam/utils"
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
		if err := sc.checkCreator(creator); err != nil {
			log.Printf("Failed to check memories for creator %s: %v", creator, err)
		}
	}
}

func (sc *Stealcamoor) checkCreator(creator common.Address) error {
	// Get all memories for a creator
	memories, err := sc.apiClient.GetMemories(creator)
	if err != nil {
		return fmt.Errorf("cannot get memories: %w", err)
	}

	// Filter out all but unminted memories
	unmintedIDs := stealcamutils.FilterUnmintedMemoryIDs(memories)
	if len(unmintedIDs) == 0 {
		log.Print("No unminted memories for ", creator)
		return nil
	}

	sc.tryMint(creator, unmintedIDs)

	return nil
}
