package actions

import (
	"fmt"
	"mirage-cli/internal/parsers"
	inf "mirage-cli/packages/informer"
	"strings"

	"github.com/urfave/cli/v2"
)

func SearchAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if len(strings.TrimSpace(name)) <= 1 {
		inf.Inform("error", "No one package was found")
		return nil
	}

	nodes, err := parsers.ParseNodes()

	if err != nil {
		panic(err)
	}

	fmt.Println(nodes)

	return nil
}
