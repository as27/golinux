package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	flagGOOS    = flag.String("GOOS", "linux", "set another GOOS value")
	flagGOARCH  = flag.String("GOARCH", "amd64", "set another GOARCH value")
	flagVersion = flag.Bool("version", false, "print out the version")
)

const version = "1.0.0"

func main() {
	flag.Parse()
	if *flagVersion {
		fmt.Println("Version:", version)
		return
	}
	args := []string{"build"}
	if len(args) > 1 {
		args = append(args, os.Args[1:]...)
	}
	cmd := exec.Command("go", args...)
	err := os.Setenv("GOOS", *flagGOOS)
	if err != nil {
		log.Fatal("error setting GOOS env", err)
	}
	err = os.Setenv("GOARCH", *flagGOARCH)
	if err != nil {
		log.Fatal("error setting GOARCH env", err)
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal("error runing go build", err)
	}
}
