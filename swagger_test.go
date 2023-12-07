package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetSwagger(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")

	client, err := NewClient(baseURL, "", "")
	assert.NoError(t, err)

	swagger, response, err := client.Swagger.GetSwagger()
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log("Nexus Version:", swagger.Info.Version)
}
