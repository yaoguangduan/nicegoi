package util

import (
	"fmt"
	"sync/atomic"
)

var seq = atomic.Uint64{}

func AllocEID() string {
	id := fmt.Sprintf("E%d", seq.Add(1))

	return id
}
