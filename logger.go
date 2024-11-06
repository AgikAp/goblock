package goblock

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"strings"
	"time"
)

var logo = `
  ________         __________.__                 __    
 /  _____/  ____   \______   \  |   ____   ____ |  | __
/   \  ___ /  _ \   |    |  _/  |  /  _ \_/ ___\|  |/ /
\    \_\  (  <_> )  |    |   \  |_(  <_> )  \___|    < 
 \______  /\____/   |______  /____/\____/ \___  >__|_ \
        \/                 \/                 \/     \/`

func printLogo() {
	fmt.Printf("%s\n\n", logo)
}

func logRequest(method string, path string, statusCode int, duration time.Duration) {
	yellow := color.New(color.FgYellow).SprintFunc()

	totalLength := 100
	formattedLog := fmt.Sprintf("%s %s", yellow(method), path)

	dots := totalLength - len(formattedLog) - len(duration.String())
	if dots < 0 {
		dots = 0
	}

	log.Printf("%s %s %s %s", formattedLog, repeatString(".", dots), writeStatusCode(statusCode), yellow(duration))
}

func writeStatusCode(statusCode int) string {
	var colorCode func(a ...interface{}) string

	if statusCode >= 400 && statusCode < 599 {
		colorCode = color.New(color.FgRed).SprintFunc()
	} else {
		colorCode = color.New(color.FgGreen).SprintFunc()
	}

	return colorCode(fmt.Sprintf("%d", statusCode))
}

func repeatString(s string, count int) string {
	return strings.Repeat(s, count)
}
