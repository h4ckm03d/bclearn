package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash is function to calculate hash
func (b *Block) SetHash() {
	ts := []byte(string(strconv.FormatInt(b.Timestamp, 10)))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, ts}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//ToString conver data to string
func (b *Block) ToString() string {
	return fmt.Sprintf("Prev. hash: %x\nData: %s\nHash: %x\n", b.PrevBlockHash, b.Data, b.Hash)
}

// NewBlock to create new instance of block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
