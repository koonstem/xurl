// xurl is a command-line HTTP client for the X (Twitter) API.
// It is a fork of xdevplatform/xurl with additional features and improvements.
package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xurl",
	Short: "A command-line HTTP client for the X (Twitter) API",
	Long: `xurl is a command-line tool that makes authenticated requests
to the X (Twitter) API. It handles OAuth 1.0a and OAuth 2.0 authentication
automatically, allowing you to quickly interact with the API from your terminal.

Example:
  xurl GET /2/tweets/:id
  xurl POST /2/tweets -d '{"text": "Hello, World!"}'
  xurl GET /2/users/me
  xurl GET /2/users/:id/tweets
  xurl GET /2/users/:id/followers
  xurl GET /2/users/:id/following`,
	SilenceUsage: true,
	// CompletionOptions hides the default 'completion' subcommand from help output
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
