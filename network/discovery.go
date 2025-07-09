package network

import(
	"net"
	"fmt"
)

func GetLocalIps() ([]string, error){

	var ips []string
	
	ifaces,err := net.Interfaces()
	if err != nil {
		return ips,err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}

			ipv4 := ipnet.IP.To4()
			if ipv4 == nil {
				continue
			}

			ips = append(ips, ipv4.String())
		}
	}

	for _, ip := range ips{
		names, _ := net.LookupAddr(ip)
		fmt.Printf("Reverse DNS: %s\n", names)
	}

	return ips, nil

}
