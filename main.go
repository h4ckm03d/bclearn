package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Lutfi")
	bc.AddBlock("Send 2 more BTC to Lutfi")

	for _, block := range bc.blocks {
		fmt.Println(block.ToString())
	}
}
