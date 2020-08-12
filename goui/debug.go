package goui

import (
	"fmt"
	"strings"

	"github.com/ttacon/chalk"
)

var (
	debug          = false
	warningMessage = []string{}
	errorMessages  = []string{}
)

func Debug(on bool) {
	debug = on
}

func debugRender(mainUI *gui, debugUI UI) {

}

const logPrefix = "[GOUI]"

func logInfo(str string) {
	fmt.Println(logPrefix + " " + str)
}

func logWarning(msg ...interface{}) {
	str := addSpaces(msg...)
	fmt.Println(chalk.Yellow.Color(logPrefix + " Warning: " + str))
	warningMessage = append(warningMessage, str)
}

func logError(msg ...interface{}) {
	str := addSpaces(msg...)
	fmt.Println(chalk.Red.Color(logPrefix + " Error: " + str))
	errorMessages = append(errorMessages, str)
}

func addSpaces(msg ...interface{}) string {
	str := fmt.Sprintln(msg)
	return strings.TrimSuffix(str, "\n")
}
