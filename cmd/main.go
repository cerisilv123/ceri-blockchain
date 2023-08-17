package main

import (
	"ceri-blockchain/internal/blockchain"
	"ceri-blockchain/pkg/api"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var NodeAddress uuid.UUID

func main() {

	// Unique address for node
	NodeAddress = uuid.New()
	NodeAddressString := NodeAddress.String()

	// Create an instance of the Blockchain struct
	blockchain := blockchain.NewBlockchain()

	// Setting up server with the custom handlers, passing the "bc" instance as an argument
	http.HandleFunc("/mine", func(w http.ResponseWriter, r *http.Request) {
		api.MineHandler(w, r, blockchain, NodeAddressString)
	})

	http.HandleFunc("/transactions/new", func(w http.ResponseWriter, r *http.Request) {
		api.CreateTransactionHandler(w, r, blockchain)
	})

	http.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		api.ReadChainHandler(w, r, blockchain)
	})

	// Starting server on port 8080
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
