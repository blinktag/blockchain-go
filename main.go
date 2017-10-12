package main

import (
	"fmt"

	"github.com/blinktag/blockchain-go/block"
)

func main() {
	// Create root block
	b := block.NewBlock(100, "")

	// P
	b.NewTransaction("me@myself.com", "john@doe.com", 99)
	b.Mine()

	for _, block := range block.BlockChain {
		fmt.Printf("%#v\n\n", block)
	}
}
