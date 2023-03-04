package ghapp

import "context"

type ref interface {
	SHA() *string
	Name() *string
}

type remote interface {
	GetRefMain(ctx context.Context) (ref, error)
	CreateNewRef(ctx context.Context, refMain ref) (ref, error)
}

type hub interface {
	PullRequest(ctx context.Context, refFrom ref, commitMessage string) (string, error)
}
