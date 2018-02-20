package main

// Blockchain represent blockchain
type Blockchain struct {
	blocks []*Block
}

//AddBlock to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//NewGenesisBlock to create genesis
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//NewBlockchain to init blockchan
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
