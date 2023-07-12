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
	newBlockIndex := blockchain.AddTransaction(sender, recipient, amount)

	sender2 := "Dave"
	recipient2 := "Thomas"
	amount2 := 54
	newBlockIndex2 := blockchain.AddTransaction(sender2, recipient2, amount2)

	// Print transactions
	fmt.Println("Transactions:")
	for _, transaction := range blockchain.CurrentTransactions {
		fmt.Printf("%+v\n", transaction)
	}

	// Adding block
	blockchain.AddBlock(65478654785, "eee7e87987eeeee78978ee")

	sender3 := "William"
	recipient3 := "Jones"
	amount3 := 123
	newBlockIndex3 := blockchain.AddTransaction(sender3, recipient3, amount3)

	// Print transactions
	fmt.Println("Transactions:")
	for _, transaction := range blockchain.CurrentTransactions {
		fmt.Printf("%+v\n", transaction)
	}

	// Adding block
	blockchain.AddBlock(787098700473, "zzzzyuyuiyiouuiyzzzzyuyizyz")

	fmt.Printf("New block index: %d\n", newBlockIndex)
	fmt.Printf("New block index: %d\n", newBlockIndex2)
	fmt.Printf("New block index: %d\n", newBlockIndex3)

	// Print blocks
	fmt.Println("Blocks:")
	for _, block := range blockchain.Chain {
		fmt.Printf("%+v\n", block)
	}
}
