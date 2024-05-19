package store

import (
	"time"

	"github.com/Durga-chikkala/apica-assignment/cacher"
	"github.com/Durga-chikkala/apica-assignment/models"
)

type Store struct {
	cache *cacher.Cache
}

func New(cache *cacher.Cache) Store {
	return Store{cache: cache}
}

// Get retrieves the value associated with the given key from the store's cache.
// It moves the accessed element to the front of the LRU list to indicate recent usage.
func (s *Store) Get(key string) string {
	s.cache.Mutex.Lock()
	defer s.cache.Mutex.Unlock()

	if elem, ok := s.cache.Cache[key]; ok {
		s.cache.MoveToFront(elem)
		return elem.Value
	}

	return ""
}

// Set stores the key-value pair in the cache, updating the value if the key already exists.
// It also maintains the LRU list by moving accessed elements to the front and removing the least recently used element when the cache exceeds its capacity.
func (s *Store) Set(cache *models.CacheData) map[string]*models.CacheData {
	s.cache.Mutex.Lock()
	defer s.cache.Mutex.Unlock()

	if elem, ok := s.cache.Cache[cache.Key]; ok {
		elem.Value = cache.Value
		s.cache.MoveToFront(elem)
	}

	if len(s.cache.Cache) >= s.cache.Capacity {
		delete(s.cache.Cache, s.cache.Tail.Key)
		s.cache.RemoveTail()

	}

	newEntry := &models.CacheData{
		Key:        cache.Key,
		Value:      cache.Value,
		Expiration: cache.Expiration,
		TimeStamp:  time.Now(),
		Next:       s.cache.Head,
	}

	if s.cache.Head != nil {
		s.cache.Head.Prev = newEntry
	}

	s.cache.Head = newEntry
	if s.cache.Tail == nil {
		s.cache.Tail = newEntry
	}

	s.cache.Cache[cache.Key] = newEntry

	return s.cache.Cache
}

// Delete removes the cache entry associated with the specified key.
func (s *Store) Delete(key string) map[string]*models.CacheData {
	s.cache.Mutex.Lock()
	defer s.cache.Mutex.Unlock()

	if elem, ok := s.cache.Cache[key]; ok {
		delete(s.cache.Cache, key)
		if elem.Prev != nil {
			elem.Prev.Next = elem.Next
		} else {
			s.cache.Head = elem.Next
		}

		if elem.Next != nil {
			elem.Next.Prev = elem.Prev
		} else {
			s.cache.Tail = elem.Prev
		}
	}

	return s.cache.Cache
}

// GetAllKeys returns a map containing all keys in the cache along with their associated cache data.
func (s *Store) GetAllKeys() map[string]*models.CacheData {
	return s.cache.Cache
}
