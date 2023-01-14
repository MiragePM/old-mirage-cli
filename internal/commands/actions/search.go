package actions

import (
	"fmt"
	"mirage-cli/internal/additions"
	"mirage-cli/internal/parsers"
	inf "mirage-cli/packages/informer"
	"strings"

	"github.com/urfave/cli/v2"
)

func SearchAction(ctx *cli.Context) error {
	var agreement string
	name := ctx.Args().Get(0)

	if len(strings.TrimSpace(name)) <= 1 {
		inf.Inform("error", "No one package was found")
		return nil
	}

	nodes, err := parsers.ParseNodes()

	for iter, url := range nodes {
		inf.Inform("info", fmt.Sprintf("Searching for package `%s`;", name))
		inf.Inform("info", fmt.Sprintf("Now we`re calling for `%s` node;", url))

		pkg, error := additions.SearchByNameQuery(name, url)

		if error == nil && len(pkg.GitUrl) >= 19 {
			inf.Inform("good", "We`ve found a package by name, now we`re showing you package details; \n")
			inf.Inform("good", "Are you sure it`s correct? ", false)

			if _, err := fmt.Scan(&agreement); err != nil {
				inf.Inform("error", "You have entered something strange, please try again or create issue")
			}

			if agreement = strings.ToLower(agreement); strings.Contains(agreement, "y") {
				inf.Inform("good", "We`ll take stock this choice")
			}
			inf.Inform("info", "Ok, we`ll search more...")
		}

		if error != nil || len(pkg.GitUrl) <= 19 {
			inf.Inform("warn", fmt.Sprintf("Well, problems in the %d node. Maybe it doesn't work, maybe the package just doesn't exist.", iter))
		}

		if (iter + 1) == len(nodes) {
			inf.Inform("error", "After calling all APIs from your config, we didn't find anything. Maybe you a typo?")
		}
	}

	if err != nil {
		inf.Inform("error", "Some problem caused by parsing your config. Pls check it in default folder")
	}

	return nil
}
