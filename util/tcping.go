package util

import (
	"fmt"
	"net"
	"time"
)

func GetTcping(ipAddr string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ipAddr, port), time.Second*5)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
