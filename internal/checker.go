package internal

import (
	"bytes"
	"errors"
	"fmt"
	"mirage-cli/internal/parsers"

	log "mirage-cli/packages/logger"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func inpUrl() string {
	var input string
	fmt.Scanln(&input)

	if len(input) <= 5 && IsUrl(input) {
		(&log.Message{Type: log.Error, Message: "Your URL is invalid, please try again: "}).Log()
		return inpUrl()
	}

	return input
}

func checkExistConf() {
	homePath, _ := os.UserHomeDir()
	pathToConfigFolder := homePath + "/.config/mirage/"
	cnfPath := pathToConfigFolder + "nodes.toml"

	nodesArray, _ := parsers.ParseNodes()

	if _, err := os.Stat(cnfPath); errors.Is(err, os.ErrNotExist) || len(nodesArray) <= 0 {
		(&log.Message{Type: log.Error, Message: "No one config file exists, creating one..."}).Log()
		(&log.Message{
			Type:    log.Error,
			Message: "Please input one node url (e.g. http://zueffc.ml:1984): ",
			NoBreak: true,
		}).Log()

		url := inpUrl()

		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(map[string]interface{}{
			"Nodes": []string{parsers.NormalizeURL(url)},
		})

		if err != nil {
			(&log.Message{Type: log.Error, Message: "Error while adding url, please retry or create issue..."}).Log()
			return
		}

		os.Mkdir(pathToConfigFolder, os.ModePerm)
		os.WriteFile(pathToConfigFolder+"nodes.toml", buf.Bytes(), 0755)
	}
}
