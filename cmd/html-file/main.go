package main

import (
	"log"
	"os"

	"github.com/mgred/html-filer/pkg/cli"
)

var opts = cli.Options{
	Base: "/",
}

func main() {
	handleError(cli.ProcessArgs(os.Args[1:], &opts))

	handleError(RunApp(&opts))

	os.Exit(0)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
