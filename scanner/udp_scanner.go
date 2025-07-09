package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
	"strings"
)

func ScanUDPPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("udp", address, time.Millisecond*900)
	if err != nil {
		return false
	}
	defer conn.Close()

	// Send dummy packet
	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return false
	}

	// Wait for response (if any)
	conn.SetReadDeadline(time.Now().Add(time.Millisecond * 500))
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	return err == nil
}

func ScanUdpPorts(host string, ports []int) {
	var wg sync.WaitGroup
	results := make(chan int, len(ports))

	fmt.Printf("\n[~] Scanning UDP ports on %s...\n", host)
	fmt.Printf("%-6s %-6s %-6s\n", "Proto", "Port", "State")
	fmt.Println(strings.Repeat("-", 30))

	for _, port := range ports {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			if ScanUDPPort(host, p) {
				results <- p
			} else {
				results <- -1
			}
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	openPorts := []int{}
for range ports {
	p := <-results
	if p != -1 {
		openPorts = append(openPorts, p)
		fmt.Printf("[+] %-6s %-6d %-6s\n", "UDP", p, "OPEN")
	}
}
if len(openPorts) == 0 {
	fmt.Println("[!] No open UDP ports found.")
} else {
	fmt.Printf("[✓] %d UDP port(s) open.\n", len(openPorts))
}

}
