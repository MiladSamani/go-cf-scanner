package pkg

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func LoadIpList() []string {

	req, err := http.NewRequest("GET", "https://www.cloudflare.com/ips-v4", nil)
	req
	if err != nil {
		log.Fatalln(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(string(res), "\n")
}
