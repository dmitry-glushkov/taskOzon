package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStore_Set(t *testing.T) {
	store := New()

	store.Set("no_ttl", "value", 0)
	time.Sleep(1 * time.Second)
	assert.NotNil(t, store.Data["no_ttl"])

	store.Set("with_ttl", "value_to_ex", 1)
	time.Sleep(2 * time.Second)
	assert.Nil(t, store.Data["with_ttl"])
}

func TestStore_Get(t *testing.T) {
	store := New()

	store.Set("key", "value", 0)
	assert.NotNil(t, store.Get("key"))
}

func TestStore_GetAllKeys(t *testing.T) {
	store := New()

	assert.Equal(t, []string{}, store.GetAllKeys())

	store.Set("key1", "val", 0)
	store.Set("key2", "val", 0)
	assert.Equal(t, []string{"key1", "key2"}, store.GetAllKeys())
}

func TestStore_Delete(t *testing.T) {
	store := New()

	store.Set("key", "val", 0)
	assert.NotNil(t, store.Data["key"])
	store.Delete("key")
	assert.Nil(t, store.Data["key"])
}
