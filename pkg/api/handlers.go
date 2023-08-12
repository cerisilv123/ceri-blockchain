package api

import (
	"ceri-blockchain/internal/blockchain"
	"encoding/json"
	"fmt"
	"net/http"
)

// Custom handler type that includes a reference to the blockchain object which is instantiated in main.go.
type blockchainHandler func(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain)

// Handler for the "/mine" route (using POST)
func MineHandler(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "We'll mine a new Block")
}

// Handler for the "/transactions/new" route
func CreateTransactionHandler(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var transaction blockchain.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	index := bc.AddTransaction(transaction.Sender, transaction.Recipient, transaction.Amount)

	response := map[string]interface{}{
		"message": fmt.Sprintf("Transaction will be added to Block %d", index),
	}

	// Convert the response map to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client with HTTP status code 201 (Created)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// Handler for the "/chain/read" route
func ReadChainHandler(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"chain":  bc.Chain,
		"length": len(bc.Chain),
	}

	// Convert the response map to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client with HTTP status code 200 (OK)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
