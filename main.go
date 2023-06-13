package main

import (
	"fmt"
)


func main () {
	chain := InitBlockchain()

	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	chain.AddBlock("Fourth block")

	for _, block := range chain.blocks {
		fmt.Printf("PreviousHash: %x\n", block.Prevhash )
		fmt.Printf("Block data: %s\n", block.Data )
		fmt.Printf("Hash: %x\n", block.Prevhash )
	}
}

