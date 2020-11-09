package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"time"
)

func PingWithICMP(ipAddr string) bool {
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
	err = p.Run()
	if err != nil {
		fmt.Print(err.Error())
	}
	return passed
}
