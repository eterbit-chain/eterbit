# Getting Started with Eterbit Core

Follow this guide to set up your development environment, compile the binary, and run your first Eterbit node.

## Prerequisites
* **Go (Golang):** Version 1.20 or higher recommended.
* **OS:** Linux (Ubuntu/Debian/Termux) or macOS.

## 1. Installation & Compilation
Clone the repository to your local machine and compile the source code:

```bash
git clone [https://github.com/eterbit-chain/eterbit.git
cd eterbit
go build -o eterbit eterbit.go
```

## 2. Managing Wallets & Accounts
Before interacting with the ledger or mining coins, create your local cryptographic account profile:

```bash
./eterbit create -label "MainNode"
