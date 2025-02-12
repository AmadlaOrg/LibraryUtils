package cache

import "time"

type ICache interface {
	Open() error
	Close() error
	Insert(entry *map[string]string, ttl time.Duration) error
	Select() (*map[string]string, error)
	Delete(entry string) error
}

type SCache struct {
	database database.IDatabase
}

// Open is for opening a connection with the cache storage (e.g.: SQLite3 database, Ristretto memory-bound Go cache)
func (service *SCache) Open() error {

}
