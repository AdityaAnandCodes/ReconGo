package main

import (
	"github.com/AdityaAnandCodes/ReconGo/scanner"
)

func main() {
	// Commonly open TCP ports: 21 (FTP), 22 (SSH), 23 (Telnet), 25 (SMTP), 53 (DNS), 80 (HTTP), 110 (POP3), 143 (IMAP), 443 (HTTPS), 3306 (MySQL), 3389 (RDP), 5900 (VNC), 8080 (HTTP-alt)
	ports := []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 3306, 3389, 5900, 8080}
	scanner.ScanPorts("scanme.nmap.org", ports)
}