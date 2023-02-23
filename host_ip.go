// The original code is from miguelmota/local_ip.go in Github Gist.
// It is modified by Taewhan Yoo in Feb. 20, 2023

package network

import (
	"errors"
	"net"
)

// LocalIP get the host machine local IP address
func HostIP(ip_block string) (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if isInBlock(ip, ip_block) {
				return ip, nil
			}
		}
	}

	return nil, errors.New("no IP")
}

func isInBlock(ip net.IP, ip_block string) bool {
	var iPBlocks []*net.IPNet
	for _, cidr := range []string{
		// don't check loopback ips
		//"127.0.0.0/8",    // IPv4 loopback
		//"::1/128",        // IPv6 loopback
		//"fe80::/10",      // IPv6 link-local
		//"10.0.0.0/8",     // RFC1918
		//"172.16.0.0/12",  // RFC1918
		//"192.168.0.0/16", // RFC1918
		ip_block,
	} {
		_, block, _ := net.ParseCIDR(cidr)
		iPBlocks = append(iPBlocks, block)
	}

	for _, block := range iPBlocks {
		if block.Contains(ip) {
			return true
		}
	}

	return false
}
