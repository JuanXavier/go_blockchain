package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	Prevhash []byte
}

// This code defines a method named "DeriveHash" on the Block struct.
// The method calculates the hash value for a block based on its data and previous hash.
func (b *Block) DeriveHash() {
	// Concatenate the data and previous hash using bytes.Join()
	info := bytes.Join([][]byte{b.Data, b.Prevhash}, []byte{})

	// Calculate the SHA-256 hash of the concatenated info
	hash := sha256.Sum256(info)

	// Assign the calculated hash value to the Hash field of the Block struct
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
