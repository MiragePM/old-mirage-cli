package informer

import (
	"mirage-cli/internal/additions"

	"github.com/fatih/color"
)

func Inform(typeOf string, message string) {
	if typeOf == "error" {
		color.Red("[ERROR] " + message)
	} else if typeOf == "info" {
		color.Blue("[INFO] " + message)
	}
}

func PrintPkgInf(res *additions.PackageData) {
	color.Cyan("[INFO] Found 1 package: %s", res.Name)
	color.Cyan("[INFO] Its description:")
	color.Yellow(res.Description)
}
