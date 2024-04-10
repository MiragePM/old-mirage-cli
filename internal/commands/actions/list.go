package actions

import (
	"fmt"
	"quanta-cli/internal/parsers"
	log "quanta-cli/packages/logger"

	"github.com/urfave/cli/v2"
)

func ListAction(ctx *cli.Context) error {
	dirs := parsers.GetArrayOfDirs()

	if len(dirs) > 0 {

		(log.Message{Type: log.Good, Message: "Now we`ll show all installed packages: "}).Log()

		for _, dir := range dirs {
			if dir != ".git" {
				fmt.Println(" * ", dir)
			}
		}
	}

	return nil
}
