package blockchain

import (
	"time"
)

type Block struct {
	index        int
	Timestamp    time.Time
	Transaction  Transaction
	proof        int
	previousHash string
}
