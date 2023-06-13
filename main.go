package main

import (
	"fmt"
	"go_blockchain/blockchain"
)



func main () {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	chain.AddBlock("Fourth block")

	for _, block := range chain.Blocks {
		fmt.Printf("PreviousHash: %x\n", block.Prevhash )
		fmt.Printf("Block data: %s\n", block.Data )
		fmt.Printf("Hash: %x\n", block.Prevhash )
	}
}

