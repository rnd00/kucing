package cache

import (
	"log"
	"sync"
	"time"
)

// build a key-value cache here

type Cac struct {
	data    map[string]bool
	mu      sync.RWMutex
	timeout time.Duration
}

func NewCache() *Cac {
	return &Cac{
		data: make(map[string]bool),
	}
}

func (c *Cac) SetKey(k string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = true
}

func (c *Cac) CheckKey(k string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.data[k]
}

func (c *Cac) DeleteKey(k string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, k)
}

// ---

func (c *Cac) AddTimeout(dur time.Duration) *Cac {
	Cache := NewCache()
	Cache.timeout = dur
	return Cache
}

func (c *Cac) SetKeyWithTimeout(k string) {
	c.SetKey(k)
	if c.timeout > 0 {
		go c.timeoutFunc(k)
	}
}

func (c *Cac) timeoutFunc(k string) {
	select {
	case <-time.After(c.timeout):
		log.Println("Cache: Timeout for key", k, "-- Deleting")
		c.DeleteKey(k)
	}
}
