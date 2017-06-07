package main

import (
	"github.com/ahelal/ci-bully/pkg/action"
	"github.com/ahelal/ci-bully/pkg/config"

	docopt "github.com/docopt/docopt-go"
)

var runConfig config.Config
var arguments map[string]interface{}

func main() {
	usage := `Usage: ci-bully [-c FILE] [-evhs]
Push for PRs to be merged or closed.

Options:
  -c FILE --config=FILE             Config the yaml file [default: config.yml].
  -e --enable                       Take action by default it runs in drymode [default: false].
  -h --help                         Show this screen.
  -v --version                      Show version.`
	arguments, _ = docopt.Parse(usage, nil, true, "", false)
	config.ParseConfig(arguments["--config"].(string), &runConfig)
	action.loopOverRepos(&runConfig)
}
