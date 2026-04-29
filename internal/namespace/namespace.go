// Package namespace is for namespacing a clone path.
package namespace

import (
	"path/filepath"

	"github.com/spenserblack/gh-namespace-clone/internal/repository"
)

// Namespace builds a namespaced path for a clone.
type Namespace struct {
	// Prefix is the optional prefix path for clones.
	Prefix string
	// Repository is the repository used to build the namespace.
	Repository repository.Repository
	// UseDomain will cause the namespace to include the domain in the path if it is
	// true.
	UseDomain bool
}

// Path creates a path for cloning.
func (n Namespace) Path() string {
	segments := make([]string, 0, 4)
	if n.Prefix != "" {
		segments = append(segments, n.Prefix)
	}
	if n.UseDomain {
		segments = append(segments, n.Repository.Host)
	}
	segments = append(segments, n.Repository.Owner, n.Repository.Name)

	return filepath.Join(segments...)
}
