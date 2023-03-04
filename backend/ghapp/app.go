package ghapp

import (
	"context"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	"net/http"
)

type FileToCommit interface {
	Path() string
	CommitMessage() string
	AsBytes() ([]byte, error)
}

type fileCommitter interface {
	Commit(ctx context.Context, branch ref, file FileToCommit) error
}

type app struct {
	hub           hub
	fileCommitter fileCommitter
	branchMaker   branchMaker
}

func (a app) commitAndPull(ctx context.Context, file FileToCommit, branch ref) (string, error) {
	link := ""
	commitMessage := file.CommitMessage()
	err := a.fileCommitter.Commit(ctx, branch, file)
	if err == nil {
		link, err = a.hub.PullRequest(ctx, branch, commitMessage)
	}
	return link, err
}

func (a app) MakePullRequest(ctx context.Context, file FileToCommit) (string, error) {
	link := ""
	branch, err := a.branchMaker.make(ctx)
	if err == nil {
		link, err = a.commitAndPull(ctx, file, branch)
	}
	return link, err
}

type App interface {
	MakePullRequest(ctx context.Context, file FileToCommit) (string, error)
}

func New(transport *ghinstallation.Transport, repository string) App {

	client := github.NewClient(&http.Client{Transport: transport})
	ghRemote := githubRemote{
		gitService:   client.Git,
		pullRequests: client.PullRequests,
		repositories: client.Repositories,
		owner:        "gofetchbot",
		repo:         repository,
	}
	return app{
		hub:           ghRemote,
		fileCommitter: ghRemote,
		branchMaker: branchMaker{
			remote: ghRemote,
		},
	}
}
