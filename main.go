package main

import (
	"githun.com/heuristicwave/buildingBlockchain/blockchain"
	"githun.com/heuristicwave/buildingBlockchain/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
