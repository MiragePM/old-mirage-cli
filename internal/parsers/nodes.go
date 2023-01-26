package parsers

import (
	log "mirage-cli/packages/logger"

	"github.com/BurntSushi/toml"
)

func ParseNodes() ([]string, error) {
	var nodes []string
	var nodesList struct {
		Nodes []string
	}

	_, err := toml.DecodeFile(_pathToConfig, &nodesList)

	if err != nil {
		(&log.Message{
			Type:    log.Error,
			Message: "While checking available nodes... Please, check your configuration file at: " + _pathToConfig,
		}).Log()

		return nil, err
	}

	nodes = append(nodes, nodesList.Nodes...)

	return nodes, nil
}
