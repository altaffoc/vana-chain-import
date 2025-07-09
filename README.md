# VANA Token Balance Checker

A lightweight Go application to securely import an Ethereum-compatible wallet from a private key and check its VANA token balance on the VANA network.

This project demonstrates minimalistic interaction with an ERC-20 smart contract using the Go Ethereum SDK, ideal for scripting, tooling, or integrations with on-chain token data.

---

## Features

- Load private key securely from `.env` (never hardcoded)
- Connect to the VANA RPC endpoint
- Derive wallet address from private key
- Interact with VANA token contract using minimal ABI
- Retrieve and format token balance based on decimals
- Built with `go-ethereum` and `godotenv` libraries

---

## Prerequisites

- Go 1.18 or higher
- Git (optional, for cloning)
- A `.env` file with your Ethereum-compatible private key
- VANA token contract address (ERC-20)

---

## Setup

1. Clone this repository or copy `main.go` into your project directory.

2. Install dependencies:

```bash
go get github.com/ethereum/go-ethereum
go get github.com/joho/godotenv
