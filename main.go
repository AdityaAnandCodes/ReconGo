package main

import (
	"github.com/AdityaAnandCodes/ReconGo/scanner"
)

func main() {
	scanner.ScanPortsRanging("scanme.nmap.org", "200-400")
}