package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestListAssets(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")
	var repository = os.Getenv("GO_NEXUS_ASSETS_REPOSITORY")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	requestQuery := &ListAssetsRequest{
		Repository: repository,
	}

	pageAssetXO, response, err := client.Assets.ListAssets(requestQuery)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log("ContinuationToken:", pageAssetXO.ContinuationToken)
	t.Log("Items Len:", len(pageAssetXO.Items))
}
