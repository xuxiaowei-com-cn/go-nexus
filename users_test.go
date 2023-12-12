package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListUsers(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	requestQuery := ListUsersQuery{}

	apiUsers, response, err := client.Users.ListUsers(requestQuery)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	for _, apiUser := range apiUsers {
		t.Log("UserId", apiUser.UserId)
	}

}
