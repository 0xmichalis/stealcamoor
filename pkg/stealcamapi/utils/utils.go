package utils

import (
	"github.com/0xmichalis/stealcamoor/pkg/stealcamapi"
)

func FilterUnmintedMemories(memories []stealcamapi.Memory) []stealcamapi.Memory {
	unminted := make([]stealcamapi.Memory, 0)
	for _, m := range memories {
		if m.Owner == nil {
			unminted = append(unminted, m)
		}
	}
	return unminted
}
