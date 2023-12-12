package nexus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetStatusCheck(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	getStatusCheck, response, err := client.Status.GetStatusCheck()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log("AvailableCPUs", getStatusCheck.AvailableCPUs)
	t.Log("BlobStoresQuota", getStatusCheck.BlobStoresQuota)
	t.Log("BlobStoresReady", getStatusCheck.BlobStoresReady)
	t.Log("CoordinateContentSelectors", getStatusCheck.CoordinateContentSelectors)
	t.Log("DefaultAdminCredentials", getStatusCheck.DefaultAdminCredentials)
	t.Log("DefaultRoleRealm", getStatusCheck.DefaultRoleRealm)
	t.Log("FileBlobStoresPath", getStatusCheck.FileBlobStoresPath)
	t.Log("FileDescriptors", getStatusCheck.FileDescriptors)
	t.Log("LifecyclePhase", getStatusCheck.LifecyclePhase)
	t.Log("NuGetV2Repositories", getStatusCheck.NuGetV2Repositories)
	t.Log("ReadOnlyDetector", getStatusCheck.ReadOnlyDetector)
	t.Log("Scheduler", getStatusCheck.Scheduler)
	t.Log("Scripting", getStatusCheck.Scripting)
	t.Log("ThreadDeadlockDetector", getStatusCheck.ThreadDeadlockDetector)
	t.Log("Transactions", getStatusCheck.Transactions)
}

func TestGetStatus(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")

	client, err := NewClient(baseURL, "", "")
	assert.NoError(t, err)

	response, err := client.Status.GetStatus()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestGetStatusWritable(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")

	client, err := NewClient(baseURL, "", "")
	assert.NoError(t, err)

	response, err := client.Status.GetStatusWritable()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
