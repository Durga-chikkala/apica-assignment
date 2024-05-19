package service

import (
	"github.com/Durga-chikkala/apica-assignment/models"
	"github.com/Durga-chikkala/apica-assignment/store"
)

type Service struct {
	lruCacherStore store.LRUCache
}

func New(lruCacherStore store.LRUCache) Service {
	return Service{lruCacherStore: lruCacherStore}
}

// Get retrieves the value associated with the given key from the store's LRU cache.
func (s Service) Get(key string) string {
	return s.lruCacherStore.Get(key)
}

// Set stores the key-value pair in the store's LRU cache.
func (s Service) Set(cache *models.CacheData) map[string]*models.CacheData {
	return s.lruCacherStore.Set(cache)
}

// Delete removes the value associated with the given key from the store's LRU cache.
func (s Service) Delete(key string) map[string]*models.CacheData {
	return s.lruCacherStore.Delete(key)
}

// GetAllKeys retrieves all the keys from the store's LRU cache.
func (s Service) GetAllKeys() map[string]*models.CacheData {
	return s.lruCacherStore.GetAllKeys()
}
