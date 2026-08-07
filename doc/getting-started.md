# Getting Started with Eterbit Core

Follow this guide to set up your development environment, compile the binary, and run your first Eterbit node.

## Prerequisites
* **Go (Golang):** Version 1.20 or higher recommended.
* **OS:** Linux (Ubuntu/Debian/Termux) or macOS.

## Installation & Compilation
Clone the repository to your local machine and compile the source code:

```bash
git clone [https://github.com/eterbit-chain/eterbit.git
cd eterbit
go build -o eterbit eterbit.go
```

## Managing Wallets & Accounts
Before interacting with the ledger or mining coins, create your local cryptographic account profile:

```bash
./eterbit create -label "MainNode"
```

## Running a Node & Mining
Start a local network node instance or test block mining manually:

```bash
# Check initial account balance
./eterbit balance

# Start the P2P network node instance
./eterbit node

# Run manual proof-of-work mining targeting your address
./eterbit mining <your_address>
```

## Connecting to Peers & Network Options
To connect your node to other peers on the network or check network statistics, use the following optional flags and commands:

```bash
# Connect to a specific peer node manually
./eterbit addnode 192.168.1.50:19333

# Or start your node with an initial bootstrap connection and custom port
./eterbit node --port :19333 --connect 192.168.1.50:19333

# Check active P2P network connections
./eterbit peers
```

## Checking Supply. Uptime. blockchain size
You can monitor your node's operational health and token metrics at any time using:

```bash
# Check node uptime duration
./eterbit uptime

# Check circulating supply and remaining coins to be mined
./eterbit supply

# View total blockchain storage size and location
./eterbit blocksize
```

## Transactions & Fee Management
To transfer coins to another address or monitor transaction fee statistics in the mempool, use:

```bash
# Send coins to another address with an optional fee and custom sender address
./eterbit send -to <recipient_address> -amount <val> -fee <val> -from <sender_address>

# Analyze fee market statistics derived from the active mempool
./eterbit fees
```

## Blockchain Inspection
To inspect individual blocks, retrieve block hashes by index, or parse raw blockchain data from storage, use:

```bash
# Retrieve the block hash corresponding to a numerical index
./eterbit getblockhash <index>

# Retrieve and render complete structural block data in JSON format based on a hash
./eterbit getblock <hash>

# Parse and inspect structural blockchain blocks directly from disk storage
./eterbit explorer
```

## Troubleshooting

### Port Already in Use
If your node fails to start with an error indicating that the port (e.g., `:19333`) is already bound, you can specify a different custom port using the `--port` flag:

```bash
./eterbit node --port :19334
```

## Resetting Blockchain Data
If your local ledger database becomes out of sync or corrupted during testing, you can safely remove the local data directory to start fresh (make sure to back up your wallet.dat file first):

```bash
rm -rf ~/.eterbit
```
