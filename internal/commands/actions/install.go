package actions

import (
	log "mirage-cli/packages/logger"
	"strings"

	"github.com/urfave/cli/v2"
)

func InstallAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if len(strings.TrimSpace(name)) <= 1 {
		(&log.Message{Type: log.Error, Message: "No one package was found"}).Log()
		return nil
	}

	return nil
}
