package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

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
