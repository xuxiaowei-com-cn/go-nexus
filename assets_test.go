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

	// 默认 Nexus 仓库，访问下列地址，即可产生新文件
	// http://127.0.0.1:8081/repository/maven-central/org/springframework/boot/spring-boot/maven-metadata.xml
	// http://127.0.0.1:8081/repository/maven-central/org/springframework/boot/spring-boot/2.7.18/spring-boot-2.7.18.pom
	// http://127.0.0.1:8081/repository/maven-public/org/springframework/boot/spring-boot/maven-metadata.xml
	// http://127.0.0.1:8081/repository/maven-public/org/springframework/boot/spring-boot/2.7.18/spring-boot-2.7.18.pom

	client, err := NewClient(baseURL, username, password)
	assert.NoError(t, err)

	requestQuery := &ListAssetsQuery{
		Repository: repository,
	}

	pageAssetXO, response, err := client.Assets.ListAssets(requestQuery)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	t.Log("ContinuationToken:", pageAssetXO.ContinuationToken)
	t.Log("Items Len:", len(pageAssetXO.Items))
}
