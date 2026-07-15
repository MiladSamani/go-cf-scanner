package util

import (
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func GetICMP(ipAddr string, timeout int) bool {
	c, err := icmp.ListenPacket("ip4:icmp", "")
	if err != nil {
		return false
	}
	defer c.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{ID: os.Getpid() & 0xffff, Seq: 1, Data: []byte("ping")},
	}
	data, err := msg.Marshal(nil)
	if err != nil {
		return false
	}

	c.WriteTo(data, &net.IPAddr{IP: net.ParseIP(ipAddr)})
	// timeout
	c.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))

	replyBtye := make([]byte, 1500)
	reply, _, err := c.ReadFrom(replyBtye)
	if err != nil {
		return false
	}
	rm, err := icmp.ParseMessage(1, replyBtye[:reply])
	if err != nil {
		return false
	}

	if rm.Type == ipv4.ICMPTypeEchoReply {
		return true
	} else {
		return false
	}
}
