package ristretto

import (
	"errors"

	"github.com/dgraph-io/ristretto/v2"
)

type Ristretto interface {
	Connect() error
	Close()
	Set(key, value string, cost int64)
	Get(key string) (string, error)
	Delete(key string) error
}

type ristrettoImpl struct {
	ristrettoInstance *ristretto.Cache[string, string]
}

var (
	ristrettoNewCache = ristretto.NewCache[string, string]
)

func (service *ristrettoImpl) Connect() error {
	var err error
	service.ristrettoInstance, err = ristrettoNewCache(&ristretto.Config[string, string]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return err
	}
	return nil
}

func (service *ristrettoImpl) Close() {
	service.ristrettoInstance.Close()
}

func (service *ristrettoImpl) Set(key, value string, cost int64) {
	service.ristrettoInstance.Set(key, value, cost)
	service.ristrettoInstance.Wait()
}

func (service *ristrettoImpl) Get(key string) (string, error) {
	val, found := service.ristrettoInstance.Get(key)
	if !found {
		return "", errors.New("key not found")
	}

	return val, nil
}

func (service *ristrettoImpl) Delete(key string) error {
	service.ristrettoInstance.Del(key)
	return nil
}
