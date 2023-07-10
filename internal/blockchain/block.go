package blockchain

import (
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Transaction  Transaction
	Proof        int
	PreviousHash string
}
