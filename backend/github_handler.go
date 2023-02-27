package main

import (
	"gofetch/postedimage"
)

type gitHubHandler struct {
	gitHandler GitHandler
}

type GitHubHandler interface {
	PostToGithub(image postedimage.Image) (string, error)
}

func (g gitHubHandler) makeGitBranchWithImage(image postedimage.Image) error {
	return g.gitHandler.NewBranch()
}

func (g gitHubHandler) PostToGithub(image postedimage.Image) (string, error) {
	err := g.makeGitBranchWithImage(image)
	if err != nil {
		return "", err
	}
	err = g.gitHandler.AddAndCommitAll("")
	if err != nil {
		return "", err
	}
	err = g.gitHandler.Push()
	if err != nil {
		return "", err
	}
	return "", nil
}
