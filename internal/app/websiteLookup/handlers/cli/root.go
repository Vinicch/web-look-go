package cli

import (
	"websiteLookup/internal/app/websiteLookup/global"

	"github.com/spf13/cobra"
)

var (
	address string
	rootCmd = &cobra.Command{
		Use:     "wlc",
		Short:   "Let's you query IPs, CNAMEs, MX records and Name Servers!",
		Version: global.Version,
	}
)

// Execute will initiate the application
func Execute() error {
	return rootCmd.Execute()
}
