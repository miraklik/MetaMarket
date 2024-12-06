package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func ValidateEthereumAddress(owner string) error {
	if !common.IsHexAddress(owner) {
		return fmt.Errorf("invalid address: %s", owner)
	}

	return nil
}

func ValidateFromAddress(from string) error {
	if !common.IsHexAddress(from) {
		return fmt.Errorf("invalid address: %s", from)
	}

	return nil
}

func ValidateAmount(amount string) error {
	if _, ok := new(big.Int).SetString(amount, 10); !ok {
		return fmt.Errorf("invalid amount: %s", amount)
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 5 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}
