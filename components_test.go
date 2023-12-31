package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestListComponents_GetComponents_DeleteComponents_maven_proxy(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY", "maven-central")
	var tempDir = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY_TEMP_DIR", filepath.Join(os.TempDir(), "GO_NEXUS_MAVEN_PROXY_REPOSITORY", repository))
	t.Log("tempDir", tempDir)

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	downloadMavenProxyRepository(t, client, repository, tempDir)

	requestQuery := &ListComponentsQuery{
		Repository: repository,
	}

	pageComponentXO, response, err := client.Components.ListComponents(requestQuery)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NotEqual(t, 0, len(pageComponentXO.Items))

	item := pageComponentXO.Items[0]

	componentXO, response, err := client.Components.GetComponents(item.Id)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log("Id:", componentXO.Id)

	response, err = client.Components.DeleteComponents(item.Id)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

}

func TestListComponentsRecursion_maven_proxy(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY", "maven-central")
	var tempDir = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY_TEMP_DIR", filepath.Join(os.TempDir(), "GO_NEXUS_MAVEN_PROXY_REPOSITORY", repository))
	t.Log("tempDir", tempDir)

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	downloadMavenProxyRepository(t, client, repository, tempDir)

	ListComponentsRecursion(t, repository, "", client)
}

func ListComponentsRecursion(t *testing.T, repository string, continuationToken string, client *Client) {
	requestQuery := &ListComponentsQuery{
		Repository: repository,
	}

	if continuationToken != "" {
		requestQuery.ContinuationToken = continuationToken
	}

	pageComponentXO, response, err := client.Components.ListComponents(requestQuery)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NotEqual(t, 0, len(pageComponentXO.Items))

	for _, item := range pageComponentXO.Items {
		t.Logf("%s %s:%s:%s", item.Repository, item.Group, item.Name, item.Version)

		_, response, err = client.Components.GetComponents(item.Id)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		for _, asset := range item.Assets {
			t.Logf("%s %d", asset.DownloadUrl, asset.FileSize)
		}

	}

	if pageComponentXO.ContinuationToken != "" {
		ListComponentsRecursion(t, repository, pageComponentXO.ContinuationToken, client)
	}

}
