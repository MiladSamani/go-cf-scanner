package util

import (
	"os"
	"sync"
)

func ScanWorker(jobs <-chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	for ip := range jobs {
		if GetTcping(ip, 443) {
			f, err := os.OpenFile("hit.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return
			}
			f.WriteString(ip + "\n")
		}
		// GetTcping(ip, 443)
	}
}
