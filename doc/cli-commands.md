# Eterbit CLI Reference Guide

All persistent node data, LevelDB blockchain state, multi-wallet credentials (`wallet.dat`), transaction mempool, and P2P peer tables are securely stored in the centralized external directory at `~/.eterbit/`.

*(Note: You can use `go run eterbit.go <command>` or build the binary and use `./eterbit <command>`)*

## Available Commands

| Command | Description |
| :--- | :--- |
| `create [-label <account_label>]` | Provisions a new post-quantum cryptographic keypair account inside the centralized wallet container (`wallet.dat`). |
| `balance` | Queries the state database and displays all registered account balances and nonces. |
| `supply` | Displays maximum coin supply, circulating supply, and remaining coins available to mine. |
| `send -to <addr> -amount <val> [-fee <val>] [-from <sender_addr>]` | Constructs, digitally signs (Dilithium Mode 3), and broadcasts a value transfer transaction to the mempool. |
| `node [--port :port] [--connect host:port]` | Starts the P2P networking server and handles incoming/outgoing peer connections. |
| `addnode <host:port>` | Manually registers and adds a target peer address into the addrman database. |
| `mine [-blocks <num>] [-address <addr>]` | Executes iterative Proof-of-Work block mining with optional block count and target reward address flags. |
| `mining <target_address>` | Shortcut command to execute manual block mining targeting a specific reward address. |
| `explorer` | Parses and inspects structural blockchain blocks directly from disk storage. |
| `peers` | Displays the list and active status of currently connected P2P network peers. |
| `fees` | Analyzes fee market statistics (highest priority fee, average fee, pending count) derived from the active mempool. |
| `uptime` | Computes and displays the active operational duration of the node instance. |
| `getnettotals` | Retrieves and outputs network traffic statistics in JSON format. |
| `blocksize` | Displays the physical storage path and current disk size occupied by the node database. |
| `getblockhash <index>` | Retrieves and outputs the hexadecimal block hash corresponding to a numerical block index. |
| `getblock <hash>` | Retrieves and renders complete structural block data in JSON format based on a target hash. |
