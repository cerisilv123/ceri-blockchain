# ceri-blockchain
A simple Proof of Work Blockchain built in Go. This project showcases the fundamental concepts of blockchain technology through a functional blockchain system. It includes features such as proof-of-work, mining, creating transactions, resolving conflicts, and managing nodes in a decentralized network using a consensus mechanism.

## What is ceri-blockchain?

At its core, a blockchain is a decentralized, immutable, and transparent ledger that allows secure transactions without the need for intermediaries. The ceri-blockchain project was a small project I undertook to provide myself with a hands-on understanding of these concepts.

## Features

- **Proof-of-Work Mechanism:** The system employs a proof-of-work algorithm to validate and add new blocks to the blockchain. This algorithm requires miners to solve a computational puzzle, ensuring that adding a block is resource-intensive and time-consuming.

- **Mining Functionality:** The project enables mining, which is the process of solving the proof-of-work puzzle to add new blocks. Mining serves the dual purpose of validating transactions and creating new blocks.

- **Transaction Management:** Ceri Blockchain allows users to create transactions by specifying the sender, recipient, and amount. These transactions are then included in the blockchain after being validated through mining.

- **Decentralized Network:** The project supports the creation of a decentralized network of nodes. Nodes can register themselves in the network and participate in the consensus process to ensure that all nodes have the same version of the blockchain.

- **Conflict Resolution:** In a decentralized network, conflicts may arise when different nodes add blocks simultaneously. Ceri Blockchain implements a consensus algorithm that resolves conflicts by choosing the longest and validated chain.

## Getting Started

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/ceri-blockchain.git
2. Running the program:
   - You can either run the program by building and running the Docker image provided in the repo or by running it using 'go run .'
   - When running the program on to a node you would need to notify other nodes on the network by calling the endpoint /nodes/register (the code could be cloned and this process could be improved!)

### Models, Services, Handlers & Persistence

In the context of GoLang directory structure, the terms "models," "services," "handlers," and "persistence" are commonly used to categorize different components of an application. Here's a breakdown of their typical roles:

1. Models:
    - The "models" folder typically contains the data structures and logic that represent the domain entities or objects in your application.
    - These models often correspond to the data stored in a database or exchanged through APIs.
    - They define the structure, behavior, and relationships of the data within the application.
    - Models encapsulate the business rules and data manipulation logic related to the entities they represent.
2. Services:
    - The "services" folder typically contains the business logic or application-specific services of your application.
    - Services encapsulate the core functionality and behavior of your application.
    - They orchestrate the operations performed on the models and handle the main business use cases.
    - Services might include operations like data validation, business rules enforcement, coordination between different components, and interaction with external systems.
3. Handlers:
    - The "handlers" folder typically contains the code responsible for handling incoming requests and generating responses in a web application or API.
    - Handlers act as the entry point for HTTP requests and define the specific routes and actions to be performed.
    - They parse incoming requests, extract relevant data, and invoke the appropriate services or business logic to fulfill the request.
    - Handlers often handle request validation, error handling, and transformation of data between request/response formats (e.g., JSON).
4. Persistence:
    - The "persistence" folder (sometimes referred to as "repositories" or "data access") typically contains the code related to data storage and retrieval.
    - This includes database operations, data access logic, and any interactions with external storage systems.
    - Persistence components handle reading from and writing to the database, executing queries, and managing data transactions.
    - They abstract away the specific database implementation details and provide a consistent interface for interacting with the data layer.

---

### Important Concepts

1. **Blockchain:** A blockchain is a distributed and decentralized digital ledger that records a series of transactions across multiple computers or nodes. It is designed to be secure, transparent, and tamper-resistant. Each transaction in a blockchain is grouped into a block, and blocks are linked together chronologically, forming a chain of blocks.
2. **Block**: A block is a data structure that contains a set of transactions in a blockchain. It typically consists of several components, including a block header and a list of transactions. The block header contains metadata such as a timestamp, a reference to the previous block's hash, and a nonce (a number used in the proof-of-work consensus algorithm). Each block is cryptographically linked to the previous block, creating a chain of blocks.
3. **Transaction**: A transaction represents an exchange of data or value between participants in a blockchain network. It can involve transferring cryptocurrency tokens, recording ownership information, or executing smart contract functions. Transactions typically include information about the sender, recipient, the amount transferred, and any additional data required by the specific blockchain protocol.
4. **Hash**: A hash refers to the output of a cryptographic hash function. In the context of blockchain, hashes are used to uniquely identify data and ensure its integrity. A hash function takes an input (such as a block or a transaction) and produces a fixed-size string of characters, which is unique to that input. Even a small change in the input data will result in a completely different hash value. Hash functions in blockchains provide security by making it computationally infeasible to reverse-engineer the original input data from the hash value.
5. **Proof**: In blockchain, the term "proof" is often used in the context of consensus mechanisms. It refers to a mathematical or computational process that demonstrates the validity or authenticity of a block or transaction. Different consensus mechanisms employ different proofs to achieve agreement among network participants. For example, proof of work and proof of stake are two commonly used consensus mechanisms, which I'll explain in more detail below.
6. **Node**: In a blockchain network, a node refers to any computer or device that participates in the network by maintaining a copy of the blockchain and validating transactions. Nodes can be categorized into full nodes and lightweight nodes. Full nodes store a complete copy of the blockchain and participate in the consensus process by validating and relaying transactions. Lightweight nodes, on the other hand, rely on full nodes for transaction validation and only store a subset of the blockchain.
7. **Proof of Work:** Proof of Work (PoW) is a consensus mechanism used in some blockchains, such as Bitcoin. It requires network participants, known as miners, to solve a computationally intensive puzzle. Miners compete to find a solution to the puzzle, which involves repeatedly hashing the block's data with a nonce until a specific condition is met (e.g., the hash value starts with a certain number of zeros). Once a miner finds a valid solution, they can propose the next block to be added to the blockchain and are rewarded with newly minted cryptocurrency. PoW ensures the security and integrity of the blockchain by making it expensive and time-consuming to alter past blocks, as it would require re-mining all subsequent blocks.

---

### 1) **Blockchain (Class)**

A Blockchain class is created, wherein the constructor establishes:

1. an initial empty list to serve as the storage for the blockchain
2. another list designated for ‘current’ transaction storage ie transactions that have not been added to a chain. 
3. another array to store neighbouring Nodes in the network. 

```jsx
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
```

The Blockchain class assumes the role of managing the chain by handling transaction storage

and providing auxiliary methods for incorporating (**mining**) new blocks into the chain. 

### 2) **Block (Data Structure)**

Every Block in the blockchain possesses an index, a timestamp (measured in Unix time), a

collection of transactions, a proof (details to follow), and the hash of the previous Block.

```
package blockchain

import (
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	Proof        int
	PreviousHash string
}
```

**Previous hash:** 

The **`previous_hash`** field in a block plays a crucial role in ensuring the immutability and integrity

of the blockchain. However, it alone does not guarantee immutability.

The **`previous_hash`** field contains the hash of the previous block in the blockchain. By including

the **`previous_hash`** in each block, the blocks are cryptographically linked together, forming a

chain. This linkage creates a dependency where altering the data of any block in the chain would

result in a change in its hash, which in turn would affect the subsequent blocks' **`previous_hash`**

values.

**Proof:** 

In a blockchain, a proof is a crucial concept that ensures the security and integrity of the network

by making it computationally expensive to create new blocks. The Proof of Work (PoW) algorithm

is used to find a specific value, known as the proof, that meets certain criteria. This proof is

added to a new block and proves that a significant amount of computational work has been done

before adding the block to the blockchain. We’ll cover this in more detail later when we dive in to

implementing proof of work!

### 3) **Transaction (Data Structure)**

For this basic blockchain we will keep it simple on the transaction side of things. We won’t be

implementing public/private key cryptography in this project as we will focus more on creating a

chain, proof of work and a consensus mechanism. Therefore the simple transaction struct will be

in the following format: 

```go
package blockchain

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}
```

### 4) Node (Data Structure)

It’s important we store data about other nodes on the network. This is to ensure nodes can share blockchain data between each-other peer to peer and resolve any chain conflicts. 

```go
package blockchain

type Node struct {
	URL       string
	IPAddress string
	Location  string
}
```

### 5) Implementing the Blockchain Class methods

**AddTransaction:**

So, we are going to add the first class method to the Blockchain Class which is to add a transaction to the blockchain. The **`AddTransaction`** function allows users to add new transactions to the blockchain's current transaction pool. These transactions will be later included in a block when the block is mined (more on this when we cover the mine function). The function returns the index of the next block where the transaction will be included. This mechanism ensures that the transaction is ready to be processed and added to the blockchain's history.

```go
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
```

**Hash:** 

Ahhh onto the good stuff! We want a function that can hash a Block. This is important because

we need to store a hash of the previous Block in each Block data structure. A hash refers to a

function that takes an input (or message) of arbitrary size and produces a fixed-size output,

often represented as a sequence of characters. The primary characteristics of a good hash

function are that it should be fast to compute, deterministic (same input always produces the

same output), and should ideally be irreversible (it's computationally infeasible to retrieve the

original input from the hash)

```go
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
```

Let me break it down for you so you can see how it works: 

1. The function is a method of the **`Blockchain`** struct, and it takes a single parameter **`block`**, which is an instance of the **`Block`** struct for which we want to calculate the hash.
2. Inside the function, the provided **`block`** is converted into a JSON-encoded byte slice using the **`json.Marshal`** function. This converts the block's data into a format that can be easily processed and hashed.
3. The **`json.Marshal`** function can potentially return an error if the encoding process fails. The code includes an error check to handle such cases. If an error occurs, the function returns an empty string. In a production scenario, better error handling mechanisms could be implemented.
4. The **`sha256.Sum256`** function is then used to calculate the SHA-256 hash of the JSON-encoded block. This function takes a byte slice as input and returns a fixed-size hash as an array of bytes.
5. The calculated hash is in the form of a byte array. To represent the hash in a more human-readable and compact way, the code uses the **`fmt.Sprintf`** function to convert the byte array to a hexadecimal string.
6. Finally, the function returns the hexadecimal string representation of the calculated hash.

**AddBlock:** 

Now we have our hashing method, lets add a method that allows us to add a Block to the Blockchain array. 

```go
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
```

Let me break this down for you:

1. The function takes two parameters:
    - **`proof`**: An integer representing the proof of work associated with the new block.
    - **`previousHash`**: A string representing the hash of the previous block in the blockchain. If this is the first block (genesis block) or if the previous hash is not provided, the function calculates the hash of the last block in the chain.
2. The function begins by checking if the **`previousHash`** is empty. If it's empty, it means we're adding the first block or creating a new block that links to the latest block in the chain. In this case, the hash of the last block in the chain is calculated using the **`Hash`** function, and that hash is set as the **`previousHash`** for the new block.
3. The function then creates a new instance of the **`Block`** struct with the following attributes:
    - **`Index`**: The index of the new block, which is one more than the index of the last block in the chain.
    - **`Timestamp`**: The current time when the block is being added.
    - **`Transactions`**: The transactions that were present in the current transaction pool at the time of adding the block.
    - **`Proof`**: The proof of work associated with the new block.
    - **`PreviousHash`**: The hash of the previous block, which was either provided or calculated earlier.
4. After creating the new block, the function empties the pool of current transactions (**`b.CurrentTransactions`**). This is because transactions that are added to a block are removed from the pool to prepare for new transactions.
5. The newly created block is then appended to the blockchain's chain using the **`append`** function.
6. Finally, the function returns the newly added block.

**ProofOfWork:** 

Here is a super simple Proof of Work Algorithm used in our Blockchain. In a PoW system, miners

(or nodes) perform computational work to find a value (proof) that, when combined with certain

data, produces a hash with specific properties. Here's how the PoW algorithm is implemented in

the given code:

```go
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
```

1. The function takes the **`lastProof`** as input, which is the proof of the previous block. This is used as a reference for the current proof generation.
2. Inside the function, a variable **`proof`** is initialized to 0. The goal of the PoW algorithm is to find a value for **`proof`** that, when combined with the **`lastProof`**, produces a hash that meets certain criteria (in this case, a hash with leading zeros).
3. The loop continues until a valid proof is found. The condition for validity is determined by the **`ValidateProof`** function. The **`ValidateProof`** function checks if the hash of the combination of **`lastProof`** and **`proof`** has the required number of leading zeros (in this case, 5 zeros).
4. If the current **`proof`** does not meet the criteria, the loop increments the **`proof`** and continues to the next iteration.
5. When a valid **`proof`** is found (i.e., the condition in the **`ValidateProof`** function is satisfied), the loop exits, and the valid **`proof`** is returned.

**ValidateProof:**

The **`ValidateProof`** function calculates a hash of the concatenation of **`lastProof`** and **`proof`**,

and then checks if the first five characters of the hash are zeros. This ensures that the hash

meets the criteria specified in the PoW algorithm.

```go
// ValidateProof checks if the provided proof is valid by looking for a hash with 5 leading zeroes.
// It takes the last proof and the new proof as inputs and returns true if valid, false otherwise.
func (b *Blockchain) ValidateProof(lastProof int, proof int) bool {
	var guess string = strconv.Itoa(lastProof) + strconv.Itoa(proof)
	var guessHash [32]byte = sha256.Sum256([]byte(guess))
	var guessHashString string = fmt.Sprintf("%x", guessHash)

	return guessHashString[:5] == "00000"
}
```

**ValidateChain:**

```go
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

```

This function (**`ValidateChain`**) checks whether a given chain is valid according to the consensus mechanism. It performs the following steps:

- It takes a chain as input and initializes **`prevBlock`** to the first block in the blockchain.
- It sets **`currentIndex`** to 1, indicating that we're starting validation from the second block.
- The function loops through each block in the chain:
    - It checks if the **`PreviousHash`** of the current block matches the hash of the previous block. If not, it means the previous hash has been tampered with, and the function returns **`false`**.
    - It then checks if the proof of the current block is valid using the **`ValidateProof`** function. If not, it means the proof of work is incorrect, and the function returns **`false`**.
    - The **`prevBlock`** is updated to the current block, and the loop moves to the next block.
- If all blocks in the chain pass the validation checks, the function returns **`true`**, indicating that the chain is valid.

**ResolveChainConflicts:**

```go
// Consensus algorithm that resolves chain conflicts and shares the longest chain in the network
func (b *Blockchain) ResolveChainConflicts() bool {
	var newChain []Block
	maxLength := len(b.Chain)

	for i := 0; i < len(b.Nodes); i++ {
		// ... (HTTP request logic)
		// ... (JSON response handling)

		length := receivedData.Length
		chain := receivedData.Chain

		if length > maxLength && b.ValidateChain(chain) {
			maxLength = length
			newChain = chain
		}
	}

	if newChain != nil {
		b.Chain = newChain
		return true
	}

	return false
}

```

This function (**`ResolveChainConflicts`**) implements a consensus algorithm that resolves conflicts in the blockchain network by adopting the longest valid chain. Here's how it works:

- It initializes **`newChain`** to an empty slice of blocks and **`maxLength`** to the length of the current chain.
- The function iterates through each node in the network:
    - It sends an HTTP GET request to the node's **`/chain`** endpoint to retrieve its blockchain data.
    - It parses the JSON response to extract the received chain and its length.
    - If the received chain's length is greater than the current **`maxLength`** and the chain is valid according to **`ValidateChain`**, the function updates **`maxLength`** and sets **`newChain`** to the received chain.
- After iterating through all nodes, if a longer valid chain (**`newChain`**) was found, the function updates the current blockchain's chain with the new one and returns **`true`**.
- If no longer valid chain was found, the function returns **`false`**.

**RegisterNode:**

```go
// Registering a node on the network with URL, IP Address, and Location (City)
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

```

This function (**`RegisterNode`**) is used to add a new node to the blockchain network:

- It takes **`url`**, **`ipAddress`**, and **`location`** as inputs to create a new **`Node`** instance.
- It checks if a node with the same URL and IP address already exists in the list of nodes (**`b.Nodes`**). If it exists, **`nodeExists`** is set to **`true`**.
- If the node does not exist, the new node is appended to the list of nodes.
- Note: The comment "Handle exception here" indicates that in a real-world scenario, proper error handling logic would be added here to deal with any exceptional cases.

**GetLastBlock:**

```go
// Get last block in the blockchain
func (b *Blockchain) GetLastBlock() Block {
	lastBlock := b.Chain[len(b.Chain)-1]
	return lastBlock
}

```

This function (**`GetLastBlock`**) retrieves the last block in the blockchain:

- It returns the last block in the **`Chain`** slice, which is accessed using the index **`len(b.Chain)-1`**.

### Handlers

**MineHandler:**

```go

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

```

This handler (**`MineHandler`**) is used for mining new blocks in the blockchain:

- It checks if the HTTP method is POST; otherwise, it responds with "Method Not Allowed".
- It calculates the proof of work using the **`ProofOfWork`** function and the last proof from the last block in the chain.
- A final transaction is added to the current transactions, indicating the creation of a new block.
- The previous hash is calculated, and a new block is added to the blockchain.
- The response is formatted as JSON and contains information about the newly created block.
- The JSON response is written to the client with a status code of 200 (OK).

**CreateTransactionHandler**

```go

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

```

This handler (**`CreateTransactionHandler`**) is used for creating new transactions:

- It checks if the HTTP method is POST; otherwise, it responds with "Method Not Allowed".
- It decodes the incoming JSON request body to extract the transaction data.
- The transaction data is used to add a new transaction to the blockchain's current transactions.
- The index of the block where the transaction will be added is obtained.
- A response message is created, indicating the block where the transaction will be added.
- The response message is converted to JSON and written to the client with a status code of 201 (Created).

**ReadChainHandler:**

```go

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

```

This handler (**`ReadChainHandler`**) is used for reading the blockchain:

- It checks if the HTTP method is GET; otherwise, it responds with "Method Not Allowed".
- It constructs a response map containing the chain, nodes, and length of the blockchain.
- The response map is converted to JSON and written to the client with a status code of 200 (OK).

**RegisterNodeHandler**

```go

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

```

This handler (**`RegisterNodeHandler`**) is used for registering new nodes in the blockchain network:

- It checks if the HTTP method is POST; otherwise, it responds with "Method Not Allowed".
- It decodes the incoming JSON request body to extract the node data.
- The node data is used to register the node in the blockchain's list of nodes.
- A response message is created, indicating that the node has been added to the network.
- The response message is converted to JSON and written to the client with a status code of 201 (Created).

**ResolveNodeHandler**

```go

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
			"chain":   bc.Chain,
		}
	} else {
		response = map[string]interface{}{
			"message": "Chain has not been replaced and is the valid chain.",
			"chain":   bc.Chain,
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

```

This handler (**`ResolveNodeHandler`**) is used for resolving conflicts and achieving consensus in the blockchain network:

- It checks if the HTTP method is PUT; otherwise, it responds with "Method Not Allowed".
- It invokes the **`ResolveChainConflicts`** method to attempt to replace the current chain with a longer and validated chain.
- Depending on whether the chain was substituted or not, a response message is created.
- The response message is converted to JSON and written to the client with a status code of 200 (OK).
