package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type action struct {
	Day     int    `yaml:"day"`
	Action  string `yaml:"action"`
	Message string `yaml:"message"`
}

type actionSlice []action

//Len is part of sort.Interface.
func (d actionSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d actionSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d actionSlice) Less(i, j int) bool {
	return d[i].Day > d[j].Day
}

type options struct {
	Token   string      `yaml:"token"`
	Actions actionSlice `yaml:"actions"`
	Repos   []string    `yaml:"repos"`
}

func parseConfig(filePath string) {
	var data []byte
	var err error

	data, err = ioutil.ReadFile(filePath)
	checkError("Failed to read config file.", err)

	err = yaml.Unmarshal([]byte(data), &runConfig)
	checkError("Failed to parse config file.", err)

	// Check syntax and defaults
	configDefaults()
}

func configDefaults() {
	// Token section
	if runConfig.Token == "" {
		runConfig.Token = os.Getenv("GITHUB_TOKEN")
		if runConfig.Token == "" {
			checkError("You need to define github in config or ENV 'GITHUB_TOKEN'", errors.New("token not defined"))
		}
	}

	// Repo section check
	if len(runConfig.Repos) == 0 {
		printError("You must define a repo section in your config")
	}
	// Sort by day
	sort.Sort(runConfig.Actions)

	for _, repository := range runConfig.Repos {
		repositoryValue := strings.Split(repository, "/")
		if len(repositoryValue) != 2 {
			printError(fmt.Sprintf("Repository %s is in wrong format", repository))
		}
	}
	if len(runConfig.Actions) == 0 {
		printError("No actions defined.")
	}

	for _, actionItem := range runConfig.Actions {
		if strings.ToLower(actionItem.Action) != "warn" && strings.ToLower(actionItem.Action) != "close" {
			printError(fmt.Sprintf("Unsupported action '%s'.", actionItem.Action))
		}
	}

}
