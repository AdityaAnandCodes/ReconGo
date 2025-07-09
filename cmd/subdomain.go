package cmd

import (
	"fmt"

	"github.com/AdityaAnandCodes/ReconGo/finder"
	"github.com/spf13/cobra"
)

var (
	baseDomain   string
	wordlistPath string
)

var subdomainCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "Enumerate subdomains using a wordlist",
	Long: `Perform passive subdomain enumeration by brute-forcing using a provided wordlist.

Example:
  ReconGo subdomain --domain example.com --wordlist ./helpers/subdomains.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		if baseDomain == "" || wordlistPath == "" {
			fmt.Println("[!] Both --domain and --wordlist are required.")
			cmd.Help()
			return
		}

		fmt.Printf("[~] Starting subdomain enumeration for: %s\n", baseDomain)
		finder.DiscoverSubdomains(baseDomain, wordlistPath)
	},
}

func init() {
	rootCmd.AddCommand(subdomainCmd)

	subdomainCmd.Flags().StringVarP(&baseDomain, "domain", "d", "", "Base domain to scan (e.g. example.com)")
	subdomainCmd.Flags().StringVarP(&wordlistPath, "wordlist", "w", "", "Path to subdomain wordlist file")
}
