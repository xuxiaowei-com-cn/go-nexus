package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetSwagger(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")

	client, err := NewClient(baseURL, "", "")
	assert.NoError(t, err)

	swagger, response, err := client.Swagger.GetSwagger()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log("Nexus Version:", swagger.Info.Version)
}
