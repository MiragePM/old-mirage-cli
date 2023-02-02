package parsers

import (
	log "mirage-cli/packages/logger"

	"github.com/BurntSushi/toml"
)

type PackageInfo struct {
	Name               string
	Description        string
	Authors            []string
	Dependencies       []string
	AfterInstallScript string
	RunApplication     []string
}

func ParsePackageInfo(path string) PackageInfo {
	var pkgInfo PackageInfo

	_, err := toml.DecodeFile(path+"/Mirage.toml", &pkgInfo)
	if err != nil {
		(&log.Message{
			Type:    log.Error,
			Message: "There are some problem caused parsing Mirage.toml file.",
		}).Log()
	}

	return pkgInfo
}
