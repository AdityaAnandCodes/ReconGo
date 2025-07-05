package scanner

import (
	"fmt"
	"net"
	"strings"
	"strconv"
	"sync"
)


func ScanPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d",host,port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}
	conn.Close()
	return true	
}


func ScanPorts(host string,ports []int){
	results := make(chan int)

	for _ ,port := range ports {
		go func(p int){
			if ScanPort(host,p){
				results <- p
			} else {
				results <- -1
			}
		}(port)
	}

	for range ports {
		p :=  <-results
		if p != -1 {
			fmt.Printf("Port :%d is open\n", p)
		}
	}
	fmt.Printf("No Ports Open")

}


func ScanPortsRanging(host string, ranging string) {
	ports := strings.Split(ranging, "-")
	if len(ports) != 2 {
		fmt.Println("Invalid range format. Use start-end.")
		return
	}

	startPort, err1 := strconv.Atoi(ports[0])
	endPort, err2 := strconv.Atoi(ports[1])
	if err1 != nil || err2 != nil || startPort > endPort {
		fmt.Println("Invalid port range values.")
		return
	}

	var wg sync.WaitGroup
	results := make(chan int, endPort-startPort+1)
	sem := make(chan struct{}, 100) // semaphore for max 100 concurrent goroutines

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		sem <- struct{}{} // acquire slot

		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }() // release slot

			if ScanPort(host, p) {
				results <- p
			} else {
				results <- -1
			}
		}(port)
	}

	// Close results channel after all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	openFound := false
	for p := range results {
		if p != -1 {
			fmt.Printf("Port :%d is open\n", p)
			openFound = true
		}
	}

	if !openFound {
		fmt.Println("No Ports Open")
	}
}