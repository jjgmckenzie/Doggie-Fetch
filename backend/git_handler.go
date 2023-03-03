package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"time"
)

type Auth struct {
	username, token string
}

type GitHandler interface {
	NewBranch(name string) error
	PushWithCommit(msg string) error
	GetPath() string
}

type workTree interface {
	Checkout(options *git.CheckoutOptions) error
	AddWithOptions(options *git.AddOptions) error
	Commit(msg string, options *git.CommitOptions) (plumbing.Hash, error)
}

type repository interface {
	Head() (*plumbing.Reference, error)
	Push(options *git.PushOptions) error
}

type gitHandler struct {
	worktree   workTree
	repository repository
	auth       Auth
	path       string
}

func (g gitHandler) GetPath() string {
	return g.path
}

func (g gitHandler) NewBranch(name string) error {
	return g.worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(name),
		Create: true})
}

func (g gitHandler) addAll() error {
	return g.worktree.AddWithOptions(&git.AddOptions{All: true})
}

func (g gitHandler) commitAll(msg string) error {
	_, err := g.worktree.Commit(msg, &git.CommitOptions{
		All: true,
		Author: &object.Signature{Name: "GoFetchBot[bot]",
			Email: "126431462+GoFetchBot[bot]@users.noreply.github.com",
			When:  time.Now()},
	})
	return err
}

func (g gitHandler) getRefSpec() (config.RefSpec, error) {
	local, err := g.repository.Head()
	refSpec := config.RefSpec("")
	if err == nil {
		refSpec = config.RefSpec("+" + local.Name() + ":" + local.Name())
	}
	return refSpec, err
}

func (g gitHandler) push() error {
	refSpec, err := g.getRefSpec()
	if err == nil {
		err = g.repository.Push(&git.PushOptions{
			RefSpecs: []config.RefSpec{refSpec},
			Auth:     &http.BasicAuth{Username: g.auth.username, Password: g.auth.token},
			Force:    true,
		})
	}
	return err
}

func (g gitHandler) addAndCommitAll(msg string) error {
	err := g.addAll()
	if err == nil {
		err = g.commitAll(msg)
	}
	return err
}

func (g gitHandler) PushWithCommit(msg string) error {
	err := g.addAndCommitAll(msg)
	if err == nil {
		err = g.push()
	}
	return err
}

func NewGitHandler(path string, auth Auth) (GitHandler, error) {

	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	return gitHandler{repository: repo, worktree: worktree, auth: auth, path: path}, err
}
