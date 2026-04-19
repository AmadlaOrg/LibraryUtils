package config

import (
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

type Config struct {
	// Auth credentials, if required, to use with the remote repository.
	Auth transport.AuthMethod

	// Name of the remote to be added, by default `origin`.
	RemoteName string

	// InsecureSkipTLS skips ssl verify if protocol is https
	InsecureSkipTLS *bool

	// CABundle specify additional ca bundle with system cert pool
	CABundle []byte

	// ProxyOptions provides info required for connecting to a proxy.
	ProxyOptions transport.ProxyOptions

	// Timeout specifies the timeout for everything Git
	Timeout int

	// CloneOptions describes how a clone should be performed.
	CloneOptions *CloneOptions
}

type CloneOptions struct {
	// Limit fetching to the specified number of commits.
	Depth int

	// RecurseSubmodules after the clone is created, initialize all submodules
	// within, using their default settings. This option is ignored if the
	// cloned repository does not have a worktree.
	RecurseSubmodules SubmoduleRescursivity

	// ShallowSubmodules limit cloning submodules to the 1 level of depth.
	// It matches the git command --shallow-submodules.
	ShallowSubmodules bool

	// Progress is where the human-readable information sent by the server is
	// stored, if nil nothing is stored and the capability (if supported)
	// no-progress, is sent to the server to avoid send this information.
	Progress sideband.Progress
}

type TagMode int

const (
	InvalidTagMode TagMode = iota
	// TagFollowing any tag that points into the histories being fetched is also
	// fetched. TagFollowing requires a server with `include-tag` capability
	// in order to fetch the annotated tags objects.
	TagFollowing
	// AllTags fetch all tags from the remote (i.e., fetch remote tags
	// refs/tags/* into local tags with the same name)
	AllTags
	// NoTags fetch no tags from the remote at all
	NoTags
)

// SubmoduleRescursivity defines how depth will affect any submodule recursive
// operation.
type SubmoduleRescursivity uint
