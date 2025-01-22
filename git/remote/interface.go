package remote

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type IGoGitRemote interface {
	Config() *config.RemoteConfig
	String() string
	Push(o *git.PushOptions) error
	PushContext(ctx context.Context, o *git.PushOptions) (err error)
	FetchContext(ctx context.Context, o *git.FetchOptions) error
	Fetch(o *git.FetchOptions) error
	ListContext(ctx context.Context, o *git.ListOptions) (rfs []*plumbing.Reference, err error)
	List(o *git.ListOptions) (rfs []*plumbing.Reference, err error)
}
