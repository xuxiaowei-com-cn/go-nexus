package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetBrowseRepository(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")
	var repository = os.Getenv("GO_NEXUS_EXT_DIRECT_REPOSITORY")

	var path = ""

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	GetBrowseRepositoryRecursion(t, baseURL, username, password, repository, path, client)
}

func GetBrowseRepositoryRecursion(t *testing.T, baseURL string, username string, password string, repository string, path string, client *Client) {

	browses, response, err := client.ExtDirect.GetBrowseRepository(repository, path)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	for _, browse := range browses {
		if browse.Type == "file" {
			t.Log(browse.Url)
		} else {
			GetBrowseRepositoryRecursion(t, baseURL, username, password, repository, path+browse.Href, client)
		}
	}

}
