package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Sleep for 2s to give Kamailio a chance to boot up.
	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "sip-proxy:8787")
	if err != nil {
		fmt.Fprintln(os.Stderr, "connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to EVAPI server")

	buf := make([]byte, 8000)
	total := 0
	for total < 8000 {
		n, err := conn.Read(buf[total:])
		total += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read ended after %d bytes: %v\n", total, err)
			break
		}
	}

	fmt.Printf("read %d bytes, now holding connection open without reading\n", total)

	// Block forever, keeping the connection open but never reading again.
	<-time.After(600 * time.Second)
}
