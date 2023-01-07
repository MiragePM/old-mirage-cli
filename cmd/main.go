package main

import (
	"mirage-cli/internal"
	"mirage-cli/packages/informer"
	"os"
)

func main() {
	internal.Initialize()

	if err := internal.App.Run(os.Args); err != nil {
		informer.Inform("error", "There is some kind of error")
	}
}
