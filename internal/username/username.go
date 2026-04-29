// Package username provides utilities to get the current user.
package username

import "github.com/cli/go-gh/v2/pkg/api"

// Getter a type that tries to get a user.
type Getter interface {
	// Get tries to get a user's name.
	Get() (string, error)
}

// DefaultGetter returns the default Getter.
func DefaultGetter() Getter {
	return getter{}
}

// getter is the concrete type that gets a user's name from the API.
type getter struct{}

func (getter) Get() (string, error) {
	client, err := api.DefaultRESTClient()
	response := struct{ Login string }{}
	if err != nil {
		return "", err
	}
	err = client.Get("user", &response)
	if err != nil {
		return "", err
	}

	return response.Login, nil
}
