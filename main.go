package main

import (
	"flag"
	"fmt"
	"github.com/a0s/httpproxy-go"
	"log"
	"net/http"
	"strings"
)

var (
	httpEnabled  = flag.Bool("http", false, "enable http proxy")
	httpsEnabled = flag.Bool("https", false, "enable https proxy")
	port         = flag.Uint("port", 8080, "bind to port")
	host         = flag.String("host", "127.0.0.1", "bind to host")
)

func buildAddressString(host string, port uint) string {
	return fmt.Sprintf("%v:%v", host, port)
}

func buildStatusString(address string, httpEnabled bool, httpsEnabled bool) string {
	var protocols []string

	if httpEnabled == true {
		protocols = append(protocols, "http")
	}
	if httpsEnabled == true {
		protocols = append(protocols, "https")
	}
	if len(protocols) == 0 {
		protocols = append(protocols, "none")
	}

	withProtocols := strings.Join(protocols, ",")
	statusString := fmt.Sprintf("bind to %v with %v", address, withProtocols)

	return statusString
}

func main() {
	flag.Parse()

	address := buildAddressString(*host, *port)
	statusString := buildStatusString(address, *httpEnabled, *httpsEnabled)
	log.Printf(statusString)

	proxy, _ := httpproxy.NewProxy()
	proxy.OnConnect = OnConnect
	proxy.HttpEnabled = *httpEnabled
	proxy.HttpsEnabled = *httpsEnabled

	if err := http.ListenAndServe(address, proxy); err != nil {
		panic(err)
	}
}

func OnConnect(_ *httpproxy.Context, host string) (ConnectAction httpproxy.ConnectAction, newHost string) {
	return httpproxy.ConnectProxy, host
}
