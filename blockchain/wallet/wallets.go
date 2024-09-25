package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

const WalletPath = "/tmp/wallet.data"

type Wallets struct {
	Wallets map[string]*Wallet
}

func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(WalletPath); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets

	fileContent, err := os.ReadFile(WalletPath)
	if err != nil {
		return err
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		return err
	}

	ws.Wallets = wallets.Wallets

	return nil
}

func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile()
	if err != nil {
		log.Panic(err)
	}

	return &wallets, nil
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()

	address := fmt.Sprintf("%s", wallet.Address())

	ws.Wallets[address] = wallet

	return address

}

func (ws *Wallets) GetAllAddress() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) SaveFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encode := gob.NewEncoder(&content)
	err := encode.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(WalletPath, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
