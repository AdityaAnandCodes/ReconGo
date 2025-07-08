package finder

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func DiscoverSubdomains(baseDomain string, wordlistpath string){
	file,err := os.Open(wordlistpath)
	if err != nil {
		fmt.Println("Error Fetching The Wordlist")
		return
	}
	defer file.Close()
	var wordlist []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			wordlist = append(wordlist, word)
		}	
	}
	if len(wordlist) == 0 {
		fmt.Println("Wordlist is Empty")
		return
	}

	for _, domain := range wordlist {
		subdomain := fmt.Sprintf("%s.%s", domain, baseDomain )
		ips, err := net.LookupHost(subdomain)
		if err == nil {
			fmt.Printf("Found: %s → %v\n", subdomain , ips)
		}
	}
}