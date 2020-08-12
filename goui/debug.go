package goui

import (
	"fmt"

	"github.com/ttacon/chalk"
)

var (
	debug    = false
	warnings = []string{}
	errors   = []string{}
)

func Debug(on bool) {
	debug = on
}

func debugRender(mainUI *gui, debugUI UI) {

}

const logPrefix = "[GOUI]"

func logInfo(str string) {
	if !debug {
		return
	}
	fmt.Println(logPrefix + " " + str)
}

func logWarning(str string) {
	if !debug {
		return
	}
	fmt.Println(chalk.Yellow.Color(logPrefix + " Warning: " + str))
	warnings = append(warnings, str)
}

func logError(str string) {
	if !debug {
		return
	}
	fmt.Println(chalk.Red.Color(logPrefix + " Error: " + str))
	errors = append(errors, str)
}
