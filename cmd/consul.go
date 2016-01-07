// +build linux darwin freebsd

package cmd

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	consulServer = "127.0.0.1:8500"
	consulToken  = "anonymous"
	consulTries  = 5
)

// ConsulSetup sets up a connection to Consul.
func ConsulSetup() (*consul.Client, error) {
	server := ""
	token := ""
	if server = viper.GetString("consul"); server == "" {
		server = consulServer
	}
	if token = viper.GetString("token"); token == "" {
		token = consulToken
	}
	consul, err := consulConnect(server, token)
	if err != nil {
		Log("Consul connection is bad.", "info")
	}
	return consul, err
}

// consulConnect to the Consul server and hand back a client object.
func consulConnect(server, token string) (*consul.Client, error) {
	var cleanedToken = ""
	config := consul.DefaultConfig()
	config.Address = server
	// Let's clean up the token so it doesn't appear in the logs.
	if token != "" {
		config.Token = token
	}
	consul, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	Log(fmt.Sprintf("server='%s' token='%s'", server, cleanedToken), "debug")
	return consul, nil
}

// Standard Consul tokens have lots of dashes in them - let's split on the dash
// so that we can see part of the token in the logs - helps with debugging.
// We don't want the full token in any logs - that's bad.
func cleanupToken(token string) string {
	first := strings.SplitN(token, "-", 2)[0]
	return first
}

// ConsulGet the value from a key in the Consul KV store.
func ConsulGet(c *consul.Client, key string) string {
	var str string
	Retry(func() error {
		var err error
		str, err = consulGet(c, key)
		return err
	}, consulTries)
	return str
}

// Retry loops through the callback func and tries several times to do the thing.
func Retry(callback func() error, tries int) {
	var err error
	for i := 1; i <= tries; i++ {
		err = callback()
		if err == nil {
			return
		}
		waitTime := time.Duration(tries) * time.Second
		Log(fmt.Sprintf("Consul Failure (%d) - trying again. Max: %d", i, tries), "info")
		if i < tries {
			time.Sleep(waitTime)
		}
	}
	Log("Panic: Giving up on Consul.", "info")
}

// consulGet the value from a key in the Consul KV store.
func consulGet(c *consul.Client, key string) (string, error) {
	var value string
	kv := c.KV()
	key = strings.TrimPrefix(key, "/")
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return "", err
	}
	if pair != nil {
		value = string(pair.Value[:])
	} else {
		value = ""
	}
	Log(fmt.Sprintf("action='consulGet' key='%s'", key), "debug")
	return value, err
}

// ConsulSet the value for a key in the Consul KV store.
func ConsulSet(c *consul.Client, key string, value string) bool {
	var success bool
	Retry(func() error {
		var err error
		success, err = consulSet(c, key, value)
		return err
	}, consulTries)
	return true
}

// consulSet a value for a key in the Consul KV store.
func consulSet(c *consul.Client, key string, value string) (bool, error) {
	key = strings.TrimPrefix(key, "/")
	p := &consul.KVPair{Key: key, Value: []byte(value)}
	kv := c.KV()
	_, err := kv.Put(p, nil)
	if err != nil {
		return false, err
	}
	Log(fmt.Sprintf("action='consulSet' key='%s'", key), "debug")
	return true, err
}

// ConsulDel removes a key from the Consul KV store.
func ConsulDel(c *consul.Client, key string) bool {
	var success bool
	Retry(func() error {
		var err error
		success, err = consulDel(c, key)
		return err
	}, consulTries)
	return true
}

// consulDel removes a key from the Consul KV store.
func consulDel(c *consul.Client, key string) (bool, error) {
	kv := c.KV()
	key = strings.TrimPrefix(key, "/")
	_, err := kv.Delete(key, nil)
	if err != nil {
		return false, err
	}
	Log(fmt.Sprintf("action='consulDel' key='%s'", key), "info")
	return true, err
}
