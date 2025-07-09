package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AdityaAnandCodes/ReconGo/finder"
	"github.com/AdityaAnandCodes/ReconGo/network"
	"github.com/AdityaAnandCodes/ReconGo/scanner"
)

func main() {

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

	
	scanner.ScanPorts("upload.facebook.com",commonTCPPorts ,true)

	// Example UDP ports to scan
	udpPorts := []int{53, 67, 68, 69, 123, 161, 500, 514, 520, 33434}
	start := time.Now()
	scanner.ScanUdpPorts("upload.facebook.com", udpPorts)
	elapsed := time.Since(start)
	fmt.Printf("UDP scan completed in %s\n", elapsed)


	finder.DiscoverSubdomains("facebook.com", "./helpers/subdomains.txt")


	ips, err := network.GetLocalIps()
	if err != nil {
		log.Fatalf("Failed to get local IPs: %v\n", err)
	}
	fmt.Println("Local IPs found:")
	for _, ip := range ips {
		fmt.Println("  -", ip)
	}
	

}