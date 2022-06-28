/* ----------------------------------
*  @author suyame 2022-06-28 10:50:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package lru

import "testing"

func TestCache_Get(t *testing.T) {
	cache := NewCache("cache", 5)
	_, err := cache.Get("key1")
	if err != KeyNotFoundError {
		t.Error(err)
	}
	// put
	err = cache.Put("key1", "value1")
	if err != nil {
		t.Error(err)
	}
	// Again get
	v, err := cache.Get("key1")
	if err != nil {
		t.Error(err)
	}
	if v != "value1" {
		t.Error(v)
	}

	// Put key2~6
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")
	cache.Get("key1")
	cache.Put("key4", "value4")
	cache.Put("key5", "value5")
	cache.Put("key6", "value6")
	// get key1
	v, err = cache.Get("key1")
	if err != nil {
		t.Error(err)
	}
	if v != "value1" {
		t.Error(v)
	}
}
