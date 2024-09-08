package config

import (
	"log"
	"time"

	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
)

var LimiterStore limiter.Store

func NewLimiterStore() {
	var err error
	LimiterStore, err = memorystore.New(&memorystore.Config{
		Tokens:   1000,
		Interval: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
}
