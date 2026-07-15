package pkg

import (
	"io"
	"log"
	"net/http"
)

func LoadIpList() []byte {

	req, err := http.NewRequest("GET", "https://www.cloudflare.com/ips-v4", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}
