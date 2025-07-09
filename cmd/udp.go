package cmd

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/AdityaAnandCodes/ReconGo/scanner"
	"github.com/spf13/cobra"
)

var (
	udpHost     string
	udpPortsStr string
)

var commonUDPPorts = []int{
	53,   // DNS
	67,   // DHCP server
	68,   // DHCP client
	69,   // TFTP
	123,  // NTP
	137,  // NetBIOS name service
	138,  // NetBIOS datagram
	161,  // SNMP
	162,  // SNMP traps
	500,  // ISAKMP / VPN
	514,  // Syslog
	520,  // RIP
	33434,// traceroute
}


var udpCmd = &cobra.Command{
	Use:   "udp",
	Short: "Scan UDP ports on a host",
	Long: `Scan specified UDP ports on a given host.

Example:
  ReconGo udp --host 192.168.0.1 --ports 53,123,161`,
  Run: func(cmd *cobra.Command, args []string) {
	if udpHost == "" {
		fmt.Println("[!] Error: --host is required")
		cmd.Help()
		return
	}

	// Use common UDP ports if none specified
	if udpPortsStr == "" {
		fmt.Println("[*] No ports specified. Scanning common UDP ports...")
		scanner.ScanUdpPorts(udpHost, commonUDPPorts)
		return
	}

	// Parse --ports
	var ports []int
	for _, p := range strings.Split(udpPortsStr, ",") {
		port, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			fmt.Printf("[!] Invalid port: %s\n", p)
			return
		}
		ports = append(ports, port)
	}

	scanner.ScanUdpPorts(udpHost, ports)
},
}

func init() {
	rootCmd.AddCommand(udpCmd)

	udpCmd.Flags().StringVarP(&udpHost, "host", "H", "", "Target host IP or domain")
	udpCmd.Flags().StringVarP(&udpPortsStr, "ports", "p", "", "Comma-separated UDP ports to scan (e.g. 53,161)")
}
