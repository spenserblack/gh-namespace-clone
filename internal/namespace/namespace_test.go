package namespace

import (
	"testing"

	"github.com/spenserblack/gh-namespace-clone/internal/repository"
)

func TestPath(t *testing.T) {
	tests := []struct {
		name      string
		namespace Namespace
		want      string
	}{
		{
			name: "Default namespace",
			namespace: Namespace{
				Prefix: "",
				Repository: repository.Repository{
					Host:  "example.com",
					Owner: "my-name",
					Name:  "my-repo",
				},
				UseDomain: false,
			},
			want: "my-name/my-repo",
		},
		{
			name: "With prefix",
			namespace: Namespace{
				Prefix: "/home/me/Development",
				Repository: repository.Repository{
					Host:  "example.com",
					Owner: "my-name",
					Name:  "my-repo",
				},
				UseDomain: false,
			},
			want: "/home/me/Development/my-name/my-repo",
		},
		{
			name: "With domain",
			namespace: Namespace{
				Prefix: "",
				Repository: repository.Repository{
					Host:  "example.com",
					Owner: "my-name",
					Name:  "my-repo",
				},
				UseDomain: true,
			},
			want: "example.com/my-name/my-repo",
		},
		{
			name: "With prefix and domain",
			namespace: Namespace{
				Prefix: "/home/me/Development/",
				Repository: repository.Repository{
					Host:  "example.com",
					Owner: "my-name",
					Name:  "my-repo",
				},
				UseDomain: true,
			},
			want: "/home/me/Development/example.com/my-name/my-repo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.namespace.Path(); got != tt.want {
				t.Fatalf(`Path() = %s, want %s`, got, tt.want)
			}
		})
	}
}
