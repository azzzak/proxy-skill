package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var (
	proxy   *httputil.ReverseProxy
	version string
)

func main() {
	forwardTo := os.Getenv("FORWARD_TO")
	if forwardTo == "" {
		fmt.Fprintln(os.Stderr, "You must set FORWARD_TO environment variable")
		os.Exit(1)
	}

	fmt.Printf("proxy-skill %s\n", version)
	fmt.Printf("Forwarding to: %s\n", forwardTo)

	http.HandleFunc("/", handle(forwardTo))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handle(forwardTo string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		url, _ := url.Parse(forwardTo)

		proxy = httputil.NewSingleHostReverseProxy(url)
		proxy.Director = func(req *http.Request) {
			req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
			req.URL.Host = url.Host
			req.URL.Scheme = url.Scheme
			req.Host = url.Host
			req.URL.Path = url.Path
		}

		proxy.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 3 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 3 * time.Second,
		}

		proxy.ServeHTTP(res, req)
	}
}
