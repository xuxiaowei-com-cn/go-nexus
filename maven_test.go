package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetMavenRepositoryGroup(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")
	var repository = os.Getenv("GO_NEXUS_MAVEN_REPOSITORY")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	groupRepository, response, err := client.Maven.GetMavenRepository(MavenTypeGroup, repository)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(groupRepository.Url)
}

func TestGetMavenRepositoryProxy(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	groupRepository, response, err := client.Maven.GetMavenRepository(MavenTypeProxy, "maven-central")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(groupRepository.Url)
}

func TestGetMavenRepositoryHosted(t *testing.T) {

	var baseURL = os.Getenv("GO_NEXUS_BASE_URL")
	var username = os.Getenv("GO_NEXUS_USERNAME")
	var password = os.Getenv("GO_NEXUS_PASSWORD")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	groupRepository, response, err := client.Maven.GetMavenRepository(MavenTypeHosted, "maven-snapshots")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(groupRepository.Url)
}
