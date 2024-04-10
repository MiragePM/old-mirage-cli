package actions

import (
	"os"
	"os/exec"
	"quanta-cli/internal/additions"
	"quanta-cli/internal/parsers"
	log "quanta-cli/packages/logger"

	"github.com/urfave/cli/v2"
)

func RunAction(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if flag, dirName := additions.IsInstalled(name); flag {
		path := parsers.ParseHomePath() + "/.mirage/" + dirName
		pkgInfo := parsers.ParsePackageInfo(path)

		for _, v := range pkgInfo.RunApplication {
			cmd := exec.Command("/bin/sh", "-c", v)
			cmd.Stderr, cmd.Stdin, cmd.Stdout = os.Stderr, os.Stdin, os.Stdout
			cmd.Dir = path
			err := cmd.Run()

			if err != nil {
				(&log.Message{
					Type:    log.Error,
					Message: "There is an error in run script.",
				}).Log()
			}
		}

	}

	return nil
}
