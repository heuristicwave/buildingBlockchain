package cli

import (
	"flag"
	"fmt"
	"os"

	"githun.com/heuristicwave/buildingBlockchain/explorer"
	"githun.com/heuristicwave/buildingBlockchain/rest"
)

func usage() {
	fmt.Printf("\nWelcome to Cryptocurrency\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-hport:		Set the PORT of the html server\n")
	fmt.Printf("-rport:		Set the PORT of the rest server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest' or 'all'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	hport := flag.Int("hport", 4000, "Set port of the html server")
	rport := flag.Int("rport", 4001, "Set port of the rest server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' or 'all'")

	flag.Parse()

	switch *mode {
	case "all":
		go rest.Start(*rport)
		explorer.Start(*hport)
	case "rest":
		rest.Start(*rport)
	case "html":
		explorer.Start(*hport)
	default:
		usage()
	}
}
