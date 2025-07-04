package scanner

import (
	"fmt"
	"net"
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