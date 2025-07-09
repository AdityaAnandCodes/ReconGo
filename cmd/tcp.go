package cmd

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/AdityaAnandCodes/ReconGo/scanner"
	"github.com/spf13/cobra"
)

var (
	host      string
	portRange string
	banner    bool
	portList string 
)

var commonTCPPorts = []int{
    20,   // FTP (data)
    21,   // FTP (control)
    22,   // SSH
    23,   // Telnet
    25,   // SMTP
    53,   // DNS (TCP rarely used, mostly UDP)
    80,   // HTTP
    110,  // POP3
    111,  // RPCbind (Linux/Unix)
    135,  // Microsoft RPC
    139,  // NetBIOS Session
    143,  // IMAP
    161,  // SNMP
    179,  // BGP
    389,  // LDAP
    443,  // HTTPS
    445,  // Microsoft-DS (SMB)
    465,  // SMTPS
    514,  // Syslog (TCP)
    587,  // SMTP (submission)
    636,  // LDAPS
    873,  // rsync
    990,  // FTPS
    993,  // IMAPS
    995,  // POP3S
    1080, // SOCKS proxy
    1433, // Microsoft SQL Server
    1521, // Oracle DB
    1723, // PPTP VPN
    2049, // NFS
    2375, // Docker API (no TLS)
    2376, // Docker API (TLS)
    2483, // Oracle DB (default)
    2484, // Oracle DB (secure)
    3128, // Squid Proxy
    3306, // MySQL
    3389, // RDP (Remote Desktop)
    3690, // Subversion
    4000, // Common for custom apps
    5000, // Flask / uPnP
    5060, // SIP
    5432, // PostgreSQL
    5900, // VNC
    5985, // WinRM (HTTP)
    5986, // WinRM (HTTPS)
    6379, // Redis
    7001, // WebLogic
    8080, // HTTP-alt
    8443, // HTTPS-alt
    8888, // Dev servers
    9200, // Elasticsearch
    9300, // Elasticsearch internal
    11211,// Memcached
    27017,// MongoDB
    27018,// MongoDB alternate
}

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Scan TCP ports on a host",
	Long:  `Scan a host's TCP ports within a specified range. Optionally grab banners from open ports.`,
	Run: func(cmd *cobra.Command, args []string) {
		if host == "" {
			fmt.Println("Error: --host is required")
			cmd.Help()
			return
		}

		if portList == "" && portRange == "" {
			scanner.ScanPorts(host,commonTCPPorts,banner)
			return
		}
	
		// Handle specific ports
		if portList != "" {
			// Validate and parse portList
			var ports []int
			for _, p := range strings.Split(portList, ",") {
				num, err := strconv.Atoi(strings.TrimSpace(p))
				if err != nil {
					fmt.Printf("Invalid port: %s\n", p)
					return
				}
				ports = append(ports, num)
			}
	
			scanner.ScanPorts(host, ports, banner)
			return
		}
	
		// Handle port range
		if portRange != "" {
			scanner.ScanPortsRanging(host, portRange, banner)
			return
		}
	
		fmt.Println("Error: You must specify either --range or --ports")
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(tcpCmd)

	tcpCmd.Flags().StringVarP(&host, "host", "H", "", "Target host (e.g. 192.168.0.1)")
	tcpCmd.Flags().StringVarP(&portRange, "range", "r", "", "Port range (e.g. 20-80)")
	tcpCmd.Flags().BoolVarP(&banner, "banner", "b", false, "Enable banner grabbing")
	tcpCmd.Flags().StringVarP(&portList, "ports", "p", "", "Comma-separated list of ports (e.g. 22,80,443)")

}
