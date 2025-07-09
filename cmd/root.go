package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ReconGo",
	Short: "A lightweight network reconnaissance CLI tool",
	Long: `ReconGo is a modular CLI tool for performing network reconnaissance tasks like:
- TCP and UDP port scanning
- Banner grabbing
- Subdomain enumeration
- Local IP discovery

Use 'ReconGo help [command]' for detailed usage.`,
	// Optional: display help if no subcommand is given
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'ReconGo help' to view available commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// You can add global flags here if needed
	// Example: config file support or logging verbosity

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ReconGo.yaml)")
}
