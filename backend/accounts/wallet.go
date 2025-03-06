package accounts

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"nft-marketplace/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SingInTransaction(address string) {
	config := config.LoadConfig()

	client, err := ethclient.Dial(config.BlockChainRPC)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}

	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	account := ks.Accounts()[0]

	err = ks.Unlock(account, "password")
	if err != nil {
		log.Fatal("Failed to unlock keystore:", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), account.Address)
	if err != nil {
		log.Fatal("Failed to get pending nonce:", err)
	}

	toAddress := common.HexToAddress(address)
	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to suggest gas price:", err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Failed to get network ID:", err)
	}

	signedTx, err := ks.SignTx(account, tx, chainID)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Failed to send transaction:", err)
	}

	log.Printf("Transaction sent: %s", signedTx.Hash().Hex())
}

func SingDate(privateKey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal("Failed to sign data:", err)
		return nil, err
	}

	return signature, nil
}

func GetBalance(client *ethclient.Client, address string) (*big.Float, error) {
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	return value, nil
}

func GetTransactor(privateKeyHex string, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Invalid public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(21000)
	auth.GasPrice = gasPrice

	return auth, nil
}
