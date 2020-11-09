package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
)

func showHelp() {
	fmt.Print("Usage: wait-up IP [--help]\n" +
		"    IP      IP Address\n")
}

func PingOneTime(ipAddr string) bool {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ipAddr)
	if err != nil {
		return false
	}
	passed := false
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		passed = true
	}
	p.OnIdle = func() {
	}
	_ = p.Run()
	return passed
}

func WaitUpByPing(ipAddr string) {
	fmt.Printf("Wait up [%s] ", ipAddr)
	for true {
		fmt.Print(".")
		if PingOneTime(ipAddr) {
			break
		} else {
			continue
		}
	}
	fmt.Printf(" connected")
}

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "--help" {
		showHelp()
		return
	}

	WaitUpByPing(os.Args[1])

}
