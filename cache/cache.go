package cache

import (
	"time"

	"github.com/AmadlaOrg/LibraryUtils/database/sqlite"
)

type Cache interface {
	Open() error
	Close() error
	Insert(entry *map[string]string, ttl time.Duration) error
	Select() (*map[string]string, error)
	Delete(entry string) error
}

type cacheImpl struct {
	database sqlite.Database
}

// Open is for opening a connection with the cache storage (e.g.: SQLite3 database, Ristretto memory-bound Go cache)
func (service *cacheImpl) Open() error {
	return nil
}
