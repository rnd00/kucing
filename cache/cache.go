package cache

import (
	"sync"
)

// build a key-value cache here

type Cac struct {
	data map[string]bool
	mu   sync.RWMutex
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
