package util

import (
	"os"
	"sync"
)

func ScanWorker(jobs <-chan string, wg *sync.WaitGroup, timeout, method int) {
	defer wg.Done()
	for ip := range jobs {
		valid := false
		if method == 1 {
			valid = GetICMP(ip, timeout)
		} else {
			valid = GetTcping(ip, 443, timeout)
		}
		if valid {
			f, err := os.OpenFile("hit.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return
			}
			f.WriteString(ip + "\n")
		}
	}
}
