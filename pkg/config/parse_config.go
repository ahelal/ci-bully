package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/ahelal/ci-bully/pkg/output"
	yaml "gopkg.in/yaml.v2"
)

// ParseConfig parse our config file
func ParseConfig(filePath string, runConfig *Config) {
	var data []byte
	var err error

	data, err = ioutil.ReadFile(filePath)
	output.CheckError("Failed to read config file.", err)

	err = yaml.Unmarshal([]byte(data), runConfig)
	output.CheckError("Failed to parse config file.", err)

	adaptConfig(runConfig)
}

//Len is part of sort.Interface.
func (d ActionSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d ActionSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d ActionSlice) Less(i, j int) bool {
	return d[i].Day > d[j].Day
}

// Check syntax and required value and sort actions
func adaptConfig(runConfig *Config) {
	// Token section
	if runConfig.Token == "" {
		runConfig.Token = os.Getenv("GITHUB_TOKEN")
		if runConfig.Token == "" {
			output.CheckError("You need to define github in config or ENV 'GITHUB_TOKEN'", errors.New("token not defined"))
		}
	}

	// Repo section check
	if len(runConfig.Repos) == 0 {
		output.PrintError("You must define a repo section in your config")
	}
	// Sort by day
	sort.Sort(runConfig.Actions)

	for _, repository := range runConfig.Repos {
		repositoryValue := strings.Split(repository, "/")
		if len(repositoryValue) != 2 {
			output.PrintError(fmt.Sprintf("Repository %s is in wrong format", repository))
		}
	}
	if len(runConfig.Actions) == 0 {
		output.PrintError("No actions defined.")
	}

	for _, actionItem := range runConfig.Actions {
		if strings.ToLower(actionItem.Action) != "warn" && strings.ToLower(actionItem.Action) != "close" {
			output.PrintError(fmt.Sprintf("Unsupported action '%s'.", actionItem.Action))
		}
	}

}
