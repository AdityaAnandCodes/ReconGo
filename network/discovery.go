package network

import (
	"net"
)

type IPInfo struct {
	InterfaceName string
	IPAddress     string
	ReverseDNS    string
}

func GetLocalIps(withReverseDNS bool) ([]IPInfo, error) {
	var results []IPInfo

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
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

			ip := ipnet.IP.To4()
			if ip == nil {
				continue
			}

			ipStr := ip.String()
			ipInfo := IPInfo{
				InterfaceName: iface.Name,
				IPAddress:     ipStr,
			}

			if withReverseDNS {
				names, err := net.LookupAddr(ipStr)
				if err == nil && len(names) > 0 {
					ipInfo.ReverseDNS = names[0]
				} else {
					ipInfo.ReverseDNS = "N/A"
				}
			}

			results = append(results, ipInfo)
		}
	}

	return results, nil
}
