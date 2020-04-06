package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// command-line flag parsing
	var d = flag.String("d", "", "`domain`: 域名 IP")
	var p = flag.Int("p", 443, "`port` : 端口")
	var t = flag.Int("t", 5, "`timeout` : 连接超时")
	var s = flag.Bool("s", false, "seconds since 1970-01-01 00:00:00 UTC")

	flag.Parse()

	if *d == "" {
		usage := "Usage: " + os.Args[0] + " -d domain [-p port] [-t timeout] [-s]\n"
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(2)
	}

	// define dialer with timeout
	duration := time.Duration(*t * 1000 * 1000 * 1000)
	var dialer net.Dialer
	dialer.Timeout = duration

	// dial with dialer
	addr := *d + ":" + strconv.Itoa(*p)
	conn, err := tls.DialWithDialer(&dialer, "tcp", addr, &tls.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		state := conn.ConnectionState()
		certs := state.PeerCertificates
		cert := *certs[0]
		notAfter := cert.NotAfter.Local()
		if *s {
			fmt.Println(notAfter.Unix())
		} else {
			fmt.Println(notAfter)
		}
		conn.Close()
	}
}
