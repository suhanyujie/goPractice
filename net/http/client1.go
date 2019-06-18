package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:3001/?m=api&c=message&a=setMessageSatus", strings.NewReader(fmt.Sprintf(`120
--69
Content-Disposition: form-data; name="token"

%s
--69
Content-Disposition: form-data; name="type"

1
--69--
0

`, "bc6c6f9088c850725ffe19eeefabcee8")))
	req.Header.Set("Test-Header-1", "chunked")
	req.Header.Set("Test-Header-2", "chunked")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=69")
	req.Header.Set("Transfer-Encoding", "chunked")
	req.Header.Set("Accept-Encoding", "chunked,gzip")
	req.TransferEncoding = []string{"chunked"}
	//proxyUrl, err := url.Parse("http://127.0.0.1:3001")
	// c := http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	fmt.Println(req.Header)
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(res.StatusCode)
}
