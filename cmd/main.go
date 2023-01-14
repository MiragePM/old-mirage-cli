package main

import (
	"mirage-cli/internal"
	inf "mirage-cli/packages/informer"
	"os"
)

func main() {
	internal.Initialize()

	if err := internal.App.Run(os.Args); err != nil {
		inf.Inform("error", "There is some kind of error")
	}
}
