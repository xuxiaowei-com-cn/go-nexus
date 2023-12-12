package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestGetBrowseRepositoryRecursion_maven_proxy(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY", "maven-central")
	var tempDir = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY_TEMP_DIR", filepath.Join(os.TempDir(), "GO_NEXUS_MAVEN_PROXY_REPOSITORY", repository))
	t.Log("tempDir", tempDir)

	var path = ""

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	downloadMavenProxyRepository(t, client, repository, tempDir)

	GetBrowseRepositoryRecursion(t, baseURL, username, password, repository, path, client)
}

func GetBrowseRepositoryRecursion(t *testing.T, baseURL string, username string, password string, repository string, path string, client *Client) {

	browses, response, err := client.ExtDirect.GetBrowseRepository(repository, path)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NotEqual(t, 0, len(browses))

	for _, browse := range browses {
		if browse.Type == "file" {
			t.Log(browse.Url)
		} else {
			GetBrowseRepositoryRecursion(t, baseURL, username, password, repository, path+browse.Href, client)
		}
	}

}
