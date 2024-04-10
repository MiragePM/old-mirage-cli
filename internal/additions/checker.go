package additions

import (
	"quanta-cli/internal/parsers"
	"strings"
)

func IsInstalled(name string) (bool, string) {
	arrayOfDirs := parsers.GetArrayOfDirs()

	for _, dirName := range arrayOfDirs {
		if strings.Contains(dirName, name) {
			return true, dirName
		}
	}

	return false, ""
}
