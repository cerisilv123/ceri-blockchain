package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Blockchain represents a chain of blocks and a pool of current transactions & nodes on the network.
type Blockchain struct {
	Chain               []Block       // The chain of blocks.
	CurrentTransactions []Transaction // Pool of current transactions.
	Nodes               []Node        // An array of nodes currently running on the network.
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

// Checking whether a chain is a valid chain for consensus mechanism. Checks hashes and proofs.
func (b *Blockchain) ValidateChain(chain []Block) bool {

	prevBlock := b.Chain[0]
	currentIndex := 1

	for currentIndex < len(b.Chain) {
		currentBlock := b.Chain[currentIndex]

		// Checking block hashes are correct and have not been tampered with
		if currentBlock.PreviousHash != b.Hash(prevBlock) {
			return false
		}

		// Checking that proof is valid (proof is valid when prev and current proofs hashed contain leading 00000)
		if !b.ValidateProof(prevBlock.Proof, currentBlock.Proof) {
			return false
		}

		prevBlock = currentBlock
		currentIndex++
	}

	return true
}

// Consensus algorithm that resolves chain conflicts and shares the longest chain in the network
func (b *Blockchain) ResolveChainConflicts() bool {

	var newChain []Block

	// Only want to replace with a chain if it is greater in length
	maxLength := len(b.Chain)

	for i := 0; i < len(b.Nodes); i++ {
		nodeUrl := b.Nodes[i].URL

		// Make a GET request to get the chain for next node
		response, err := http.Get(nodeUrl)
		if err != nil {
			fmt.Println("Error:", err)
			// Throw error here
		}
		defer response.Body.Close()

		// Read the response body
		jsonBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			// Throw error here
		}

		var receivedData struct {
			Chain  []Block
			Length int
			Nodes  []Node
		}

		// Unmarshal the JSON response into the receivedChain slice
		err = json.Unmarshal(jsonBody, &receivedData)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			// Handle error here
		}

		length := receivedData.Length
		chain := receivedData.Chain

		if length > maxLength && b.ValidateChain(chain) {
			maxLength = length
			newChain = chain
		}
	}

	// Chain of current node is replaced with longer, validated chain.
	if newChain != nil {
		b.Chain = newChain
		return true
	}

	return false
}

// Registering a node on the network with URL, IP Address and Location (City)
func (b *Blockchain) RegisterNode(url string, ipAddress string, location string) {
	node := Node{
		URL:       url,
		IPAddress: ipAddress,
		Location:  location,
	}

	var nodeExists bool = false

	for _, existingNode := range b.Nodes {
		if existingNode.URL == node.URL && existingNode.IPAddress == node.IPAddress {
			nodeExists = true
			break
		}
	}

	if !nodeExists {
		b.Nodes = append(b.Nodes, node)
	} // Handle exception here
}

// Get last block in the blockchain
func (b *Blockchain) GetLastBlock() Block {
	lastBlock := b.Chain[len(b.Chain)-1]
	return lastBlock
}

// hash computes the SHA-256 hash of a given block and returns it as a string.
func (b *Blockchain) Hash(block Block) string {
	blockJSON, err := json.Marshal(block)
	if err != nil {
		return "" // Handle error better in production code.
	}

	hash := sha256.Sum256(blockJSON)
	hashString := fmt.Sprintf("%x", hash)

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
		previousHash = b.Hash(b.Chain[len(b.Chain)-1]) // Handle error if this returns nil in production code.
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
