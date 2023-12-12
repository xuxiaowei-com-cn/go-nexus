package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPostExtDirectRecursion_maven_proxy(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY", "maven-central")

	var node = "/"

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	ExtDirectRecursion(t, baseURL, username, password, repository, node, client)
}

func ExtDirectRecursion(t *testing.T, baseURL, username, password, repository string, node string, client *Client) {

	extDirect, response, err := client.ExtDirect.PostExtDirectRead(repository, node)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	assert.Equal(t, true, extDirect.Result.Success)
	//assert.NotEqual(t, 0, len(extDirect.Result.Data))

	//t.Logf("%s %s ExtDirect Result Len: %d", repository, node, len(extDirect.Result.Data))

	for _, data := range extDirect.Result.Data {
		if "folder" == data.Type {

			ExtDirectRecursion(t, baseURL, username, password, repository, data.Id, client)

		} else if "component" == data.Type {

			ExtDirectRecursion(t, baseURL, username, password, repository, data.Id, client)

		} else if "asset" == data.Type {

			postExtDirectReadAsset, response, err := client.ExtDirect.PostExtDirectReadAsset(repository, data.AssetId)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, response.StatusCode)

			assert.Equal(t, true, postExtDirectReadAsset.Result.Success)

			t.Logf("%s/repository/%s/%s %d", baseURL, repository, data.Id, postExtDirectReadAsset.Result.Data.Size)

		} else {
			t.Fatalf("Unknown Type: %s", data.Type)
		}
	}

}
