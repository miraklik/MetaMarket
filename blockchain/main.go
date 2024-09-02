package main

import (
	"fmt"
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous block: %x\n", block.PrevHash)
		fmt.Printf("Data is block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
