package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestListRepository(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	repositories, response, err := client.Repository.ListRepository()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NotEqual(t, 0, len(repositories))
	t.Log("Repositories Len:", len(repositories))
}
