package main

import (
	"ceri-blockchain/internal/blockchain"
	"fmt"
)

func main() {
	// Create an instance of the Blockchain struct
	chain := []blockchain.Block{}
	transactions := []blockchain.Transaction{}
	blockchain := blockchain.Blockchain{
		Chain:               chain,
		CurrentTransactions: transactions,
	}

	// Call the newTransaction function
	sender := "Alice"
	recipient := "Bob"
	amount := 23

	newBlock := blockchain.NewTransaction(sender, recipient, amount)

	fmt.Printf("New block index: %d\n", newBlock)
}
