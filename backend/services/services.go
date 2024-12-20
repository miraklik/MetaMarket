package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"nft-marketplace/db"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumService struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	PrivateKey      *ecdsa.PrivateKey
	Contract        *bind.BoundContract
}

type NFTContract struct {
	*bind.BoundContract
}

// NewEthereumService creates a new instance of EthereumService.
//
// It takes two string parameters: rpcURL and contractAddress. The rpcURL is the
// URL of the Ethereum node to connect to, and the contractAddress is the address
// of the smart contract to interact with.
func NewEthereumService(rpcURL, contractAddress, privateKeyHex, abiJSON string, chainID *big.Int) (*EthereumService, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("Invalid private key: %v", err)
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Printf("Failed to connect to Ethereum client: %v", err)
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Printf("Failed to parse contract ABI: %v", err)
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		return nil, err
	}

	auth.GasLimit = uint64(22000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return nil, err
	}
	auth.GasPrice = new(big.Int).Div(gasPrice, big.NewInt(2))

	contract := bind.NewBoundContract(common.HexToAddress(contractAddress), parsedABI, client, client, client)

	return &EthereumService{
		Client:          client,
		ContractAddress: common.HexToAddress(contractAddress),
		PrivateKey:      privateKey,
		Contract:        contract,
	}, nil
}

// CheckOwnership checks whether the given token ID is owned by the given Ethereum address.
//
// It will query the smart contract and compare the owner of the token ID with the given ownerAddress.
//
// If the owner matches, it will return true; otherwise, it will return false.
//
// If there is an error during the smart contract query, it will return false and log the error.
func (es *EthereumService) CheckOwnership(tokenID string, ownerAddress string) bool {
	tokenIDBigInt := new(big.Int)
	tokenIDBigInt.SetString(tokenID, 10)

	owner := common.HexToAddress(ownerAddress)

	var actualOwner common.Address
	err := es.Contract.Call(nil, &[]interface{}{actualOwner}, "ownerOf", tokenIDBigInt)
	if err != nil {
		log.Printf("Error checking ownership: %v\n", err)
		return false
	}

	return actualOwner == owner
}

func (es *EthereumService) GetBalance(address common.Address) (*big.Int, error) {
	var balance big.Int
	err := es.Contract.Call(nil, &[]interface{}{balance}, "balanceOf", address)
	if err != nil {
		log.Printf("Failed to fetch balance: %v", err)
		return nil, fmt.Errorf("failed to fetch balance: %w", err)
	}
	return &balance, nil
}

// GetNFTs returns a list of NFTs owned by the given address. Currently, this
// function is not implemented and will return an error.
func (es *EthereumService) GetNFTs(accounts common.Address) ([]*big.Int, error) {
	log.Printf("Getting NFTs for address: %s\n", accounts.Hex())

	contractABI := `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_nftContractAddress",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_commissionPercent",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "newPercent",
				"type": "uint256"
			}
		],
		"name": "CommissionUpdated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "message",
				"type": "string"
			}
		],
		"name": "Debug",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "FundsWithdrawn",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "seller",
				"type": "address"
			}
		],
		"name": "ListingCancelled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "seller",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "ListingCreated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "buyer",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "PurchaseCompleted",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_listingId",
				"type": "uint256"
			}
		],
		"name": "cancelListing",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "commissionPercent",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint128",
				"name": "_tokenId",
				"type": "uint128"
			},
			{
				"internalType": "uint128",
				"name": "_price",
				"type": "uint128"
			}
		],
		"name": "createListing",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "listings",
		"outputs": [
			{
				"internalType": "address",
				"name": "seller",
				"type": "address"
			},
			{
				"internalType": "uint128",
				"name": "tokenId",
				"type": "uint128"
			},
			{
				"internalType": "uint128",
				"name": "price",
				"type": "uint128"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftContract",
		"outputs": [
			{
				"internalType": "contract IERC721",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftEnumerable",
		"outputs": [
			{
				"internalType": "contract IERC721Enumerable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address payable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "pendingWithdrawals",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_listingId",
				"type": "uint256"
			}
		],
		"name": "purchaseListing",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_newPercent",
				"type": "uint256"
			}
		],
		"name": "setCommissionPercent",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_owner",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "index",
				"type": "uint256"
			}
		],
		"name": "tokenOfOwnerByIndex",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawFunds",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
]`

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Printf("Failed to parse ABI: %v", err)
		return nil, err
	}

	contract := bind.NewBoundContract(es.ContractAddress, parsedABI, es.Client, es.Client, es.Client)

	var balance *big.Int
	err = contract.Call(nil, &[]interface{}{&balance}, "balanceOf", accounts)
	if err != nil {
		log.Printf("Failed to call balanceOf: %v", err)
		return nil, err
	}

	tokens := make([]*big.Int, 0, balance.Uint64())
	for i := uint64(0); i < balance.Uint64(); i++ {
		var tokenId big.Int
		err := contract.Call(nil, &[]interface{}{tokenId}, "tokenOfOwnerByIndex", accounts, big.NewInt(int64(i)))
		if err != nil {
			log.Printf("Failed to call tokenOfOwnerByIndex: %v", err)
			return nil, err
		}
		tokens = append(tokens, &tokenId)
	}
	return tokens, nil
}

// MintNFT creates a new NFT and lists it on the marketplace with the given name, symbol, description, and price.
func (es *EthereumService) MintNFT(tokenID, price, recipient string) error {
	log.Printf("Minting NFT with token ID: %s for recipient: %s with price: %s, name: %s, symbol: %s, description: %s", tokenID, recipient, price, name, symbol, description)

	if !common.IsHexAddress(recipient) {
		log.Printf("Invalid recipient address: %s", recipient)
		return fmt.Errorf("invalid recipient address")
	}

	recipientAddress := common.HexToAddress(recipient)
	if recipientAddress == (common.Address{}) {
		log.Printf("Invalid recipient address: %s", recipient)
		return fmt.Errorf("invalid recioient address")
	}

	tokenIDBigInt := new(big.Int)
	if _, ok := tokenIDBigInt.SetString(tokenID, 10); !ok {
		log.Printf("Invalid token ID: %s", tokenID)
		return fmt.Errorf("invalid token ID: %s", tokenID)
	}

	priceBigInt := new(big.Int)
	if _, ok := priceBigInt.SetString(price, 10); !ok {
		log.Printf("Invalid price: %s", price)
		return fmt.Errorf("invalid price: %s", price)
	}

	chainID, err := es.Client.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Printf("Invalid to load private key: %v", err)
		return fmt.Errorf("invalid to load private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.GasLimit = uint64(22000)
	gasPrice, err := es.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}
	auth.GasPrice = new(big.Int).Div(gasPrice, big.NewInt(2))

	contractABI := `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_nftContractAddress",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_commissionPercent",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "newPercent",
				"type": "uint256"
			}
		],
		"name": "CommissionUpdated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "message",
				"type": "string"
			}
		],
		"name": "Debug",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "FundsWithdrawn",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "seller",
				"type": "address"
			}
		],
		"name": "ListingCancelled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "seller",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "ListingCreated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "buyer",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "PurchaseCompleted",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_listingId",
				"type": "uint256"
			}
		],
		"name": "cancelListing",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "commissionPercent",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint128",
				"name": "_tokenId",
				"type": "uint128"
			},
			{
				"internalType": "uint128",
				"name": "_price",
				"type": "uint128"
			}
		],
		"name": "createListing",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "listings",
		"outputs": [
			{
				"internalType": "address",
				"name": "seller",
				"type": "address"
			},
			{
				"internalType": "uint128",
				"name": "tokenId",
				"type": "uint128"
			},
			{
				"internalType": "uint128",
				"name": "price",
				"type": "uint128"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftContract",
		"outputs": [
			{
				"internalType": "contract IERC721",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftEnumerable",
		"outputs": [
			{
				"internalType": "contract IERC721Enumerable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address payable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "pendingWithdrawals",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_listingId",
				"type": "uint256"
			}
		],
		"name": "purchaseListing",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_newPercent",
				"type": "uint256"
			}
		],
		"name": "setCommissionPercent",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_owner",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "index",
				"type": "uint256"
			}
		],
		"name": "tokenOfOwnerByIndex",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawFunds",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
]`

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Printf("invalid to parse ABI: %v", err)
		return fmt.Errorf("invalid to parse ABI: %w", err)
	}

	contract := bind.NewBoundContract(es.ContractAddress, parsedABI, es.Client, es.Client, es.Client)

	tx, err := contract.Transact(auth, "createListing", tokenIDBigInt, priceBigInt)
	if err != nil {
		log.Printf("failed to mint NFT: %v", err)
		return fmt.Errorf("failed to mint NFT: %w", err)
	}

	fmt.Printf("NFT minted successfully! Transaction hash: %s\n", tx.Hash().Hex())
	return nil
}

// TransferNFT transfers an NFT to the buyer, given the token ID.
//
// This method will first check if the token ID is valid, and if the buyer's address is valid.
// Then, it will get the chain ID, and construct a new transactor with the private key and chain ID.
// Next, it will get the network ID, and suggest a gas price.
// After that, it will set the gas limit and gas price for the transactor.
// Finally, it will call the transfer function on the contract, and return the transaction hash.
//
// Parameters:
//
//	tokenID: The token ID of the NFT to transfer.
//	buyer: The address of the buyer.
//
// Returns:
//
//	An error if something goes wrong.
func (es *EthereumService) TransferNFT(tokenID, buyer string) error {
	log.Printf("Starting NFT transfer: tokenID=%s, buyer=%s", tokenID, buyer)

	buyerAddress := common.HexToAddress(buyer)
	if buyerAddress == (common.Address{}) {
		log.Printf("invalid buyer address")
		return fmt.Errorf("invalid address")
	}
	tokenIDBigInt := new(big.Int)
	if _, ok := tokenIDBigInt.SetString(tokenID, 10); !ok {
		log.Printf("invalid token ID")
		return fmt.Errorf("invalid token ID")
	}

	chainID, err := es.Client.ChainID(context.Background())
	if err != nil {
		log.Printf("failed to get chain ID: %v", err)
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	if es.PrivateKey == nil {
		log.Printf("invalid private key")
		return fmt.Errorf("invalid private key")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(es.PrivateKey, chainID)
	if err != nil {
		log.Printf("failed to create transactor: %v", err)
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	_, err = es.Client.NetworkID(context.Background())
	if err != nil {
		log.Printf("failed to get network ID: %v", err)
		return fmt.Errorf("failed to get network ID: %w", err)
	}

	auth.GasLimit = uint64(22000)
	gasPrice, err := es.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("failed to suggest gas price: %v", err)
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}
	auth.GasPrice = new(big.Int).Div(gasPrice, big.NewInt(2))

	contractABI := `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_nftContractAddress",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_commissionPercent",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "newPercent",
				"type": "uint256"
			}
		],
		"name": "CommissionUpdated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "message",
				"type": "string"
			}
		],
		"name": "Debug",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "FundsWithdrawn",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "seller",
				"type": "address"
			}
		],
		"name": "ListingCancelled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "seller",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "ListingCreated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "buyer",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "price",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "timestamp",
				"type": "uint256"
			}
		],
		"name": "PurchaseCompleted",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_listingId",
				"type": "uint256"
			}
		],
		"name": "cancelListing",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "commissionPercent",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint128",
				"name": "_tokenId",
				"type": "uint128"
			},
			{
				"internalType": "uint128",
				"name": "_price",
				"type": "uint128"
			}
		],
		"name": "createListing",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "listings",
		"outputs": [
			{
				"internalType": "address",
				"name": "seller",
				"type": "address"
			},
			{
				"internalType": "uint128",
				"name": "tokenId",
				"type": "uint128"
			},
			{
				"internalType": "uint128",
				"name": "price",
				"type": "uint128"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftContract",
		"outputs": [
			{
				"internalType": "contract IERC721",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "nftEnumerable",
		"outputs": [
			{
				"internalType": "contract IERC721Enumerable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address payable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "pendingWithdrawals",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_listingId",
				"type": "uint256"
			}
		],
		"name": "purchaseListing",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_newPercent",
				"type": "uint256"
			}
		],
		"name": "setCommissionPercent",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_owner",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "index",
				"type": "uint256"
			}
		],
		"name": "tokenOfOwnerByIndex",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawFunds",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
]`

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Printf("Failed to parse contract ABI: %v", err)
		return fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	contract := bind.NewBoundContract(es.ContractAddress, parsedABI, es.Client, es.Client, es.Client)

	tx, err := contract.Transact(auth, "purchaseListing", tokenIDBigInt)
	if err != nil {
		log.Printf("Failed to transfer NFT: %v", err)
		return fmt.Errorf("failed to transfer NFT: %w", err)
	}

	log.Printf("Transfer successful! Transaction hash: %s", tx.Hash().Hex())
	return nil
}

// SearchNFTs searches for NFTs with the given name in the database.
//
// It takes a single parameter `name` which is the name of the NFT to search for.
// The function logs the search operation and returns a list of NFTs that match
// the given name. If there is an error during the database query, it logs the
// error and returns an empty result along with the error.
//
// Returns:
// - A `db.Nfts` containing the list of NFTs with the specified name.
// - An `error` if the database query fails.
func (es *EthereumService) SearchNFTs(name string) (db.Nfts, error) {
	log.Printf("Searching for NFTs by event with name: %s", name)

	result, err := db.GetNFTsByName(name)
	if err != nil {
		log.Printf("Failed to get NFTs by name: %v", err)
		return db.Nfts{}, fmt.Errorf("failed to get NFTs by name: %w", err)
	}

	return result, nil
}
