package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestInmemoryCache struct {
	cache Storage
	suite.Suite
}

func (test *TestInmemoryCache) SetupTest() {
	test.cache = NewInMemoryCache()
	_ = test.cache.Flush()
}

func (test *TestInmemoryCache) TestCanSetKey() {
	_ = test.cache.Set("key", "value")
	actual, err := test.cache.Get("key")

	assert.NoError(test.T(), err)
	assert.Equal(test.T(), actual, "value")
}

func (test *TestInmemoryCache) TestKeyNotFoundReturnsError() {
	_ = test.cache.Set("key", "value")
	_, err := test.cache.Get("random-key")

	assert.Error(test.T(), err)
}

func (test *TestInmemoryCache) TestCanFindKey() {
	_ = test.cache.Set("random-key", "random-value")
	actual, err := test.cache.Get("random-key")

	assert.NoError(test.T(), err)
	assert.Equal(test.T(), actual, "random-value")
}

func (test *TestInmemoryCache) TestCanRemoveKey() {
	_ = test.cache.Set("key-1", "random-value")

	err := test.cache.Remove("key-1")

	assert.NoError(test.T(), err)
}

func (test *TestInmemoryCache) TestCanFlushKeys() {
	_ = test.cache.Set("key-1", "random-value")

	flushError := test.cache.Flush()
	_, getKeyError := test.cache.Get("key-1")

	assert.NoError(test.T(), flushError)
	assert.Error(test.T(), getKeyError)
}

func TestNewInMemoryCache(t *testing.T) {
	suite.Run(t, new(TestInmemoryCache))
}
