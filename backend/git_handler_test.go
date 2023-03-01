package main

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage/memory"
	"testing"
)

func initGitHandler(storage *memory.Storage, fs billy.Filesystem) (gitHandler, *git.Repository, *git.Worktree) {
	repo, _ := git.Init(storage, fs)
	worktree, _ := repo.Worktree()
	worktree.Commit("initial commit", &git.CommitOptions{AllowEmptyCommits: true})
	return gitHandler{
		worktree:   worktree,
		repository: repo,
	}, repo, worktree
}

func getAmountOfBranches(branches storer.ReferenceIter) int {
	i := 0
	_ = branches.ForEach(func(branch *plumbing.Reference) error { i++; return nil })
	return i
}

func getAmountOfCommits(repository *git.Repository) int {
	headRef, _ := repository.Head()
	headCommit, _ := repository.CommitObject(headRef.Hash())
	return headCommit.NumParents() + 1
}

func TestGitHandlerMakesNewBranch(t *testing.T) {
	// given a git handler with 1 branch,
	handler, repo, _ := initGitHandler(memory.NewStorage(), memfs.New())
	branches, _ := repo.Branches()
	branchCount := getAmountOfBranches(branches)
	if branchCount != 1 {
		println("started with more than one branch")
		t.Fail()
	}
	// when a new branch is asked to be created,
	err := handler.NewBranch("test")

	// there are no errors,
	if err != nil {
		println(err.Error())
		t.Fail()
		return
	}
	// and there is now a new branch
	branches, _ = repo.Branches()
	branchCount = getAmountOfBranches(branches)
	if branchCount != 2 {
		println(branchCount)
		t.Fail()
	}
}

func TestGitHandlerAddsAllAndCommitsWithMessage(t *testing.T) {
	// given a git handler with 1 commit,
	fs := memfs.New()
	handler, repo, _ := initGitHandler(memory.NewStorage(), fs)
	if getAmountOfCommits(repo) != 1 {
		println("started with more than one commit")
		t.Fail()
	}
	// when a new file is made, and committed with a given commit message
	fs.Create("Readme.md")
	handler.addAndCommitAll("given commit message")

	// then a new commit is added,
	if getAmountOfCommits(repo) != 2 {
		println("commit did not increment")
		t.Fail()
	}
	headRef, _ := handler.repository.Head()
	headCommit, _ := repo.CommitObject(headRef.Hash())

	// and that commit has the given commit message
	if headCommit.Message != "given commit message" {
		println("message is not the same")
		t.Fail()
	}
	// and that commit contains the file
	parentCommit, _ := headCommit.Parent(0)
	patch, _ := parentCommit.Patch(headCommit)
	fromFile, toFile := patch.FilePatches()[0].Files()
	if fromFile != nil || toFile == nil || toFile.Path() != "Readme.md" {
		t.Fail()
	}
}

func TestRefSpec(t *testing.T) {
	// given a git handler with a branch named master,
	handler, _, _ := initGitHandler(memory.NewStorage(), memfs.New())
	// when asked for the refSpec
	refSpec, _ := handler.getRefSpec()
	// then is the expected result
	if refSpec != "+refs/heads/master:refs/heads/master" {
		t.Fail()
	}
	// and the expected result is valid
	if refSpec.Validate() != nil {
		t.Fail()
	}
}

type mockWorkTree struct {
	wasCalled chan bool
}

func (m mockWorkTree) Checkout(options *git.CheckoutOptions) error {
	return nil
}

func (m mockWorkTree) AddWithOptions(options *git.AddOptions) error {
	return nil
}

func (m mockWorkTree) Commit(msg string, options *git.CommitOptions) (plumbing.Hash, error) {
	close(m.wasCalled)
	return [20]byte{}, nil
}

type mockRepository struct {
	wasCalled chan bool
}

func (m mockRepository) Head() (*plumbing.Reference, error) {
	return &plumbing.Reference{}, nil
}

func (m mockRepository) Push(options *git.PushOptions) error {
	close(m.wasCalled)
	return nil
}

func TestGitHandler_PushWithCommit(t *testing.T) {
	// given a gitHandler with a mock worktree and repo
	mockTree := &mockWorkTree{make(chan bool)}
	mockRepo := &mockRepository{make(chan bool)}
	g := gitHandler{worktree: mockTree, repository: mockRepo}
	// when pushwithcommit is called
	g.PushWithCommit("")
	// then commit and push were called
	<-mockRepo.wasCalled
	<-mockTree.wasCalled
}
