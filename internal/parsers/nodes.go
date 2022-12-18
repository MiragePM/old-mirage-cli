package parsers

import (
	"mirage-cli/packages/informer"
	"os"

	"github.com/BurntSushi/toml"
)

func ParseNodes() ([]string, error) {
	var nodes []string
	var nodesList struct {
		Nodes []string
	}

	homePath, _ := os.UserHomeDir()
	pathToConfig := homePath + "/.config/mirage/nodes.toml"

	_, err := toml.DecodeFile(pathToConfig, &nodesList)

	if err != nil {
		informer.Inform("error", "While checking avaliable nodes... Please, check your configuration file at: "+pathToConfig)
		return nil, err
	}

	nodes = append(nodes, nodesList.Nodes...)

	return nodes, nil
}
