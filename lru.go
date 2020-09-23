/*
  https://leetcode.com/problems/lru-cache/
*/

package main

import "fmt"

// CacheSize defines the size of the LRU cache
const CacheSize = 4

// Entry defines each unit of the LRU cache
type Entry struct {
  left *Entry
  right *Entry
  key int
  value int 
}

// LruCache struct consists of the map and a doubly linked list
// to track recent usage of the cache entries
type LruCache struct {
   size  int
   start *Entry
   end   *Entry
   cache map[int]Entry 
}

func (c *LruCache) moveToTop(entry *Entry) {
  if c.start == nil {
    c.start = entry
    c.end = entry
    return
  }

  origLeft := entry.left
  origRight := entry.right

  if origLeft != nil {
    origLeft.right = origRight
  }

  if origRight != nil {
    origRight.left = origLeft
  }

  entry.left = nil
  entry.right = c.start
  c.start.left = entry

  c.start = entry
}

func (c *LruCache) removeEnd() {
   if c.end == nil {
     fmt.Printf("No elements in the doubly linked list") 
     return
   }
   delete(c.cache, c.end.key);
   c.end = c.end.left
   c.end.right = nil
}

func (c *LruCache) get(key int) {
   if entry, ok := c.cache[key]; ok {
     c.moveToTop(&entry)
     fmt.Printf("Key: %v, Value: %v \n", entry.key, entry.value)
   } else {
     fmt.Println("Error: Entry does not exist")
   }
}

func (c *LruCache) put(key int, val int) {
  if (len(c.cache) == c.size) {
    c.removeEnd()
  }
  if entry, ok := c.cache[key]; !ok {
    entry := Entry {
      left: nil,
      right: nil,
      key: key,
      value: val,
    }
    c.cache[key] = entry
    c.moveToTop(&entry)
  } else {
    entry.value = val
    c.moveToTop(&entry)
  }
} 

func main() {
  lruCache := LruCache {
    size: CacheSize,
    start: nil,
    end: nil,
    cache: make(map[int]Entry),
  }
  lruCache.put(10, 20);
  lruCache.put(20, 20);
  lruCache.put(30, 20);
  lruCache.put(40, 20);
  lruCache.put(50, 20);
  lruCache.get(10);
  lruCache.get(30);
}
