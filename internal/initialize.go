package internal

import (
	"strconv"
	"time"

	"mirage-cli/internal/commands"

	"github.com/urfave/cli/v2"
)

var App cli.App

func Initialize() {
	checkExistConf()

	App.Name = "mirage"
	App.Usage = "lightweight package manager"
	App.UsageText = "mirage [command] [name of package] [flags]"
	App.Copyright = "Copyright (c) " + strconv.Itoa(time.Now().Year()) + " ZueffC. Licensed under WTFPL License."

	App.Authors = []*cli.Author{
		{
			Name:  "Зуев Даниил",
			Email: "zueffc@gmail.com",
		},
	}

	App.Commands = []*cli.Command{
		commands.Search(),
		commands.Install(),
		//commands.Uninstall(),
		//commands.List(),
	}
}
