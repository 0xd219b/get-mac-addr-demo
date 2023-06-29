package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println(GetOutboundInterfaceInfo())
}

// GetOutboundInterfaceInfo Get preferred outbound ip of this machine
func GetOutboundInterfaceInfo() (ip, macAddr string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// get all network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	// get all ip addresses
	for _, i := range ifaces {
		faceAdders, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		// find the interface with the same ip address
		for _, ifaceAddr := range faceAdders {
			if strings.Contains(ifaceAddr.String(), localAddr.IP.String()) {
				return localAddr.IP.String(), i.HardwareAddr.String()
			}
		}
	}

	return localAddr.IP.String(), ""
}
