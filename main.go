package main

import (
	"flag"
	"fmt"
	"github.com/a0s/httpproxy-go"
	"log"
	"net/http"
)

var (
	httpEnabled  bool
	httpsEnabled bool
	port         uint
	host         string
)

func init() {
	flag.BoolVar(&httpEnabled, "http", false, "enable http proxy")
	flag.BoolVar(&httpsEnabled, "https", false, "enable https proxy")
	flag.UintVar(&port, "port", 8080, "bind to port")
	flag.StringVar(&host, "host", "127.0.0.1", "bind to host")
}

func main() {
	flag.Parse()

	proxy, _ := httpproxy.NewProxy()
	proxy.OnConnect = OnConnect
	proxy.HttpEnabled = httpEnabled
	proxy.HttpsEnabled = httpsEnabled

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("Proxy listening on", addr)
	if err := http.ListenAndServe(addr, proxy); err != nil {
		panic(err)
	}
}

func OnConnect(_ *httpproxy.Context, host string) (ConnectAction httpproxy.ConnectAction, newHost string) {
	return httpproxy.ConnectProxy, host
}
