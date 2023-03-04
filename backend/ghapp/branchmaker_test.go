package ghapp

import (
	"context"
	"errors"
	"testing"
)

type mockRef struct {
	str string
}

func (m mockRef) SHA() *string {
	return &m.str
}

func (m mockRef) Name() *string {
	return &m.str
}

type mockRemote struct {
	ref mockRef
}

func (m mockRemote) GetRefMain(_ context.Context) (ref, error) {
	return m.ref, nil
}

func (m mockRemote) CreateNewRef(_ context.Context, refMain ref) (ref, error) {
	if *refMain.Name() != "old" {
		return nil, errors.New("refMain is different than GetRefMain()")
	}
	return mockRef{str: "new"}, nil
}

func TestNewBranch(t *testing.T) {
	oldRef := "old"
	factory := branchMaker{remote: mockRemote{mockRef{oldRef}}}
	newRef, err := factory.make(context.Background())
	if err != nil {
		t.Fail()
	}
	if newRef == nil || *newRef.Name() != "new" {
		t.Fail()
	}
}
