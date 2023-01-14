package informer

import (
	"fmt"

	"github.com/fatih/color"
)

func Inform(typeOf string, message string, breakLine ...bool) {
	if breakLine == nil {
		message = message + " \n"
	}

	if typeOf == "error" {
		red := color.New(color.FgRed).SprintfFunc()
		fmt.Print(red("[HALT] " + message))

	} else if typeOf == "warn" {
		hiYellow := color.New(color.FgHiYellow).SprintFunc()
		fmt.Print(hiYellow("[WARN] " + message))
	} else if typeOf == "info" {
		blue := color.New(color.FgBlue).SprintFunc()
		fmt.Print(blue("[INFO] " + message))
	} else if typeOf == "good" {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Print(green("[GOOD] " + message))
	}
}
