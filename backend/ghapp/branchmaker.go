package ghapp

import (
	"context"
)

type branchMaker struct {
	remote remote
}

func (b branchMaker) make(ctx context.Context) (ref, error) {
	var newBranch ref
	refMain, err := b.remote.GetRefMain(ctx)
	if err == nil {
		newBranch, err = b.remote.CreateNewRef(ctx, refMain)
	}
	return newBranch, err
}
