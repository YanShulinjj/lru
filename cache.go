/* ----------------------------------
*  @author suyame 2022-06-28 9:42:00
*  Crazy for Golang !!!
*  IDE: GoLand
*  Implementation of LRU cache-mechanism
*-----------------------------------*/

package lru

import "sync"

// ListNode is a node of bidirectional-cycle-list.
type ListNode struct {
	// key is Cache item's key.
	key   string
	next  *ListNode
	prior *ListNode
}

// Cache is a cache with capacity, it will discard item when capacity is not enough.
type Cache struct {
	sync.RWMutex
	name     string
	len      int
	capacity int
	items    map[string]interface{}
	head     *ListNode
}

// NewCache add a new cache
func NewCache(name string, capacity int) *Cache {
	// capacity must be larger than 1
	if capacity < 1 {
		panic("capacity must be larger than 1.")
	}
	head := &ListNode{}
	head.next = head
	head.prior = head
	return &Cache{
		name:     name,
		len:      0,
		capacity: capacity,
		items:    make(map[string]interface{}),
		head:     head,
	}
}

// Exist judge key if it existed.
func (cache *Cache) Exist(key string) (*ListNode, bool) {
	cache.RLock()
	defer cache.RUnlock()
	if cache.len == 0 {
		return nil, false
	}
	p := cache.head.next
	for p != cache.head {
		if p.key == key {
			break
		}
		p = p.next
	}
	return p, p != cache.head
}

// Get get the value of key,
// If not exist, return error.
func (cache *Cache) Get(key string) (interface{}, error) {
	p, ok := cache.Exist(key)
	if !ok {
		return nil, KeyNotFoundError
	}
	cache.Lock()
	defer cache.Unlock()
	if cache.len == 1 {
		return cache.items[p.key], nil
	}
	// need put p node at list head.
	// 1. remove p from link
	p.prior.next = p.next
	p.next.prior = p.prior
	// 2. p add the head of link
	p.next = cache.head.next
	cache.head.next.prior = p
	cache.head.next = p
	p.prior = cache.head

	return cache.items[p.key], nil
}

// Put add new item.
// If capacity is not enough, will remove the tail of link.
func (cache *Cache) Put(key string, value interface{}) error {
	// 1. Check exist
	_, ok := cache.Exist(key)
	if ok {
		return KeyHasExistedError
	}
	// 2. Add new item
	cache.Lock()
	defer cache.Unlock()
	cache.items[key] = value
	node := &ListNode{
		key: key,
	}
	// add node to head.
	node.next = cache.head.next
	cache.head.next.prior = node
	cache.head.next = node
	node.prior = cache.head
	// judge capacity
	if cache.len == cache.capacity {
		// full
		// need remove tail
		tail := cache.head.prior
		tail.prior.next = tail.next
		tail.next.prior = tail.prior
		delete(cache.items, tail.key)
	} else {
		cache.len++
	}
	return nil
}
