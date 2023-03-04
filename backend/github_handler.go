package main

import (
	"context"
	"flag"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	"github.com/google/uuid"
	"gofetch/postedimage"
	"log"
	"net/http"
)

type gitHubHandler struct {
	gitHandler GitHandler
	client     *github.Client
}

type GitHubHandler interface {
	PostToGithub(image postedimage.Image, commitMessage string) (string, error)
}

func (g gitHubHandler) makeGitBranchWithImage(image postedimage.Image) (string, error) {
	branch := uuid.New().String()
	err := g.gitHandler.NewBranch(branch)
	if err == nil {
		err = image.Save(g.gitHandler.GetPath())
	}
	return branch, err
}

func (g gitHubHandler) makePullRequestToMain(branch, commitMessage string) (string, error) {
	base := "main"
	newPullRequest := github.NewPullRequest{
		Title: &commitMessage,
		Head:  &branch,
		Base:  &base,
		Body:  &commitMessage,
	}
	pr, _, err := g.client.PullRequests.Create(context.Background(), "GoFetchBot", "dog-api-images", &newPullRequest)
	return *pr.HTMLURL, err
}

func (g gitHubHandler) PostToGithub(image postedimage.Image, commitMessage string) (string, error) {
	branch, err := g.makeGitBranchWithImage(image)
	if err != nil {
		return "", err
	}
	err = g.gitHandler.PushWithCommit(commitMessage)
	if err != nil {
		return "", err
	}

	return g.makePullRequestToMain(branch, commitMessage)
}

func NewGitHubHandler(path string) GitHubHandler {

	var appId int64
	var installationId int64
	var privateKey string

	flag.Int64Var(&appId, "appId", 0, "the app ID of the the bot to run")
	flag.Int64Var(&installationId, "installation ID", 0, "the installation ID of the the bot to run")
	flag.StringVar(&privateKey, "debug", "", "the private PEM encoded key of the bot to run")

	println("initializing git...")
	gh, err := ghinstallation.New(http.DefaultTransport, appId, installationId, []byte(privateKey))
	if err != nil {
		log.Fatalln(err)
	}
	token, err := gh.Token(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	git, err := NewGitHandler(path, Auth{username: "git", token: token})
	if err != nil {
		log.Fatalln(err)
	}
	println("git initialized.")

	return gitHubHandler{gitHandler: git, client: github.NewClient(&http.Client{Transport: gh})}
}
