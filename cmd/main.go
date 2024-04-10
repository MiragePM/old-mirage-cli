package main

import (
	"os"
	"quanta-cli/internal"
	log "quanta-cli/packages/logger"
)

func main() {
	internal.Initialize()

	if err := internal.App.Run(os.Args); err != nil {
		(&log.Message{Type: log.Error, Message: "There is some kind of error"}).Log()
	}
}
