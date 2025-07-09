package scanner

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)


func ScanAndGrabBanner(host string, port int) {
	timeout := time.Millisecond * 900
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(1 * time.Second))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		fmt.Printf("[+] %-6s %-6d %-6s  %s\n", "TCP", port, "OPEN", "(no banner)")
		return
	}

	fmt.Printf("[+] %-6s %-6d %-6s  %q\n", "TCP", port, "OPEN", string(buf[:n]))
}



func ScanPort(host string, port int) bool {
	timeout := time.Millisecond * 900
	address := fmt.Sprintf("%s:%d",host,port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true	
}

func ScanPorts(host string, ports []int, withBanner bool) {
	var wg sync.WaitGroup
	results := make(chan int, len(ports))

	fmt.Printf("\n[~] Scanning TCP ports on %s...\n", host)
	fmt.Printf("%-6s %-6s %-6s  %s\n", "Proto", "Port", "State", "Banner")
	fmt.Println(strings.Repeat("-", 40))

	for _, port := range ports {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			if withBanner {
				ScanAndGrabBanner(host, p)
				results <- p
			} else {
				if ScanPort(host, p) {
					fmt.Printf("[+] %-6s %-6d %-6s\n", "TCP", p, "OPEN")
					results <- p
				} else {
					results <- -1
				}
			}
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	openFound := false
	for range ports {
		p := <-results
		if p != -1 {
			openFound = true
		}
	}
	if !openFound {
		fmt.Println("[!] No open TCP ports found.")
	}
}



func ScanPortsRanging(host string, ranging string, withBanner bool) {
	ports := strings.Split(ranging, "-")
	if len(ports) != 2 {
		fmt.Println("[!] Invalid range format. Use start-end (e.g. 20-100).")
		return
	}

	startPort, err1 := strconv.Atoi(ports[0])
	endPort, err2 := strconv.Atoi(ports[1])
	if err1 != nil || err2 != nil || startPort > endPort {
		fmt.Println("[!] Invalid port range values.")
		return
	}

	fmt.Printf("\n[~] Scanning TCP port range %d-%d on %s...\n", startPort, endPort, host)
	fmt.Printf("%-6s %-6s %-6s  %s\n", "Proto", "Port", "State", "Banner")
	fmt.Println(strings.Repeat("-", 40))

	var wg sync.WaitGroup
	results := make(chan int, endPort-startPort+1)
	sem := make(chan struct{}, 100)

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }()

			if withBanner {
				ScanAndGrabBanner(host, p)
				results <- p
			} else {
				if ScanPort(host, p) {
					fmt.Printf("[+] %-6s %-6d %-6s\n", "TCP", p, "OPEN")
					results <- p
				} else {
					results <- -1
				}
			}
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	openFound := false
	for p := range results {
		if p != -1 {
			openFound = true
		}
	}
	if !openFound {
		fmt.Println("[!] No open TCP ports found.")
	}
}



