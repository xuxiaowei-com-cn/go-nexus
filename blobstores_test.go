package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetBlobStoresQuotaStatus(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var name = Getenv("GO_NEXUS_BLOB_STORES_FOR_QUOTA_STATUS", "default")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	blobStoreQuotaResultXO, response, err := client.BlobStores.GetBlobStoresQuotaStatus(name)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log(blobStoreQuotaResultXO)

}

func TestListBlobStores(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	genericBlobStoreApiResponse, response, err := client.BlobStores.ListBlobStores()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NotEqual(t, 0, len(genericBlobStoreApiResponse))

}

func TestGetBlobStoresFile(t *testing.T) {
	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var name = Getenv("GO_NEXUS_BLOB_STORES_FOR_FILE", "default")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	fileBlobStoreApiModel, response, err := client.BlobStores.GetBlobStoresFile(name)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log(fileBlobStoreApiModel)
}
