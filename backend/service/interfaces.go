package service

import "github.com/Durga-chikkala/apica-assignment/models"

type LRUCache interface {
	Get(key string) string
	Set(cache *models.CacheData) map[string]*models.CacheData
	Delete(key string) map[string]*models.CacheData
	GetAllKeys() map[string]*models.CacheData
}
