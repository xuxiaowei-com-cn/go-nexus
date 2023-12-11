package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestPostExtDirect(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")
	var repository = os.Getenv("GO_NEXUS_EXT_DIRECT_REPOSITORY")

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
