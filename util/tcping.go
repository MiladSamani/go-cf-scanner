package util

import (
	"fmt"
	"net"
	"time"
)

func GetTcping(ipAddr string, port, timeout int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ipAddr, port), time.Duration(timeout)*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
