package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/heuristicwave/buildingBlockchain/explorer"
	"github.com/heuristicwave/buildingBlockchain/rest"
)

func usage() {
	fmt.Printf("\nWelcome to Cryptocurrency\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-hport:		Set the PORT of the html server\n")
	fmt.Printf("-rport:		Set the PORT of the rest server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest' or 'all'\n\n")
	// runtime.Goexit()
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	rport := flag.Int("rport", 4000, "Set port of the rest server")
	hport := flag.Int("hport", 4001, "Set port of the html server")
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
