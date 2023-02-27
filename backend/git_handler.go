package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/google/uuid"
	"os"
)

type GitHandler interface {
	NewBranch() error
	AddAndCommitAll(msg string) error
	Push() error
}

type gitHandler struct {
	worktree   *git.Worktree
	repository *git.Repository
}

func (g gitHandler) NewBranch() error {
	return g.worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(uuid.New().String()),
		Create: true})
}

func (g gitHandler) AddAndCommitAll(msg string) error {
	err := g.worktree.AddWithOptions(&git.AddOptions{All: true})
	if err != nil {
		return err
	}
	_, err = g.worktree.Commit(msg, &git.CommitOptions{All: true})
	return err
}

func (g gitHandler) Push() error {
	return g.repository.Push(&git.PushOptions{})
}

func NewGitHandler(upstream, origin string) (GitHandler, error) {
	repository, err := git.PlainClone("./dog-api-images/", false, &git.CloneOptions{
		URL:        upstream,
		RemoteName: "upstream",
		Progress:   os.Stdout,
	})
	if err == git.ErrRepositoryAlreadyExists {
		repository, err = git.PlainOpen("./dog-api-images/")
	}
	if err != nil {
		return nil, err
	}

	_, err = repository.CreateRemote(&config.RemoteConfig{
		URLs: []string{origin},
		Name: "origin"})

	if err != nil {
		if err != git.ErrRemoteExists {
			return nil, err
		}
	}

	worktree, err := repository.Worktree()
	if err != nil {
		return nil, err
	}

	return gitHandler{worktree: worktree}, nil
}
