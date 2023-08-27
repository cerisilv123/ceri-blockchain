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
func MineHandler(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain, nodeAddress string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Getting data from last block in the chain which will link to new block
	lastBlock := bc.GetLastBlock()
	lastProof := lastBlock.Proof
	newProof := bc.ProofOfWork(lastProof)

	// Adding a final transaction to show the end of the block
	bc.AddTransaction("0", nodeAddress, 1) // Sender(default to 0), recipient & amount

	// Getting previous hash and instantiated new Block with link to previous via hash and proof
	previousHash := bc.Hash(lastBlock)
	newBlock := bc.AddBlock(newProof, previousHash)

	// Creating JSON body
	response := map[string]interface{}{
		"message":      "New block created in the blockchain.",
		"index":        newBlock.Index,
		"previousHash": newBlock.PreviousHash,
		"proof":        newBlock.Proof,
		"transactions": newBlock.Transactions,
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
		"nodes":  bc.Nodes, // Remove this after testing
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

// Handler for the "/nodes/register" route
func RegisterNodeHandler(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var node blockchain.Node
	err := json.NewDecoder(r.Body).Decode(&node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bc.RegisterNode(node.URL, node.IPAddress, node.Location)

	response := map[string]interface{}{
		"message": "Node added to the ceri-blockchain network.",
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

// Handler for the "/nodes/resolve" route
func ResolveNodeHandler(w http.ResponseWriter, r *http.Request, bc *blockchain.Blockchain) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	chainSubstituted := bc.ResolveChainConflicts()

	var response map[string]interface{}

	if chainSubstituted {
		response = map[string]interface{}{
			"message": "Current chain was replaced by a new longer & Validated chain.",
		}
	} else {
		response = map[string]interface{}{
			"message": "Chain has not been replaced and is the valid chain.",
		}
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
