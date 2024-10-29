package main

import (
	"log"
	"os"

	"github.com/mgred/html-filer/pkg/cli"
)

func main() {
	opts, err := cli.ProcessArgs()

	handleError(err)

	handleError(RunApp(opts))

	os.Exit(0)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
