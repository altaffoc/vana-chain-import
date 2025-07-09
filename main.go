package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// Replace with the VANA token contract address
const vanaTokenAddress = "0x..." // VANA token contract address
const vanaRPC = "https://rpc.vananet.org" // VANA RPC endpoint

// Minimal ERC20 ABI: balanceOf & decimals
const erc20ABI = `[{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"type":"function"},
{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"type":"function"}]`

func main() {
	// 🔐 Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 🔑 Retrieve private key from .env
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("PRIVATE_KEY not found in .env file")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Invalid public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("🔓 Wallet Address:", fromAddress.Hex())

	// 🌐 Connect to the VANA RPC
	client, err := ethclient.Dial(vanaRPC)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	defer client.Close()

	// 🔗 Load the ERC20 contract
	tokenAddress := common.HexToAddress(vanaTokenAddress)
	instance := bind.NewBoundContract(tokenAddress, erc20ABI, client, client, client)

	// 💰 Retrieve token balance
	var balance *big.Int
	err = instance.Call(nil, &balance, "balanceOf", fromAddress)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	// 🔢 Retrieve token decimals
	var decimals uint8
	err = instance.Call(nil, &decimals, "decimals")
	if err != nil {
		log.Fatalf("Failed to get token decimals: %v", err)
	}

	// 💵 Print balance with decimal formatting
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	balanceFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(divisor))
	fmt.Printf("💰 VANA Balance: %f\n", balanceFloat)
}
