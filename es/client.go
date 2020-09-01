package es

import (
	"sync"

	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
)

var (
	// used for concurrent protection
	mutex = &sync.Mutex{}

	// our underlying es client
	client *elasticsearch6.Client
)

// GetClient is a wrapper around our elasticsearch global variable
func GetClient() (*elasticsearch6.Client, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if client == nil {
		es6, err := connect()
		if err != nil {
			return nil, err
		}
		client = es6
	}

	return client, nil
}
