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

	conc := 100
	var wg sync.WaitGroup
	sem := make(chan struct{}, conc)
	for _, domain := range wordlist {
		wg.Add(1)
		sem <- struct{}{}
		go func(d string){
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer func(){
				wg.Done()
				<- sem
				cancel()
			}()
			subdomain := fmt.Sprintf("%s.%s", d, baseDomain )
			ips, err := net.DefaultResolver.LookupHost(ctx, subdomain)
			if err == nil {
				fmt.Printf("Found: %s → %v\n", subdomain, ips)
			} else if ctx.Err() == context.DeadlineExceeded {
				fmt.Printf("Timeout: %s\n", subdomain)
			}
			
		}(domain)	
	}
	wg.Wait()
}