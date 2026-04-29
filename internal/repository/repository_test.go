package repository

import "testing"

func TestParse(t *testing.T) {
	const owner string = "test-owner"
	parser := parser{
		usernameGetter: mockUsernameGetter(owner),
	}
	tests := []struct {
		name   string
		s      string
		want   Repository
	}{
		{
			name: "Only the repo name",
			s:    "my-repo",
			want: Repository{
				Host: "github.com",
				Owner: owner,
				Name: "my-repo",
			},
		},
		{
			name: "Owner and repo name",
			s:    "other-owner/my-repo",
			want: Repository{
				Host: "github.com",
				Owner: "other-owner",
				Name: "my-repo",
			},
		},
		{
			name: "With domain",
			s:    "example.com/other-owner/my-repo",
			want: Repository{
				Host: "example.com",
				Owner: "other-owner",
				Name: "my-repo",
			},
		},
		{
			name: "With HTTPS URL",
			s:    "https://example.com/other-owner/my-repo.git",
			want: Repository{
				Host: "example.com",
				Owner: "other-owner",
				Name: "my-repo",
			},
		},
		{
			name: "With SSH URL",
			s:    "git@example.com:other-owner/my-repo.git",
			want: Repository{
				Host: "example.com",
				Owner: "other-owner",
				Name: "my-repo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := parser.Parse(tt.s)
			if err != nil {
				t.Fatalf(`err = %v, want nil`, err)
			}
			if repo != tt.want {
				t.Fatalf(`repo = %#v, want %#v`, repo, tt.want)
			}
		})
	}
}

type mockUsernameGetter string

func (m mockUsernameGetter) Get() (string, error) {
	return string(m), nil
}
