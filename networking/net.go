package common

import (
	"errors"
	"net"
	"time"
)

// IsLocalIPAddress return an error when the IP is not a local IP
func IsLocalIPAddress(ip string) error {
	if ip == "0.0.0.0" {
		return nil
	}
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok {
				if ipnet.IP.String() == ip {
					return nil
				}
			}
		}
	} else {
		return err
	}

	return errors.New("Take care, the desired IP does not belong to this server")
}

//
// CheckEndpoint return an error when the endpoint is not reachable
func CheckEndpoint(network string, endpoint string) error {
	timeout := time.Duration(10) * time.Second
	conn, err := net.DialTimeout(network, endpoint, timeout)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
