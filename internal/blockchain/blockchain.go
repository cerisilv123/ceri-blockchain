package blockchain

import (
	"time"
)

type Blockchain struct {
	Chain               []Block
	CurrentTransactions []Transaction
}

// NewBlockchain is a factory function that creates a new instance of the Blockchain struct.
func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{
		Chain:               []Block{},
		CurrentTransactions: []Transaction{},
	}
	blockchain.AddBlock(1, "100")
	return blockchain
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
