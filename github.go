package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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

const (
	HOURSINADAY = 24.0
)

func actions(currentPr prType) {
	var takeAction bool
	var actionTaken action

	for _, actionItem := range runConfig.Actions {
		switch {
		case currentPr.OpenSince == actionItem.Day:
			actionTaken = actionItem
			takeAction = true
			break
		case currentPr.OpenSince > actionItem.Day && actionItem.Last:
			// Last action
			actionTaken = actionItem
			takeAction = true
		}
	}
	if !takeAction {
		fmt.Printf("[skipping] \n")
		return
	}

	if arguments["--enable"].(bool) {
		commentOnPr(currentPr, actionTaken.Message)
		switch actionTaken.Action {
		case "close":
			closePr(currentPr)
		default:
			fmt.Printf("[%s]\n", actionTaken.Action)
		}
	} else {
		fmt.Printf("[%s 'DRY MODE']\n", actionTaken.Action)
	}

}

func checkOpenPRs(ctx *context.Context, client *github.Client, owner string, repo string) {
	openPullRequests, _, err := client.PullRequests.List(*ctx, owner, repo, nil)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}

	for _, openPullRequest := range openPullRequests {
		currentPr := prType{
			client:    client,
			ctx:       ctx,
			owner:     owner,
			repo:      repo,
			OpenSince: daysSincePRCreated(openPullRequest.CreatedAt),
			GHpr:      openPullRequest,
		}
		fmt.Printf("Repo %s/%s/#%d user '%s' open since '%d' days : ", owner, repo, *openPullRequest.Number, *openPullRequest.User.Login, currentPr.OpenSince)
		// take action if any
		actions(currentPr)
	}
}

func daysSincePRCreated(CreatedAt *time.Time) int {
	duration := time.Since(*CreatedAt)
	if !runConfig.OnlyWorkdays {
		return int(duration.Hours() / HOURSINADAY)
	}

	return workdaysBetweenDates(*CreatedAt, time.Now())
}

// workdaysBetweenDates calculates the workdays between two dates.
func workdaysBetweenDates(t1, t2 time.Time) int {
	if t2.Before(t1) {
		t1, t2 = t2, t1
	}

	days := 0
	for {
		if t1.After(t2) || t1.Equal(t2) {
			return days
		}
		if t1.Weekday() != time.Saturday && t1.Weekday() != time.Sunday {
			days++
		}

		t1 = t1.Add(time.Hour * HOURSINADAY)
	}
	return days
}

func loopOverRepos() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: runConfig.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	for _, repo := range runConfig.Repos {
		owner := strings.Split(repo, "/")[0]
		repo := strings.Split(repo, "/")[1]
		checkOpenPRs(&ctx, client, owner, repo)
	}
}

func commentOnPr(pr prType, message string) {
	message = strings.Replace(message, "_USER_", *pr.GHpr.User.Login, -1)
	message = strings.Replace(message, "_SINCE_", strconv.Itoa(pr.OpenSince), -1)
	message = strings.Replace(message, "_TILL_", strconv.Itoa(pr.CloseOn), -1)

	commentMsg := &github.IssueComment{Body: &message}
	_, _, err := pr.client.Issues.CreateComment(*pr.ctx, pr.owner, pr.repo, *pr.GHpr.Number, commentMsg)
	checkError(fmt.Sprintf("Error failed to comment on %s/%s #%d", pr.owner, pr.repo, *pr.GHpr.Number), err)
}

func closePr(pr prType) {
	*pr.GHpr.State = "closed"
	_, _, err := pr.client.PullRequests.Edit(*pr.ctx, pr.owner, pr.repo, *pr.GHpr.Number, pr.GHpr)
	checkError(fmt.Sprintf("Error failed to close PR %s/%s #%d", pr.owner, pr.repo, *pr.GHpr.Number), err)
}
