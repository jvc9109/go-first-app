package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"a": "cervese1",
	"b": "sermesa 2",
}

func main() {
	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("You must specified a command beers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "find by ID")
		beersCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}
}
