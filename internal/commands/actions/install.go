package actions

import (
	inf "mirage-cli/packages/informer"
	"strings"

	"github.com/urfave/cli/v2"
)

func InstallAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if len(strings.TrimSpace(name)) <= 1 {
		inf.Inform("error", "No one package was found")
		return nil
	}

	return nil
}
