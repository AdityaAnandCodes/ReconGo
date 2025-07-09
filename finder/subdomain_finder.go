package finder

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func DiscoverSubdomains(baseDomain string, wordlistPath string) {
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Printf("[!] Failed to open wordlist: %s\n", err.Error())
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
		fmt.Println("[!] Wordlist is empty. Nothing to scan.")
		return
	}

	fmt.Printf("\n[~] Starting subdomain enumeration for: %s\n", baseDomain)
	fmt.Printf("%-40s %-20s\n", "Subdomain", "Status")
	fmt.Println(strings.Repeat("-", 60))

	concurrency := 100
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	for _, domain := range wordlist {
		wg.Add(1)
		sem <- struct{}{}
		go func(d string) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer func() {
				wg.Done()
				<-sem
				cancel()
			}()

			subdomain := fmt.Sprintf("%s.%s", d, baseDomain)
			ips, err := net.DefaultResolver.LookupHost(ctx, subdomain)
			if err == nil {
				fmt.Printf("[+] %-40s %v\n", subdomain, ips)
			} else if ctx.Err() == context.DeadlineExceeded {
				fmt.Printf("[!] %-40s Timeout\n", subdomain)
			} else {
				fmt.Printf("[-] %-40s Not Found\n", subdomain)
			}
		}(domain)
	}
	wg.Wait()
}
