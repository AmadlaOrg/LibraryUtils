package ristretto

import (
	"errors"

	"github.com/dgraph-io/ristretto/v2"
)

type IRistretto interface{}

type SRistretto struct {
	ristrettoInstance *ristretto.Cache[K, V]
}

var (
	ristrettoNewCache = ristretto.NewCache
)

func (service *SRistretto) Connect() error {
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

func (service *SRistretto) Close() {
	service.ristrettoInstance.Close()
}

func (service *SRistretto) Set(item IRistretto) {
	service.ristrettoInstance.Set("key", "value", 1)
	service.ristrettoInstance.Wait()
}

func (service *SRistretto) Get(key string) (IRistretto, error) {
	val, found := service.ristrettoInstance.Get(key)
	if !found {
		return nil, errors.New("key not found")
	}

	return val
}

func (service *SRistretto) Delete(key string) error {
	service.ristrettoInstance.Del(key)
	return nil
}
