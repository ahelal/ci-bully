package action

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/ahelal/ci-bully/pkg/config"
	"github.com/ahelal/ci-bully/pkg/output"
	"github.com/google/go-github/github"
)

func checkOpenPRs(ctx *context.Context, client *github.Client, owner string, repo string) (openPr []prType) {
	openPullRequests, _, err := client.PullRequests.List(*ctx, owner, repo, nil)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}

	for _, openPullRequest := range openPullRequests {
		duration := time.Since(*openPullRequest.CreatedAt)
		currentPr := prType{
			client:    client,
			ctx:       ctx,
			owner:     owner,
			repo:      repo,
			OpenSince: int(duration.Hours() / 24.0),
			GHpr:      openPullRequest,
		}
		openPr = append(openPr, currentPr)
	}
	return
}
func LoopOverRepos(runConfig *config.Config) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: runConfig.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	for _, repo := range runConfig.Repos {
		owner := strings.Split(repo, "/")[0]
		repo := strings.Split(repo, "/")[1]
		fmt.Printf("Checking %s/%s ...\n", owner, repo)
		checkOpenPRs(&ctx, client, owner, repo)
	}
}

func commentOnPr(pr prType, message string) {
	commentMsg := &github.IssueComment{Body: &message}
	_, _, err := pr.client.Issues.CreateComment(*pr.ctx, pr.owner, pr.repo, *pr.GHpr.Number, commentMsg)
	output.CheckError(fmt.Sprintf("Error failed to comment on %s/%s #%d", pr.owner, pr.repo, *pr.GHpr.Number), err)
}

func closePr(pr prType) {
	*pr.GHpr.State = "closed"
	_, _, err := pr.client.PullRequests.Edit(*pr.ctx, pr.owner, pr.repo, *pr.GHpr.Number, pr.GHpr)
	output.CheckError(fmt.Sprintf("Error failed to close PR %s/%s #%d", pr.owner, pr.repo, *pr.GHpr.Number), err)
}
