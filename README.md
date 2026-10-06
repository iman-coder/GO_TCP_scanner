# Go TCP Scanner

A small TCP port scanner written in Go. It checks whether selected ports are
open on a host using a connection timeout.

## Run

From the project directory:

```powershell
go run ./cmd/portscan
```

The current example scans ports `80` and `81` on `scanme.nmap.org`.

## Test

```powershell
go test ./...
```
