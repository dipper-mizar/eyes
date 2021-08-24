package eyes

import (
	"net"
	"strings"
)

var ServerIP string

func init() {
	var ips []string

	if addresses, err := net.InterfaceAddrs(); err != nil {
		panic(err)
	} else {
		for _, address := range addresses {
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ips = append(ips, ipNet.IP.String())
				}
			}
		}
	}
	ServerIP = strings.Join(ips, "/")
}
