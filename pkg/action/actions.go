package action

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ahelal/ci-bully/pkg/config"
	"github.com/google/go-github/github"
)

type prType struct {
	client    *github.Client
	ctx       *context.Context
	owner     string
	repo      string
	OpenSince int
	CloseOn   int
	GHpr      *github.PullRequest
}

func actions(currentPr prType, runConfig config.Config, enabled bool) {
	var takeAction bool
	var actionTaken config.Action

	for _, actionItem := range runConfig.Actions {
		if currentPr.OpenSince == actionItem.Day {
			actionTaken = actionItem
			takeAction = true
			break
		}
	}
	if !takeAction {
		fmt.Printf("[skipping] \n")
		return
	}
	message := constructMsg(currentPr, actionTaken.Message)
	if enabled {
		commentOnPr(currentPr, message)
		switch actionTaken.Action {
		case "close":
			closePr(currentPr)
		default:
			fmt.Printf("[%s]\n", actionTaken.Action)
		}
	} else {
		fmt.Printf("[%s 'DRY MODE']\n", actionTaken.Action)
		fmt.Println(message)
	}

}

// fmt.Printf("Repo %s/%s/#%d user '%s' open since '%d' days : ", owner, repo, *openPullRequest.Number, *openPullRequest.User.Login, currentPr.OpenSince)
// // take action if any
// actions(currentPr)

func constructMsg(pr prType, message string) string {
	message = strings.Replace(message, "_USER_", *pr.GHpr.User.Login, -1)
	message = strings.Replace(message, "_SINCE_", strconv.Itoa(pr.OpenSince), -1)
	message = strings.Replace(message, "_TILL_", strconv.Itoa(pr.CloseOn), -1)
	return message
}
