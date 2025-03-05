package utils

import (
	"net"
	"errors"
)

func GetAddr() (retAddr string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "nil",err
	}

	// 遍历网络接口，查找 eth0
	for _, iface := range interfaces {
		if iface.Name == "eth0" {
			addrs, err := iface.Addrs()
			if err != nil {
				return "nil", err
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "nil", errors.New("can not find eth0 address")
}
