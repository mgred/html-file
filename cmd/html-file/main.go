package main

import (
	"log"
)

func main() {
	opts, err := ProcessArgs()

	handleError(err)

	handleError(RunApp(opts))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
