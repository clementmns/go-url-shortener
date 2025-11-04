package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	url := "https://www.google.com"
	userUUID := "test-user-uuid"
	shortUrl := "test-short-url"

	SaveUrlMapping(shortUrl, url, userUUID)

	retrievedUrl := RetrieveInitialUrl(shortUrl)
	assert.Equal(t, url, retrievedUrl)
}
