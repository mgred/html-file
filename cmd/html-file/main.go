package main

import (
	"log"
	"os"
)

func main() {
	opts, err := ProcessArgs()

	handleError(err)

	handleError(RunApp(opts))

	os.Exit(0)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
