// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/viper"
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
func SetDirection(args string) string {
	argsSplit := strings.Split(args, " ")
	direction := "main"
	if strings.ContainsAny(args, " ") {
		if strings.HasPrefix(argsSplit[1], "-") {
			direction = "main"
		} else {
			direction = argsSplit[1]
		}
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

// LoadConfig loads the configuration from a config file.
func LoadConfig() {
	Log("Loading viper config.", "info")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/octohost/")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Log(fmt.Sprintf("No config file found: %s \n", err), "info")
	}
	viper.SetEnvPrefix("OCTO")
}

// SpaceCheck looks to see if the value contains a space.
func SpaceCheck(value string, name string) {
	if strings.Contains(value, " ") {
		fmt.Printf("A %s cannot contain a space.\n", name)
		os.Exit(1)
	}
}
