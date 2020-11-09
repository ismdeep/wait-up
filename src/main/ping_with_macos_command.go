package main

import (
	"github.com/go-cmd/cmd"
	"strings"
)

func PingWithMacOSCommand(ipAddr string) bool {
	envCmd := cmd.NewCmd("ping", "-c", "1", ipAddr)
	status := envCmd.Start()
	cursorLine := ""
	finalStatus := <-status
	for _, line := range finalStatus.Stdout {
		if strings.Contains(line, "bytes from") {
			cursorLine = line
			break
		}
	}

	return strings.Contains(cursorLine, "time=") && strings.Contains(cursorLine, "ms")
}
