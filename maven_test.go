package nexus

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
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

func TestGetMavenRepository_maven_group(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_GROUP_REPOSITORY", "maven-public")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	groupRepository, response, err := client.Maven.GetMavenRepository(MavenTypeGroup, repository)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(groupRepository.Url)
}

func TestGetMavenRepository_maven_proxy(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_PROXY_REPOSITORY", "maven-central")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	proxyRepository, response, err := client.Maven.GetMavenRepository(MavenTypeProxy, repository)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(proxyRepository.Url)
}

func TestGetMavenRepository_maven_hosted_snapshots(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_HOSTED_REPOSITORY", "maven-snapshots")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	hostedRepository, response, err := client.Maven.GetMavenRepository(MavenTypeHosted, repository)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(hostedRepository.Url)
}

func TestGetMavenRepository_maven_hosted_releases(t *testing.T) {

	var baseURL = Getenv("GO_NEXUS_BASE_URL", "http://127.0.0.1:8081/")
	var username = Getenv("GO_NEXUS_USERNAME", "admin")
	var password = Getenv("GO_NEXUS_PASSWORD", "password")
	var repository = Getenv("GO_NEXUS_MAVEN_HOSTED_REPOSITORY", "maven-releases")

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	hostedRepository, response, err := client.Maven.GetMavenRepository(MavenTypeHosted, repository)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	t.Log(hostedRepository.Url)
}
