package main

import (
	"fmt"

	"github.com/ProArash/go-cf-scanner/pkg"
)

func main() {
	fmt.Println("cf scanner in go")
	ips := pkg.LoadIpList()
	fmt.Println(ips)
}
