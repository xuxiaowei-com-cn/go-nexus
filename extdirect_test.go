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
	requestBody := &PostExtDirectRequest{
		Action: "coreui_Browse",
		Method: "read",
		Type:   "rpc",
		Data: []PostExtDirectData{
			{
				RepositoryName: repository,
				Node:           node,
			},
		},
	}

	extDirect, response, err := client.ExtDirect.PostExtDirect(requestBody)
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

			t.Logf("%s/repository/%s/%s", baseURL, repository, data.Id)

		} else {
			t.Fatalf("Unknown Type: %s", data.Type)
		}
	}

}
