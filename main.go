package main

import (
	"fmt"

	"github.com/blinktag/blockchain-go/block"
)

func main() {
	// Create root block
	b := block.NewBlock(100, "")

	// Add sample transaction
	b.NewTransaction("me@myself.com", "john@doe.com", 99)

	// Mine block
	b.Mine()

	// Print out our current block chain
	for _, block := range block.BlockChain {
		fmt.Printf("%#v\n\n", block)
	}
}
