package ghapp

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/google/uuid"
)

type githubRemote struct {
	gitService   *github.GitService
	pullRequests *github.PullRequestsService
	repositories *github.RepositoriesService
	owner, repo  string
}

func (gh githubRemote) Commit(ctx context.Context, branch ref, file FileToCommit) error {
	commitMessage := file.CommitMessage()
	var req github.RepositoryContentFileOptions
	content, err := file.AsBytes()
	req.Content = content
	if err == nil {
		req.Message = &commitMessage
		req.Branch = branch.Name()
		req.Committer = &github.CommitAuthor{Name: github.String("GoFetchBot[bot]"), Email: github.String("126431462+GoFetchBot[bot]@users.noreply.github.com")}
		_, _, err = gh.repositories.CreateFile(ctx, gh.owner, gh.repo, file.Path(), &req)
	}
	return err
}

func (gh githubRemote) PullRequest(ctx context.Context, refFrom ref, commitMessage string) (string, error) {
	pullRequestLink := ""
	base := "main"
	newPullRequest := github.NewPullRequest{
		Title: &commitMessage,
		Head:  refFrom.Name(),
		Base:  &base,
		Body:  &commitMessage,
	}
	pr, _, err := gh.pullRequests.Create(ctx, gh.owner, gh.repo, &newPullRequest)
	if err == nil {
		pullRequestLink = *pr.HTMLURL
	}
	return pullRequestLink, err
}

type githubRef struct {
	ref *github.Reference
}

func (gh githubRef) SHA() *string {
	return gh.ref.Object.SHA
}

func (gh githubRef) Name() *string {
	return gh.ref.Ref
}

func (gh githubRemote) GetRefMain(ctx context.Context) (ref, error) {
	mainRef, _, err := gh.gitService.GetRef(ctx, gh.owner, gh.repo, "refs/heads/main")
	return githubRef{mainRef}, err
}

func (gh githubRemote) CreateNewRef(ctx context.Context, refFrom ref) (ref, error) {
	newName := fmt.Sprintf("refs/heads/%s", uuid.New())
	branchRef := &github.Reference{
		Ref: &newName,
		Object: &github.GitObject{
			SHA: refFrom.SHA(),
		},
	}
	newRef, _, err := gh.gitService.CreateRef(ctx, gh.owner, gh.repo, branchRef)
	return githubRef{newRef}, err
}
