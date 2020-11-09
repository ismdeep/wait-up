package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func showHelp() {
	fmt.Print("Usage: wait-up IP [--help]\n" +
		"    IP      IP Address\n")
}

func WaitUpByPing(ipAddr string) {
	fmt.Printf("Wait up [%s] ", ipAddr)

	pingFunc := PingWithICMP

	if runtime.GOOS == "darwin" {
		pingFunc = PingWithMacOSCommand
	}

	for true {
		fmt.Print(".")
		if pingFunc(ipAddr) {
			break
		} else {
			time.Sleep(1 * time.Second)
			continue
		}
	}
	fmt.Println(" connected")
}

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "--help" {
		showHelp()
		return
	}

	WaitUpByPing(os.Args[1])
}
