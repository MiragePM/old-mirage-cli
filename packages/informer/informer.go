package informer

import "github.com/fatih/color"

func Informer(typeOf string, message string) {
	if typeOf == "error" {
		color.Red("[ERROR] " + message)
	} else if typeOf == "info" {
		color.Blue("[INFO] " + message)
	}
}
