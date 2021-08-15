package main

import (
	"fmt"

	"githun.com/heuristicwave/buildingBlockchain/person"
)

func main() {
	my := person.Person{}
	my.SetDetails("heuri", 7)
	fmt.Println("Main : ", my)
}
