// Package repository is a helper to build repository data.
package repository

import (
	"strings"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spenserblack/gh-namespace-clone/internal/username"
)

// Repository is an alias for the type from go-gh.
type Repository = repository.Repository

// Parse parses with the default parser.
func Parse(s string) (Repository, error) {
	return Default().Parse(s)
}

// Parser is a type that tries to parse a repository.
type Parser interface {
	// Tries to parse a repository.
	Parse(s string) (Repository, error)
}

// parser is the concrete type.
type parser struct {
	// usernameGetter tries to get the username if it is not provided.
	usernameGetter username.Getter
}

// Default gets the default parser.
func Default() Parser {
	return parser{
		usernameGetter: username.DefaultGetter(),
	}
}

func (p parser) Parse(s string) (Repository, error) {
	s, err := p.prependOwner(s)
	if err != nil {
		return Repository{}, err
	}
	return repository.Parse(s)
}

// prependOwner will try to prepend the username to the repository format string
// only if it looks like the owner was not specified.
func (p parser) prependOwner(s string) (string, error) {
	if strings.ContainsRune(s, '/') {
		// NOTE Repository owner probably already specified.
		return s, nil
	}
	username, err := p.usernameGetter.Get()
	if err != nil {
		return "", err
	}
	return username + "/" + s, nil
}
