package scanner

import (
	"net"
	"strconv"
	"time"
)

// ScanPort reports whether a TCP connection to host:port succeeds within
// timeout. A successful connection means something is listening = open.
func ScanPort(host string, port int, timeout time.Duration) bool {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}