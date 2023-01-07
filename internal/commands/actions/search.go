package actions

import (
	. "fmt"
	"mirage-cli/internal/additions"
	"mirage-cli/internal/parsers"
	"mirage-cli/packages/informer"
	"strings"

	"github.com/urfave/cli/v2"
)

func SearchAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if len(strings.TrimSpace(name)) <= 1 {
		informer.Inform("error", "No one package was found")
		return nil
	}

	nodes, err := parsers.ParseNodes()

	for iter, url := range nodes {
		informer.Inform("info", Sprintf("Searching for package `%s`;", name))
		informer.Inform("info", Sprintf("Now we`re calling for `%s` node;", url))

		pkg, error := additions.SearchByNameQuery(name, url)

		if error != nil || len(pkg.Name) <= 1 {
			informer.Inform("warn", Sprintf("Well, problems in the %d node. Maybe it doesn't work, maybe the package just doesn't exist.", iter))
		}

		if (iter + 1) == len(nodes) {
			informer.Inform("error", "After calling all APIs from your config, we didn't find anything. Maybe you a typo?")
		}
	}

	if err != nil {
		informer.Inform("error", "Some problem caused by parsing your config. Pls check it in default folder")
	}

	return nil
}
