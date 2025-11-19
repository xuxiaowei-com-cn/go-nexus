package nexus

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ListAssets(t *testing.T) {
	base := os.Getenv("NEXUS_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:48081"
	}
	u := os.Getenv("NEXUS_USERNAME")
	p := os.Getenv("NEXUS_PASSWORD")
	if u == "" || p == "" {
		t.Skip("missing NEXUS_USERNAME or NEXUS_PASSWORD")
	}
	c := New(base)
	c.Username = u
	c.Password = p
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"maven-proxy-%d", time.Now().UnixNano())
	{
		req := MavenProxyRepositoryApiRequest{
			Name:   name,
			Online: true,
			Storage: StorageAttributes{
				BlobStoreName:               "default",
				StrictContentTypeValidation: true,
			},
			Proxy: ProxyAttributes{
				RemoteURL:      "https://repo1.maven.org/maven2",
				ContentMaxAge:  1440,
				MetadataMaxAge: 1440,
			},
			NegativeCache: NegativeCacheAttributes{
				Enabled:    true,
				TimeToLive: 1440,
			},
			HttpClient: HttpClientAttributesWithPreemptiveAuth{
				Blocked:   false,
				AutoBlock: true,
				Connection: &HttpClientConnectionAttributes{
					Timeout: 60,
				},
			},
			Maven: MavenAttributes{
				VersionPolicy:      "MIXED",
				LayoutPolicy:       "STRICT",
				ContentDisposition: "ATTACHMENT",
			},
		}

		err := c.CreateMavenProxyRepository(ctx, req)
		require.NoError(t, err)

		// 访问构件 URL 促使代理拉取
		versions := []string{"2.6.", "2.7."}
		extensions := []string{"", ".asc", ".md5", ".sha1"}
		for i := 0; i <= 15; i++ {
			for _, version := range versions {
				for _, extension := range extensions {
					u := fmt.Sprintf("%s/repository/%s/org/springframework/boot/spring-boot-dependencies/%s%d/spring-boot-dependencies-%s%d.pom%s", base, name, version, i, version, i, extension)
					req2, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
					require.NoError(t, err)
					if c.Username != "" || c.Password != "" {
						req2.SetBasicAuth(c.Username, c.Password)
					}
					c.HTTPClient.Do(req2)
					//resp, err := c.HTTPClient.Do(req2)
					//require.NoError(t, err)
					//if resp != nil {
					//	require.Equal(t, http.StatusOK, resp.StatusCode)
					//	resp.Body.Close()
					//}
				}
			}
		}
	}

	page, err := c.ListAssets(ctx, name, "")
	require.NoError(t, err)
	require.NotNil(t, page)
	num := 0
	for _, asset := range page.Items {
		num++
		a, err := c.GetAssetByID(ctx, asset.ID)
		require.NoError(t, err)
		require.NotNil(t, a)
		if num > 3 {
			break
		}
	}

	require.NotNil(t, page.ContinuationToken)
	{
		page, err := c.ListAssets(ctx, name, page.ContinuationToken)
		require.NoError(t, err)
		require.NotNil(t, page)
		num := 0
		for _, asset := range page.Items {
			num++
			a, err := c.GetAssetByID(ctx, asset.ID)
			require.NoError(t, err)
			require.NotNil(t, a)
			if num > 3 {
				break
			}
			err = c.DeleteAssetByID(ctx, asset.ID)
			require.NoError(t, err)
		}
	}

	// 清理代理仓库
	require.NoError(t, c.DeleteRepository(ctx, name))
}
