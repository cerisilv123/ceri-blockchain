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
