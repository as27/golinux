# golinux

Specially in Windows 10 it is not so simple to set the env variables for a cross compilation.

`golinux` wraps `go build` and calls it with `GOOS=linux` and `GOARCH=amd64`. 
