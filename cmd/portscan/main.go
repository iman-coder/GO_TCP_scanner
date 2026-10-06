package main

import (
	"fmt"
	"time"

	"GO_TCP_scanner/internal/scanner"
)

func main() {
	host := "scanme.nmap.org"
	for _, port := range []int{80, 81} {
		if scanner.ScanPort(host, port, 2*time.Second) {
			fmt.Printf("%s:%d is OPEN\n", host, port)
		} else {
			fmt.Printf("%s:%d is closed/filtered\n", host, port)
		}
	}
}