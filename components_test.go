package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestListComponents(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")
	var repository = os.Getenv("GO_NEXUS_MAVEN_REPOSITORY")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	ListComponentsRecursion(t, repository, "", client)
}

func ListComponentsRecursion(t *testing.T, repository string, continuationToken string, client *Client) {
	request := &ListComponentsQuery{
		Repository: repository,
	}

	if continuationToken != "" {
		request.ContinuationToken = continuationToken
	}

	pageComponentXO, response, err := client.Components.ListComponents(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// t.Log("ContinuationToken:", pageComponentXO.ContinuationToken)
	// t.Log("Items Len:", len(pageComponentXO.Items))

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
