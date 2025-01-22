package git

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type IGoGitRepository interface {
	Grep(opts *git.GrepOptions) ([]git.GrepResult, error)
	DeleteObject(hash plumbing.Hash) error
	Prune(opt git.PruneOptions) error
	Config() (*config.Config, error)
	SetConfig(cfg *config.Config) error
	ConfigScoped(scope config.Scope) (*config.Config, error)
	Remote(name string) (*git.Remote, error)
	Remotes() ([]*git.Remote, error)
	CreateRemote(c *config.RemoteConfig) (*git.Remote, error)
	CreateRemoteAnonymous(c *config.RemoteConfig) (*git.Remote, error)
	DeleteRemote(name string) error
	Branch(name string) (*config.Branch, error)
	CreateBranch(c *config.Branch) error
	DeleteBranch(name string) error
	CreateTag(name string, hash plumbing.Hash, opts *git.CreateTagOptions) (*plumbing.Reference, error)
	Tag(name string) (*plumbing.Reference, error)
	DeleteTag(name string) error
	Fetch(o *git.FetchOptions) error
	FetchContext(ctx context.Context, o *git.FetchOptions) error
	Push(o *git.PushOptions) error
	PushContext(ctx context.Context, o *git.PushOptions) error
	Log(o *git.LogOptions) (object.CommitIter, error)
	Tags() (storer.ReferenceIter, error)
	Branches() (storer.ReferenceIter, error)
	Notes() (storer.ReferenceIter, error)
	TreeObject(h plumbing.Hash) (*object.Tree, error)
	TreeObjects() (*object.TreeIter, error)
	CommitObject(h plumbing.Hash) (*object.Commit, error)
	CommitObjects() (object.CommitIter, error)
	BlobObject(h plumbing.Hash) (*object.Blob, error)
	BlobObjects() (*object.BlobIter, error)
	TagObject(h plumbing.Hash) (*object.Tag, error)
	TagObjects() (*object.TagIter, error)
	Object(t plumbing.ObjectType, h plumbing.Hash) (object.Object, error)
	Objects() (*object.ObjectIter, error)
	Head() (*plumbing.Reference, error)
	Reference(name plumbing.ReferenceName, resolved bool) (*plumbing.Reference, error)
	References() (storer.ReferenceIter, error)
	Worktree() (*git.Worktree, error)
	ResolveRevision(in plumbing.Revision) (*plumbing.Hash, error)
	RepackObjects(cfg *git.RepackConfig) (err error)
	Merge(ref plumbing.Reference, opts git.MergeOptions) error
}

type IGoGitWorktree interface {
	Pull(o *git.PullOptions) error
	PullContext(ctx context.Context, o *git.PullOptions) error
	Checkout(opts *git.CheckoutOptions) error
	ResetSparsely(opts *git.ResetOptions, dirs []string) error
	Restore(o *git.RestoreOptions) error
	Reset(opts *git.ResetOptions) error
	Submodule(name string) (*git.Submodule, error)
	Submodules() (git.Submodules, error)
	Clean(opts *git.CleanOptions) error
	Grep(opts *git.GrepOptions) ([]git.GrepResult, error)
	Status() (git.Status, error)
	StatusWithOptions(o git.StatusOptions) (git.Status, error)
	Add(path string) (plumbing.Hash, error)
	AddWithOptions(opts *git.AddOptions) error
	AddGlob(pattern string) error
	Remove(path string) (plumbing.Hash, error)
	RemoveGlob(pattern string) error
	Move(from, to string) (plumbing.Hash, error)
	Commit(msg string, opts *git.CommitOptions) (plumbing.Hash, error)
}
