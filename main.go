package main

import (
	"github.com/heuristicwave/buildingBlockchain/cli"
	"github.com/heuristicwave/buildingBlockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
