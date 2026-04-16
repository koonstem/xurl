// Package cmd provides the CLI commands for xurl.
// xurl is a command-line HTTP client for the X (Twitter) API.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"

	// flags
	hostname    string
	verbose     bool
	prettyPrint bool
)

// rootCmd is the base command for xurl
var rootCmd = &cobra.Command{
	Use:   "xurl <method> <path> [flags]",
	Short: "A command-line HTTP client for the X (Twitter) API",
	Long: `xurl is a command-line tool for making authenticated HTTP requests
to the X (Twitter) API. It handles OAuth authentication automatically
using credentials stored via the X CLI or environment variables.`,
	Example: `  xurl GET /2/tweets/1234567890
  xurl POST /2/tweets -f text="Hello, World!"
  xurl GET /2/users/me`,
	Args:         cobra.MinimumNArgs(2),
	RunE:         runRequest,
	SilenceUsage: true,
}

// Execute runs the root command and handles any errors.
func Execute(version string) {
	Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&hostname,
		"hostname", "H",
		"api.twitter.com",
		"X API hostname to send requests to",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose", "v",
		false,
		"Enable verbose output including request/response headers",
	)
	// Default pretty-print to true for better readability in my personal workflow.
	// Use -p=false or pipe to jq if raw output is needed.
	rootCmd.PersistentFlags().BoolVarP(
		&prettyPrint,
		"pretty", "p",
		true,
		"Pretty-print JSON responses",
	)

	// Add version command
	rootCmd.AddCommand(versionCmd)
}

// versionCmd prints the current version of xurl
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of xurl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("xurl version %s\n", Version)
	},
}

// runRequest is the main handler that dispatches HTTP requests to the X API.
// It is a placeholder that will be wired to the request package.
func runRequest(cmd *cobra.Command, args []string) error {
	method := args[0]
	path := args[1]

	// Validate HTTP method
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
		// valid
	default:
		return fmt.Errorf("unsupported HTTP method: %s", method)
	}

	_ = path // will be used by request handler

	// TODO: wire up authenticated HTTP client and request execution
	return fmt.Errorf("request execution not yet implemented for %s %s", method, path)
}
