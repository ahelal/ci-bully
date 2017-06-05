package main

import (
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
)

var runConfig options
var arguments map[string]interface{}

func checkError(msg string, e error) {
	if e != nil {
		fmt.Printf("%s. %s\n", msg, e.Error())
		os.Exit(1)
	}
}

func printError(msg string) {
	fmt.Printf("%s\n", msg)
	os.Exit(1)
}

func main() {
	usage := `Usage: cibully [-c FILE] [-evhs]
Push for PRs to be merged or closed.

Options:
  -c FILE --config=FILE             Config the yaml file [default: config.yml].
  -e --enable                       Take action by default it runs in drymode [default: false].
  -h --help                         Show this screen.
  -v --version                      Show version.`
	arguments, _ = docopt.Parse(usage, nil, true, "", false)
	parseConfig(arguments["--config"].(string))
	loopOverRepos()
}
