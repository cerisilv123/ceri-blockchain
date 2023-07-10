package blockchain

import (
	"fmt"
)

type Blockchain struct {
	Chain               []Block
	CurrentTransactions []Transaction
}

func (b *Blockchain) NewTransaction(sender string, recipient string, amount int) int {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	b.CurrentTransactions = append(b.CurrentTransactions, transaction)
	fmt.Printf("Current Transactions: %v\n", b.CurrentTransactions) // Remove this when not testing

	var newBlock int = len(b.Chain)
	return newBlock
}
