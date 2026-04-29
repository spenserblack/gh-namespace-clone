package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// prefix is an optional prefix for the target clone path.
	prefix string

	// domain will include the repository's domain in the target clone path if true.
	domain bool
)

var rootCmd = &cobra.Command{
	Use:     "gh-namespace-clone [flags] [DOMAIN]/[OWNER]/<REPO>",
	Short:   "Clone a repository with automatic namespacing",
	Example: "gh-namespace-clone -P ~/Development cli/cli",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stderr := cmd.ErrOrStderr()
		stdout := cmd.OutOrStdout()

		repo := args[0]

		fmt.Fprintf(stdout, "prefix=%s domain=%t repo=%s\n", prefix, domain, repo)
		fmt.Fprintf(stderr, "TODO\n")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "P", ".", "Clone the namespaced repository under this path")
	rootCmd.PersistentFlags().BoolVarP(&domain, "domain", "D", false, "Include the domain in the namespace")
}
