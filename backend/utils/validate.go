package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func ValidateToAddress(owner string) error {
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
