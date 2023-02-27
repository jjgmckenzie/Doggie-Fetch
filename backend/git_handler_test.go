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

func initGitHandler(fs billy.Filesystem) gitHandler {
	repository, _ := git.Init(memory.NewStorage(), fs)
	worktree, _ := repository.Worktree()
	worktree.Commit("initial commit", &git.CommitOptions{AllowEmptyCommits: true})
	return gitHandler{
		worktree:   worktree,
		repository: repository,
	}
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
	handler := initGitHandler(memfs.New())
	branches, _ := handler.repository.Branches()
	branchCount := getAmountOfBranches(branches)
	if branchCount != 1 {
		println("started with more than one branch")
		t.Fail()
	}
	// when a new branch is asked to be created,
	err := handler.NewBranch()

	// there are no errors,
	if err != nil {
		println(err.Error())
		t.Fail()
		return
	}
	// and there is now a new branch
	branches, _ = handler.repository.Branches()
	branchCount = getAmountOfBranches(branches)
	if branchCount != 2 {
		println(branchCount)
		t.Fail()
	}
}

func TestGitHandlerAddsAllAndCommitsWithMessage(t *testing.T) {
	// given a git handler with 1 commit,
	fs := memfs.New()
	handler := initGitHandler(fs)
	if getAmountOfCommits(handler.repository) != 1 {
		println("started with more than one commit")
		t.Fail()
	}
	// when a new file is made, and committed with a given commit message
	fs.Create("Readme.md")
	handler.AddAndCommitAll("given commit message")

	// then a new commit is added,
	if getAmountOfCommits(handler.repository) != 2 {
		println("commit did not increment")
		t.Fail()
	}
	headRef, _ := handler.repository.Head()
	headCommit, _ := handler.repository.CommitObject(headRef.Hash())

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
