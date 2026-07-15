package main

import (
	"log"
	"net/netip"
	"strings"
	"sync"

	"github.com/ProArash/go-cf-scanner/pkg"
	"github.com/ProArash/go-cf-scanner/util"
)

const (
	botsCount = 500
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string, 2000)
	ips := pkg.LoadIpList()

	ipList := strings.SplitSeq(string(ips), "\n")

	totalIpCount := 0

	log.Println("Extracting IP addresses....")

	for range botsCount {
		wg.Add(1)
		go util.ScanWorker(jobs, &wg)
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
