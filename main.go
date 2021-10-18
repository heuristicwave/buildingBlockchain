package main

import (
	"githun.com/heuristicwave/buildingBlockchain/explorer"
	"githun.com/heuristicwave/buildingBlockchain/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
