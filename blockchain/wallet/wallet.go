package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const (
	checkSumLenght = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivatKEy ecdsa.PrivateKey
	PublicKey []byte
}

func NewPairKey() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Println(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	varsionHash := append([]byte{version}, pubHash...)
	checkSum := CheckSum(varsionHash)

	fullHash := append(varsionHash, checkSum...)
	address := Base58Encode(fullHash)

	fmt.Printf("Pub Key: %s\n", w.PublicKey)
	fmt.Printf("pub hash: %s\n", pubHash)
	fmt.Printf("address: %s\n", address)

	return address
}

func MakeWallet() *Wallet {
	private, public := NewPairKey()
	Wallet := Wallet{
		PrivatKEy: private,
		PublicKey: public,
	}

	return &Wallet
}

func PublicKeyHash(pubkey []byte) []byte {
	pubHash := sha256.Sum256(pubkey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Println(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func CheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checkSumLenght]
}
