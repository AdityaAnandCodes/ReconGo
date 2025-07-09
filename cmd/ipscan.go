package cmd

import (
	"fmt"

	"github.com/AdityaAnandCodes/ReconGo/network"
	"github.com/spf13/cobra"
)

var reverseDNS bool

// ipscanCmd represents the ipscan command
var ipscanCmd = &cobra.Command{
	Use:   "ipscan",
	Short: "Scan and display active local IP addresses",
	Long: `ipscan scans all active non-loopback network interfaces on your system 
and lists their associated IPv4 addresses. Optionally, it performs reverse DNS lookups.`,
	Run: func(cmd *cobra.Command, args []string) {
		ipInfos, err := network.GetLocalIps(reverseDNS)
		if err != nil {
			fmt.Printf("[!] Failed to retrieve IPs: %v\n", err)
			return
		}

		fmt.Println("[+] Active Network Interfaces and IP Addresses:")
		for _, ip := range ipInfos {
			if reverseDNS {
				fmt.Printf("    %s → %-15s | rDNS: %s\n", ip.InterfaceName, ip.IPAddress, ip.ReverseDNS)
			} else {
				fmt.Printf("    %s → %-15s\n", ip.InterfaceName, ip.IPAddress)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ipscanCmd)

	ipscanCmd.Flags().BoolVarP(&reverseDNS, "rdns", "r", false, "Perform reverse DNS lookup for each IP")
}
