package main

import (
	"github.com/google/uuid"
	"gofetch/postedimage"
)

type gitHubHandler struct {
	gitHandler GitHandler
}

type GitHubHandler interface {
	PostToGithub(image postedimage.Image, commitMessage string) (string, error)
}

func (g gitHubHandler) makeGitBranchWithImage(image postedimage.Image) error {
	err := g.gitHandler.NewBranch(uuid.New().String())
	if err == nil {
		err = image.Save(g.gitHandler.GetPath())
	}
	return err
}

func (g gitHubHandler) PostToGithub(image postedimage.Image, commitMessage string) (string, error) {
	err := g.makeGitBranchWithImage(image)
	if err != nil {
		return "", err
	}
	err = g.gitHandler.PushWithCommit(commitMessage)
	if err != nil {
		return "", err
	}
	return "", nil
}
