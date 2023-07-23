package main

import (
	"ceri-blockchain/internal/blockchain"
	"fmt"
)

func main() {
	// Create an instance of the Blockchain struct
	blockchain := blockchain.NewBlockchain()

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
	blockchain.AddBlock(65478654785, "")

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

	sender4 := "David"
	recipient4 := "Hugh"
	amount4 := 123
	newBlockIndex4 := blockchain.AddTransaction(sender4, recipient4, amount4)

	fmt.Printf("New block index: %d\n", newBlockIndex)
	fmt.Printf("New block index: %d\n", newBlockIndex2)
	fmt.Printf("New block index: %d\n", newBlockIndex3)
	fmt.Printf("New block index: %d\n", newBlockIndex4)

	// Adding block
	blockchain.AddBlock(547895765986, "hfukrehfekrfherlkjd")

	// Print blocks
	fmt.Println("Blocks:")
	for _, block := range blockchain.Chain {
		fmt.Printf("%+v\n", block)
	}

	// Testing proof of work algorithm
	lastProof := 12345
	proof := blockchain.ProofOfWork(lastProof)
	fmt.Printf("Valid proof: %d\n", proof)
}
