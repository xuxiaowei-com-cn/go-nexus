package nexus

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// 下载 Maven 代理仓库 测试数据
func downloadMavenProxyRepository(t *testing.T, client *Client, repository string) {

	metadataSuffixList := []string{"", ".md5", ".sha1", ".sha256", ".sha512"}
	for _, suffix := range metadataSuffixList {
		u := fmt.Sprintf("repository/%s/org/springframework/boot/spring-boot/maven-metadata.xml%s", repository, suffix)
		req, err := client.NewRequest(http.MethodGet, u, nil, nil)
		if err != nil {
			assert.NoError(t, err)
		}
		_, err = client.Do(req, "")
		if err != nil {
			assert.NoError(t, err)
		}
	}
	for i := 1; i <= 18; i++ {
		u := fmt.Sprintf("repository/%s/org/springframework/boot/spring-boot/2.7.%d/spring-boot-2.7.%d.pom", repository, i, i)
		req, err := client.NewRequest(http.MethodGet, u, nil, nil)
		if err != nil {
			assert.NoError(t, err)
		}
		_, err = client.Do(req, "")
		if err != nil {
			assert.NoError(t, err)
		}
	}
	fileSuffixList := []string{"-javadoc.jar", "-javadoc.jar.asc", "-javadoc.jar.md5", "-javadoc.jar.sha1",
		"-sources.jar", "-sources.jar.asc", "-sources.jar.md5", "-sources.jar.sha1",
		".jar", ".jar.asc", ".jar.md5", ".jar.sha1",
		".module", ".module.asc", ".module.md5", ".module.sha1",
		".pom", ".pom.asc", ".pom.md5", ".pom.sha1"}
	for i := 1; i <= 3; i++ {
		for _, file := range fileSuffixList {
			u := fmt.Sprintf("repository/%s/org/springframework/boot/spring-boot/2.7.%d/spring-boot-2.7.%d%s", repository, i, i, file)
			req, err := client.NewRequest(http.MethodGet, u, nil, nil)
			if err != nil {
				assert.NoError(t, err)
			}
			_, err = client.Do(req, "")
			if err != nil {
				assert.NoError(t, err)
			}
		}
	}
}

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
