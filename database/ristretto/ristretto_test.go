package ristretto

import (
	"errors"
	"testing"
	"time"

	ristrettoLib "github.com/dgraph-io/ristretto/v2"
	"github.com/stretchr/testify/assert"
)

func newConnectedService(t *testing.T) *ristrettoImpl {
	t.Helper()
	svc := &ristrettoImpl{}
	err := svc.Connect()
	assert.NoError(t, err)
	return svc
}

func TestConnect(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &ristrettoImpl{}
		err := svc.Connect()
		assert.NoError(t, err)
		assert.NotNil(t, svc.ristrettoInstance)
		svc.Close()
	})

	t.Run("error from ristrettoNewCache", func(t *testing.T) {
		origRistrettoNewCache := ristrettoNewCache
		defer func() { ristrettoNewCache = origRistrettoNewCache }()

		ristrettoNewCache = func(config *ristrettoLib.Config[string, string]) (*ristrettoLib.Cache[string, string], error) {
			return nil, errors.New("cache creation failed")
		}

		svc := &ristrettoImpl{}
		err := svc.Connect()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cache creation failed")
	})
}

func TestSet(t *testing.T) {
	svc := newConnectedService(t)
	defer svc.Close()

	svc.Set("mykey", "myvalue", 1)

	// Allow ristretto time to process
	time.Sleep(10 * time.Millisecond)

	val, err := svc.Get("mykey")
	assert.NoError(t, err)
	assert.Equal(t, "myvalue", val)
}

func TestGet(t *testing.T) {
	svc := newConnectedService(t)
	defer svc.Close()

	t.Run("existing key", func(t *testing.T) {
		svc.ristrettoInstance.Set("mykey", "myvalue", 1)
		svc.ristrettoInstance.Wait()
		time.Sleep(10 * time.Millisecond)

		val, err := svc.Get("mykey")
		assert.NoError(t, err)
		assert.Equal(t, "myvalue", val)
	})

	t.Run("nonexistent key", func(t *testing.T) {
		val, err := svc.Get("nonexistent")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "key not found")
		assert.Empty(t, val)
	})

	t.Run("get after delete", func(t *testing.T) {
		svc.ristrettoInstance.Set("delkey", "delvalue", 1)
		svc.ristrettoInstance.Wait()
		time.Sleep(10 * time.Millisecond)

		svc.ristrettoInstance.Del("delkey")
		time.Sleep(10 * time.Millisecond)

		val, err := svc.Get("delkey")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "key not found")
		assert.Empty(t, val)
	})
}

func TestDelete(t *testing.T) {
	svc := newConnectedService(t)
	defer svc.Close()

	svc.ristrettoInstance.Set("rmkey", "rmvalue", 1)
	svc.ristrettoInstance.Wait()
	time.Sleep(10 * time.Millisecond)

	err := svc.Delete("rmkey")
	assert.NoError(t, err)

	time.Sleep(10 * time.Millisecond)
	_, getErr := svc.Get("rmkey")
	assert.Error(t, getErr)
}

func TestClose(t *testing.T) {
	svc := newConnectedService(t)
	// Should not panic
	assert.NotPanics(t, func() {
		svc.Close()
	})
}
