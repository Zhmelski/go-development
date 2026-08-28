package cache

import (
	"sync"
	"time"
)

const (
	// 'ExpirationTimeout' - the time an element remains inactive before being automatically deleted
	ExpirationTimeout = 10 * time.Second
	// 'CleanupInterval' - automatic cache cleaning interval
	CleanupInterval = 1 * time.Second
)

// 'Cache' represents a cache with limited capacity and automatic deletion
type Cache struct {
	items    map[string]*CacheItem
	capacity int
	mu       sync.RWMutex
	stopChan chan struct{}
}

// 'NewCache' creates a new cache with the specified capacity and starts background cleaning
func NewCache(capacity int) *Cache {
	if capacity <= 0 {
		panic("cache capacity must be positive")
	}

	cache := &Cache{
		items:    make(map[string]*CacheItem),
		capacity: capacity,
		stopChan: make(chan struct{}),
	}

	// Starting a background goroutine for automatic cleanup
	go cache.startCleanup()

	return cache
}

// 'Set' adds or updates an object in the cache
func (c *Cache) Set(key string, obj any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If the element already exists, update it
	if item, exists := c.items[key]; exists {
		item.Value = obj
		item.UpdateAccess()
		return
	}

	// If the cache is full, delete the item with the oldest access time
	if len(c.items) >= c.capacity {
		c.removeLRU()
	}

	// Adding a new element
	c.items[key] = &CacheItem{
		Value:      obj,
		LastAccess: time.Now(),
	}
}

// 'Get' returns the object by key or nil if the object is not found
func (c *Cache) Get(key string) any {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.items[key]
	if !exists {
		return nil
	}

	// Updating the last access time
	item.UpdateAccess()
	return item.Value
}

// 'Remove' removes an object by key and returns true if the object was found
func (c *Cache) Remove(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.items[key]
	if exists {
		delete(c.items, key)
	}
	return exists
}

// 'Size' returns the current number of items in the cache
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// 'Stop' stops background cache cleaning
func (c *Cache) Stop() {
	close(c.stopChan)
}

// 'removeLRU' removes the item with the oldest access time (Least Recently Used)
func (c *Cache) removeLRU() {
	var oldestKey string
	var oldestTime time.Time
	first := true

	for key, item := range c.items {
		if first || item.LastAccess.Before(oldestTime) {
			oldestKey = key
			oldestTime = item.LastAccess
			first = false
		}
	}

	if oldestKey != "" {
		delete(c.items, oldestKey)
	}
}

// 'startCleanup' runs periodic cleaning of obsolete items
func (c *Cache) startCleanup() {
	ticker := time.NewTicker(CleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanupExpired()
		case <-c.stopChan:
			return
		}
	}
}

// 'cleanupExpired' removes elements that have not been accessed for more than ExpirationTimeout
func (c *Cache) cleanupExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	keysToDelete := make([]string, 0)

	for key, item := range c.items {
		if item.IsExpired(ExpirationTimeout) {
			keysToDelete = append(keysToDelete, key)
		}
	}

	for _, key := range keysToDelete {
		delete(c.items, key)
	}
}
