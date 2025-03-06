package utils

import (
	"fmt"
	"math/big"
	"strconv"

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

func ValidatePrice(price string) error {
	priceInt, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return fmt.Errorf("invalid price: %s", price)
	}
	if priceInt < 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	return nil
}
