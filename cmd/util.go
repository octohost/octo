// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// ReturnCurrentUTC returns the current UTC time in RFC3339 format.
func ReturnCurrentUTC() string {
	t := time.Now().UTC()
	dateUpdated := (t.Format(time.RFC3339))
	return dateUpdated
}

// SetDirection returns the direction.
func SetDirection() string {
	direction := ""
	if strings.HasPrefix(os.Args[1], "-") {
		direction = "main"
	} else {
		direction = os.Args[1]
	}
	return direction
}

// Log adds the global Direction to a message and sends to syslog.
// Syslog is setup in main.go
func Log(message, priority string) {
	message = fmt.Sprintf("%s: %s", Direction, message)
	if Verbose {
		time := ReturnCurrentUTC()
		fmt.Printf("%s: %s\n", time, message)
	}
	switch {
	case priority == "debug":
		if os.Getenv("OCTO_DEBUG") != "" {
			log.Print(message)
		}
	default:
		log.Print(message)
	}
}

// GetHostname returns the hostname.
func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
