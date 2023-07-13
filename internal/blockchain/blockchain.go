package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Blockchain struct {
	Chain               []Block
	CurrentTransactions []Transaction
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{
		Chain:               []Block{},
		CurrentTransactions: []Transaction{},
	}
	blockchain.AddBlock(1, "100")
	return blockchain
}

func (b *Blockchain) hash(block Block) string {
	blockJSON, err := json.Marshal(block)
	if err != nil {
		return "" // handle error better
	}

	hash := sha256.Sum256(blockJSON)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Printf("%+v\n", "hash: "+hashString)

	return hashString
}

func (b *Blockchain) AddTransaction(sender string, recipient string, amount int) int {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	b.CurrentTransactions = append(b.CurrentTransactions, transaction)

	var newBlock int = len(b.Chain)
	return newBlock
}

func (b *Blockchain) AddBlock(proof int, previousHash string) Block {

	if previousHash == "" {
		previousHash = b.hash(b.Chain[len(b.Chain)-1]) // Handle error if this returns nil
	}

	block := Block{
		Index:        len(b.Chain) + 1,
		Timestamp:    time.Now(),
		Transactions: b.CurrentTransactions,
		Proof:        proof,
		PreviousHash: previousHash,
	}

	b.CurrentTransactions = []Transaction{} // Emptying current transactions

	b.Chain = append(b.Chain, block)
	return block
}
