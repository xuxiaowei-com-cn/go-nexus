package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestListAssets_GetAssets_DeleteAssets_maven_proxy(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY", "maven-central")
	var tempDir = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY_TEMP_DIR", filepath.Join(os.TempDir(), "GO_NEXUS_MAVEN_PROXY_REPOSITORY", repository))
	t.Log("tempDir", tempDir)

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	downloadMavenProxyRepository(t, client, repository, tempDir)

	requestQuery := &ListAssetsQuery{
		Repository: repository,
	}

	pageAssetXO, response, err := client.Assets.ListAssets(requestQuery)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.NotEqual(t, 0, len(pageAssetXO.Items))

	t.Log("ContinuationToken:", pageAssetXO.ContinuationToken)
	t.Log("Items Len:", len(pageAssetXO.Items))

	item := pageAssetXO.Items[0]

	assetXO, response, err := client.Assets.GetAssets(item.Id)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log("DownloadUrl:", assetXO.DownloadUrl)

	response, err = client.Assets.DeleteAssets(item.Id)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
