package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
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


func ScanUdpPorts(host string, ports []int){
	var wg sync.WaitGroup
	results := make(chan int, len(ports))

	for _, port := range ports {
		wg.Add(1)
		go func(p int){
			defer wg.Done()
			if ScanUDPPort(host,p){
				results <- p
			} else {
				results <- -1
			}
		}(port)
		}

	go func(){
		wg.Wait()
		close(results)
	}()

	openFound :=  false
	for range ports {
		p := <- results
		if p != -1 {
			openFound = true
			fmt.Printf("UDP Port: %d Found Open\n",p)
		}
	}

	if !openFound {
		fmt.Print("No Open Ports Were Found\n")
	}
}
