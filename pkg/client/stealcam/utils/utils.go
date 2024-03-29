package utils

import (
	"github.com/0xmichalis/stealcamoor/pkg/client/stealcam"
)

func FilterUnmintedMemoryIDs(memories []stealcam.Memory) []uint64 {
	ids := make([]uint64, 0)
	for _, m := range memories {
		if m.Owner == nil {
			ids = append(ids, m.ID)
		}
	}
	return ids
}
