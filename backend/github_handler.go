package main

import (
	"github.com/google/uuid"
	"gofetch/postedimage"
)

type gitHubHandler struct {
	gitHandler GitHandler
}

type GitHubHandler interface {
	PostToGithub(image postedimage.Image) (string, error)
}

func (g gitHubHandler) makeGitBranchWithImage(image postedimage.Image) error {
	return g.gitHandler.NewBranch(uuid.New().String())
}

func (g gitHubHandler) PostToGithub(image postedimage.Image) (string, error) {
	err := g.makeGitBranchWithImage(image)
	if err != nil {
		return "", err
	}
	err = g.gitHandler.PushWithCommit("")
	if err != nil {
		return "", err
	}
	return "", nil
}

func (g gitHubHandler) AuthPullRequest() {
}
