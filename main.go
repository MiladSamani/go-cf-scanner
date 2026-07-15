package main

import (
	"log"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ProArash/go-cf-scanner/pkg"
	"github.com/ProArash/go-cf-scanner/util"
)

const (
	defaultBots    = 300
	defaultTimeout = 1
	defaultMethod  = 2
)

func main() {
	bots := defaultBots
	timeout := defaultTimeout
	method := defaultMethod

	if len(os.Args) > 1 {
		if value, err := strconv.Atoi(os.Args[1]); err == nil && value > 0 {
			bots = value
		}
	}
	if len(os.Args) > 2 {
		if value, err := strconv.Atoi(os.Args[2]); err == nil && value > 0 {
			timeout = value
		}
	}
	if len(os.Args) > 3 {
		if value, err := strconv.Atoi(os.Args[3]); err == nil && (value == 1 || value == 2) {
			method = value
		}
	}

	var wg sync.WaitGroup
	jobs := make(chan string, 2000)
	ips := pkg.LoadIpList()

	ipList := strings.SplitSeq(string(ips), "\n")

	totalIpCount := 0

	log.Println("Extracting IP addresses....")

	for range bots {
		wg.Add(1)
		go util.ScanWorker(jobs, &wg, timeout, method)
	}
	// extract the ip range
	for cidr := range ipList {
		prefix, err := netip.ParsePrefix(cidr)
		if err != nil {
			log.Println(err)
		}
		ip := prefix.Masked().Addr()
		count := 1 << (32 - prefix.Bits())
		totalIpCount += count
		for range count {
			ip = ip.Next()
			jobs <- ip.String()
		}
	}
	close(jobs)
	wg.Wait()
	log.Printf("%v valid IP extracted", totalIpCount)
}
