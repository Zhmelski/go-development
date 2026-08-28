package cache

import "time"

// 'CacheItem' represents a cache item with data and last access time
type CacheItem struct {
	Value      any
	LastAccess time.Time
}

// 'UpdateAccess' updates the last access time of an element
func (item *CacheItem) UpdateAccess() {
	item.LastAccess = time.Now()
}

// 'IsExpired' checks if an item has expired (not been used for more than a specified time)
func (item *CacheItem) IsExpired(timeout time.Duration) bool {
	return time.Since(item.LastAccess) > timeout
}
