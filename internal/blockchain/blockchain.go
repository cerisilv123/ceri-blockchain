package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Blockchain represents a chain of blocks and a pool of current transactions.
type Blockchain struct {
	Chain               []Block       // The chain of blocks.
	CurrentTransactions []Transaction // Pool of current transactions.
}

// NewBlockchain creates and initializes a new Blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{
		Chain:               []Block{},
		CurrentTransactions: []Transaction{},
	}
	blockchain.AddBlock(1, "100") // Create the genesis block with default values.
	return blockchain
}

// hash computes the SHA-256 hash of a given block and returns it as a string.
func (b *Blockchain) hash(block Block) string {
	blockJSON, err := json.Marshal(block)
	if err != nil {
		return "" // Handle error better in production code.
	}

	hash := sha256.Sum256(blockJSON)
	hashString := fmt.Sprintf("%x", hash)
	fmt.Printf("%+v\n", "hash: "+hashString)

	return hashString
}

// AddTransaction adds a new transaction to the current pool of transactions.
// It returns the index of the block that will include this transaction.
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

// AddBlock creates a new block with the provided proof and previous hash, and appends it to the chain.
// If the previous hash is empty, it calculates the hash of the last block in the chain.
func (b *Blockchain) AddBlock(proof int, previousHash string) Block {
	if previousHash == "" {
		previousHash = b.hash(b.Chain[len(b.Chain)-1]) // Handle error if this returns nil in production code.
	}

	block := Block{
		Index:        len(b.Chain) + 1,
		Timestamp:    time.Now(),
		Transactions: b.CurrentTransactions,
		Proof:        proof,
		PreviousHash: previousHash,
	}

	b.CurrentTransactions = []Transaction{} // Emptying current transactions after adding them to a block.

	b.Chain = append(b.Chain, block)
	return block
}

// ProofOfWork performs a proof-of-work algorithm to find the valid proof for the next block.
// It takes the last proof as input and returns the new valid proof.
func (b *Blockchain) ProofOfWork(lastProof int) int {
	// - Find a number p' such that hash(pp') contains leading 5 zeroes, where p is the previous p'
	// - p is the previous proof, and p' is the new proof
	var proof int = 0

	for !b.ValidateProof(lastProof, proof) {
		proof += 1
	}

	return proof
}

// ValidateProof checks if the provided proof is valid by looking for a hash with 5 leading zeroes.
// It takes the last proof and the new proof as inputs and returns true if valid, false otherwise.
func (b *Blockchain) ValidateProof(lastProof int, proof int) bool {
	var guess string = strconv.Itoa(lastProof) + strconv.Itoa(proof)
	var guessHash [32]byte = sha256.Sum256([]byte(guess))
	var guessHashString string = fmt.Sprintf("%x", guessHash)

	return guessHashString[:5] == "00000"
}
