package parsers

import (
	"mirage-cli/packages/informer"

	"github.com/BurntSushi/toml"
)

func ParseNodes() ([]string, error) {
	var nodes []string
	var nodesList struct {
		Nodes []string
	}

	_, err := toml.DecodeFile(_pathToConfig, &nodesList)

	if err != nil {
		informer.Inform("error", "While checking available nodes... Please, check your configuration file at: "+_pathToConfig)
		return nil, err
	}

	nodes = append(nodes, nodesList.Nodes...)

	return nodes, nil
}
