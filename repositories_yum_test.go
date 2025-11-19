package nexus

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_CreateYumGroupRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"yum-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"yum-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"yum-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := YumHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Yum: YumAttributes{RepodataDepth: 1, DeployPolicy: "STRICT"},
	}
	require.NoError(t, c.CreateYumHostedRepository(ctx, hreq))

	preq := YumProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://repo.almalinux.org/almalinux/9/BaseOS/x86_64/os/",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}
	require.NoError(t, c.CreateYumProxyRepository(ctx, preq))

	// 创建 group 仓库，成员为上述两个仓库
	greq := YumGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateYumGroupRepository(ctx, greq))

	got, err := c.GetYumGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetYumGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateYumGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetYumGroupRepository(ctx, group)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, greq.Group.MemberNames, repo.Group.MemberNames)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, group)
			require.NoError(t, err)
		}
	}

	// 清理：删除 proxy、hosted
	require.NoError(t, c.DeleteRepository(ctx, proxy))
	require.NoError(t, c.DeleteRepository(ctx, hosted))
}

func Test_CreateYumHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"yum-hosted-%d", time.Now().UnixNano())
	req := YumHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Yum: YumAttributes{RepodataDepth: 1, DeployPolicy: "STRICT"},
	}

	// 创建仓库
	err := c.CreateYumHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetYumHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateYumHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetYumHostedRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Storage.WritePolicy, repo.Storage.WritePolicy)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}

func Test_CreateYumProxyRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"yum-proxy-%d", time.Now().UnixNano())
	req := YumProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://repo.almalinux.org/almalinux/9/BaseOS/x86_64/os/",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateYumProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetYumProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "yum", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetYumProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateYumProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetYumProxyRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Proxy.ContentMaxAge, repo.Proxy.ContentMaxAge)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}
