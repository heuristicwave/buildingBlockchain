package main

import (
	"githun.com/heuristicwave/buildingBlockchain/cli"
	"githun.com/heuristicwave/buildingBlockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
