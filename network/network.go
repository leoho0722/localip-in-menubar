package network

import (
	"fmt"
	"log"
	"net"
)

// GetLocalNetworkIPv4Address 取得區網 IP
func GetLocalNetworkIPv4Address() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println("IP：", localAddr.IP)
	return localAddr.IP
}
