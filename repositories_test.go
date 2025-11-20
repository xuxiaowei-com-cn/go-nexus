package nexus

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const TestRepositoriesNamePrefix = "test-nexus-"

// Test_CleanRepositories 清理测试仓库
func Test_CleanRepositories(t *testing.T) {
	base := os.Getenv("NEXUS_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:48081"
	}
	c := New(base)
	u := os.Getenv("NEXUS_USERNAME")
	p := os.Getenv("NEXUS_PASSWORD")
	c.Username = u
	c.Password = p
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// 列出仓库
	repos, err := c.ListRepositories(ctx)
	require.NoError(t, err)
	require.NotNil(t, repos)

	for _, repo := range repos {
		if strings.HasPrefix(repo.Name, TestRepositoriesNamePrefix) {
			// 清理测试仓库
			err := c.DeleteRepository(ctx, repo.Name)
			require.NoError(t, err)
		}
	}
}

func Test_ListRepositories(t *testing.T) {
	base := os.Getenv("NEXUS_BASE_URL")
	if base == "" {
		base = "http://127.0.0.1:48081"
	}
	c := New(base)
	u := os.Getenv("NEXUS_USERNAME")
	p := os.Getenv("NEXUS_PASSWORD")
	c.Username = u
	c.Password = p
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// 列出仓库
	repos, err := c.ListRepositories(ctx)
	require.NoError(t, err)
	require.NotNil(t, repos)

	for _, repo := range repos {
		// 获取仓库详情
		rep, err := c.GetRepository(ctx, repo.Name)
		require.NoError(t, err)
		require.NotNil(t, rep)
	}
}

func Test_CreateAptHostedRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"apt-hosted-%d", time.Now().UnixNano())
	keypair := os.Getenv("NEXUS_APT_KEYPAIR")
	passphrase := os.Getenv("NEXUS_APT_PASSPHRASE")
	if keypair == "" {
		keypair = "-"
	}
	req := AptHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Apt:        AptHostedRepositoriesAttributes{Distribution: "bookworm"},
		AptSigning: AptSigningRepositoriesAttributes{Keypair: keypair, Passphrase: passphrase},
	}

	require.NoError(t, c.CreateAptHostedRepository(ctx, req))

	repo, err := c.GetAptHostedRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, repo)
	require.Equal(t, name, repo.Name)
	require.Equal(t, "hosted", repo.Type)
	require.Equal(t, "apt", repo.Format)

	{
		// 获取仓库详情
		repo, err := c.GetAptHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Apt.Distribution = "trixie"
		// 更新仓库
		err := c.UpdateAptHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetAptHostedRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Apt.Distribution, repo.Apt.Distribution)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}

func Test_CreateAptProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"apt-proxy-%d", time.Now().UnixNano())
	req := AptProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "http://deb.debian.org/debian",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true, Connection: &HttpClientConnectionAttributes{Timeout: 60}},
		Apt:           AptProxyRepositoriesAttributes{Distribution: "bookworm", Flat: false},
	}

	require.NoError(t, c.CreateAptProxyRepository(ctx, req))

	repo, err := c.GetAptProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, repo)
	require.Equal(t, name, repo.Name)
	require.Equal(t, "proxy", repo.Type)
	require.Equal(t, "apt", repo.Format)

	{
		// 获取仓库详情
		repo, err := c.GetAptProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Apt.Distribution = "trixie"
		// 更新仓库
		err := c.UpdateAptProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetAptProxyRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Apt.Distribution, repo.Apt.Distribution)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}

func Test_CreateCargoGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"cargo-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"cargo-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"cargo-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := CargoHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreateCargoHostedRepository(ctx, hreq))

	preq := CargoProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://crates.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		Cargo:         CargoAttributes{RequireAuthentication: false},
	}
	require.NoError(t, c.CreateCargoProxyRepository(ctx, preq))

	greq := CargoGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy, hosted}},
		Cargo: CargoAttributes{RequireAuthentication: false},
	}
	require.NoError(t, c.CreateCargoGroupRepository(ctx, greq))

	got, err := c.GetCargoGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetCargoGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateCargoGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetCargoGroupRepository(ctx, group)
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

func Test_CreateCargoHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"cargo-hosted-%d", time.Now().UnixNano())
	req := CargoHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateCargoHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetCargoHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateCargoHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetCargoHostedRepository(ctx, name)
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

func Test_CreateCargoProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"cargo-proxy-%d", time.Now().UnixNano())
	req := CargoProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://crates.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		Cargo:         CargoAttributes{RequireAuthentication: false},
	}

	err := c.CreateCargoProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetCargoProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "cargo", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetCargoProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateCargoProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetCargoProxyRepository(ctx, name)
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

func Test_CreateCocoapodsProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"cocoapods-proxy-%d", time.Now().UnixNano())
	req := CocoapodsProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://cdn.cocoapods.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateCocoapodsProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetCocoapodsProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "cocoapods", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetCocoapodsProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateCocoapodsProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetCocoapodsProxyRepository(ctx, name)
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

func Test_CreateComposerProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"composer-proxy-%d", time.Now().UnixNano())
	req := ComposerProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://repo.packagist.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateComposerProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetComposerProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "composer", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetComposerProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateComposerProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetComposerProxyRepository(ctx, name)
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

func Test_CreateConanGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"conan-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"conan-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"conan-group-%d", time.Now().UnixNano())

	hreq := ConanHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreateConanHostedRepository(ctx, hreq))

	preq := ConanProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://center.conan.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		ConanProxy:    &ConanProxyAttributes{ConanVersion: "V2"},
	}
	require.NoError(t, c.CreateConanProxyRepository(ctx, preq))

	greq := ConanGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupDeployAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateConanGroupRepository(ctx, greq))

	got, err := c.GetConanGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetConanGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateConanGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetConanGroupRepository(ctx, group)
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

func Test_CreateConanHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"conan-hosted-%d", time.Now().UnixNano())
	req := ConanHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateConanHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetConanHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateConanHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetConanHostedRepository(ctx, name)
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

func Test_CreateConanProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"conan-proxy-%d", time.Now().UnixNano())
	req := ConanProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://center.conan.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		ConanProxy:    &ConanProxyAttributes{ConanVersion: "V1"},
	}

	err := c.CreateConanProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetConanProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "conan", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetConanProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateConanProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetConanProxyRepository(ctx, name)
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

func Test_CreateDockerGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"docker-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"docker-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"docker-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := DockerHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: DockerHostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
			LatestPolicy:                true,
		},
		Docker: DockerAttributes{V1Enabled: false, ForceBasicAuth: true, PathEnabled: true},
	}
	require.NoError(t, c.CreateDockerHostedRepository(ctx, hreq))

	preq := DockerProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://registry-1.docker.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		Docker:        DockerAttributes{V1Enabled: false, ForceBasicAuth: true, PathEnabled: true},
		DockerProxy:   DockerProxyAttributes{IndexType: "HUB", CacheForeignLayers: true},
	}
	require.NoError(t, c.CreateDockerProxyRepository(ctx, preq))

	greq := DockerGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group:  GroupDeployAttributes{MemberNames: []string{proxy, hosted}},
		Docker: DockerAttributes{V1Enabled: false, ForceBasicAuth: true, PathEnabled: true},
	}
	require.NoError(t, c.CreateDockerGroupRepository(ctx, greq))

	got, err := c.GetDockerGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetDockerGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateDockerGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetDockerGroupRepository(ctx, group)
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

func Test_CreateDockerHostedRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"docker-hosted-%d", time.Now().UnixNano())
	req := DockerHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: DockerHostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
			LatestPolicy:                true,
		},
		Docker: DockerAttributes{V1Enabled: false, ForceBasicAuth: true, PathEnabled: true},
	}

	err := c.CreateDockerHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetDockerHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateDockerHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetDockerHostedRepository(ctx, name)
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

func Test_CreateDockerProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"docker-proxy-%d", time.Now().UnixNano())
	req := DockerProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://registry-1.docker.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		Docker:        DockerAttributes{V1Enabled: false, ForceBasicAuth: true, PathEnabled: true},
		DockerProxy:   DockerProxyAttributes{IndexType: "HUB", CacheForeignLayers: true},
	}

	err := c.CreateDockerProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetDockerProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "docker", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetDockerProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateDockerProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetDockerProxyRepository(ctx, name)
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

func Test_CreateGitlfsHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"gitlfs-hosted-%d", time.Now().UnixNano())
	req := GitLfsHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	err := c.CreateGitlfsHostedRepository(ctx, req)
	require.NoError(t, err)

	repo, err := c.GetGitlfsHostedRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, repo)
	require.Equal(t, name, repo.Name)
	require.Equal(t, "hosted", repo.Type)

	{
		// 获取仓库详情
		repo, err := c.GetGitlfsHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateGitlfsHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetGitlfsHostedRepository(ctx, name)
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

func Test_CreateGoGroupRepository(t *testing.T) {
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

	proxy1 := fmt.Sprintf(TestRepositoriesNamePrefix+"go-proxy-%d", time.Now().UnixNano())
	proxy2 := fmt.Sprintf(TestRepositoriesNamePrefix+"go-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"go-group-%d", time.Now().UnixNano())

	preq1 := GolangProxyRepositoryApiRequest{
		Name:   proxy1,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://proxy.golang.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}
	require.NoError(t, c.CreateGoProxyRepository(ctx, preq1))

	preq2 := GolangProxyRepositoryApiRequest{
		Name:   proxy2,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://proxy.golang.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}
	require.NoError(t, c.CreateGoProxyRepository(ctx, preq2))

	greq := GolangGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy1, proxy2}},
	}
	require.NoError(t, c.CreateGoGroupRepository(ctx, greq))

	got, err := c.GetGoGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy1, proxy2}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetGoGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{proxy1}
		// 更新仓库
		err := c.UpdateGoGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetGoGroupRepository(ctx, group)
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

	// 清理：删除 proxy1、proxy2
	require.NoError(t, c.DeleteRepository(ctx, proxy1))
	require.NoError(t, c.DeleteRepository(ctx, proxy2))
}

func Test_CreateGoHostedRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"go-proxy-%d", time.Now().UnixNano())
	req := GolangProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://proxy.golang.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateGoProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetGoProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "go", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	req.Proxy.ContentMaxAge = 60
	err = c.UpdateGoProxyRepository(ctx, name, req)
	require.NoError(t, err)

	repo, err := c.GetGoProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, repo)
	require.Equal(t, req.Proxy.ContentMaxAge, repo.Proxy.ContentMaxAge)

	err = c.DeleteRepository(ctx, name)
	require.NoError(t, err)
}

func Test_CreateHelmHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"helm-hosted-%d", time.Now().UnixNano())
	req := HelmHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateHelmHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetHelmHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateHelmHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetHelmHostedRepository(ctx, name)
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

func Test_CreateHelmProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"helm-proxy-%d", time.Now().UnixNano())
	req := HelmProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://charts.bitnami.com/bitnami",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateHelmProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetHelmProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "helm", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetHelmProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateHelmProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetHelmProxyRepository(ctx, name)
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

func Test_CreateHuggingfaceProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"huggingface-proxy-%d", time.Now().UnixNano())
	req := HuggingFaceProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://huggingface.co",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateHuggingfaceProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetHuggingfaceProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "huggingface", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetHuggingfaceProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateHuggingfaceProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetHuggingfaceProxyRepository(ctx, name)
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

func Test_CreateMavenGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"maven-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"maven-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"maven-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := MavenHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Maven: MavenAttributes{
			VersionPolicy:      "MIXED",
			LayoutPolicy:       "STRICT",
			ContentDisposition: "ATTACHMENT",
		},
	}
	require.NoError(t, c.CreateMavenHostedRepository(ctx, hreq))

	preq := MavenProxyRepositoryApiRequest{
		Name:   proxy,
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
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributesWithPreemptiveAuth{Blocked: false, AutoBlock: true},
		Maven:         MavenAttributes{VersionPolicy: "MIXED", LayoutPolicy: "STRICT", ContentDisposition: "ATTACHMENT"},
	}
	require.NoError(t, c.CreateMavenProxyRepository(ctx, preq))

	// 创建 group 仓库，成员为上述两个仓库
	greq := MavenGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateMavenGroupRepository(ctx, greq))

	got, err := c.GetMavenGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetMavenGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateMavenGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetMavenGroupRepository(ctx, group)
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

func Test_CreateMavenHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"maven-hosted-%d", time.Now().UnixNano())
	req := MavenHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Maven: MavenAttributes{
			VersionPolicy:      "MIXED",
			LayoutPolicy:       "STRICT",
			ContentDisposition: "ATTACHMENT",
		},
	}

	// 创建仓库
	err := c.CreateMavenHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetMavenHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Maven.LayoutPolicy = "PERMISSIVE"
		// 更新仓库
		err := c.UpdateMavenHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetMavenHostedRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Maven.LayoutPolicy, repo.Maven.LayoutPolicy)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}

func Test_CreateMavenProxyRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"maven-proxy-%d", time.Now().UnixNano())
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

	got, err := c.GetMavenProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "maven2", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetMavenProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Maven.LayoutPolicy = "PERMISSIVE"
		// 更新仓库
		err := c.UpdateMavenProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetMavenProxyRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Maven.LayoutPolicy, repo.Maven.LayoutPolicy)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}

func Test_CreateNpmGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"npm-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"npm-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"npm-group-%d", time.Now().UnixNano())

	hreq := NpmHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreateNpmHostedRepository(ctx, hreq))

	preq := NpmProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://registry.npmjs.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}
	require.NoError(t, c.CreateNpmProxyRepository(ctx, preq))

	greq := NpmGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupDeployAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateNpmGroupRepository(ctx, greq))

	got, err := c.GetNpmGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetNpmGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateNpmGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetNpmGroupRepository(ctx, group)
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

func Test_CreateNpmHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"npm-hosted-%d", time.Now().UnixNano())
	req := NpmHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateNpmHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetNpmHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateNpmHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetNpmHostedRepository(ctx, name)
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

func Test_CreateNpmProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"npm-proxy-%d", time.Now().UnixNano())
	req := NpmProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://registry.npmjs.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateNpmProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetNpmProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "npm", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetNpmProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateNpmProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetNpmProxyRepository(ctx, name)
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

func Test_CreateNugetGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"nuget-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"nuget-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"nuget-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := NugetHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreateNugetHostedRepository(ctx, hreq))

	preq := NugetProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://api.nuget.org/v3/index.json",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		NugetProxy:    NugetAttributes{NugetVersion: "V3", QueryCacheItemMaxAge: 3600},
	}
	require.NoError(t, c.CreateNugetProxyRepository(ctx, preq))

	// 创建 group 仓库，成员为上述两个仓库
	greq := NugetGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateNugetGroupRepository(ctx, greq))

	got, err := c.GetNugetGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetNugetGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateNugetGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetNugetGroupRepository(ctx, group)
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

func Test_CreateNugetHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"nuget-hosted-%d", time.Now().UnixNano())
	req := NugetHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateNugetHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetNugetHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateNugetHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetNugetHostedRepository(ctx, name)
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

func Test_CreateNugetProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"nuget-proxy-%d", time.Now().UnixNano())
	req := NugetProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://api.nuget.org/v3/index.json",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		NugetProxy:    NugetAttributes{NugetVersion: "V3", QueryCacheItemMaxAge: 3600},
	}

	err := c.CreateNugetProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetNugetProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "nuget", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetNugetProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateNugetProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetNugetProxyRepository(ctx, name)
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

func Test_CreateP2ProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"p2-proxy-%d", time.Now().UnixNano())
	req := P2ProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://download.eclipse.org/releases/latest",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateP2ProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetP2ProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "p2", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetP2ProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateP2ProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetP2ProxyRepository(ctx, name)
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

func Test_CreatePypiGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"pypi-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"pypi-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"pypi-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := PypiHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreatePypiHostedRepository(ctx, hreq))

	preq := PypiProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://pypi.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		Pypi:          PyPiProxyAttributes{RemoveQuarantined: false},
	}
	require.NoError(t, c.CreatePypiProxyRepository(ctx, preq))

	// 创建 group 仓库，成员为上述两个仓库
	greq := PypiGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Group: GroupDeployAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreatePypiGroupRepository(ctx, greq))

	got, err := c.GetPypiGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetPypiGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdatePypiGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetPypiGroupRepository(ctx, group)
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

func Test_CreatePypiHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"pypi-hosted-%d", time.Now().UnixNano())
	req := PypiHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreatePypiHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetPypiHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdatePypiHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetPypiHostedRepository(ctx, name)
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

func Test_CreatePypiProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"pypi-proxy-%d", time.Now().UnixNano())
	req := PypiProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://pypi.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
		Pypi:          PyPiProxyAttributes{RemoveQuarantined: false},
	}

	err := c.CreatePypiProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetPypiProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "pypi", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetPypiProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdatePypiProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetPypiProxyRepository(ctx, name)
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

func Test_CreateRGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"r-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"r-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"r-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := RHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreateRHostedRepository(ctx, hreq))

	preq := RProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://cloud.r-project.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}
	require.NoError(t, c.CreateRProxyRepository(ctx, preq))

	// 创建 group 仓库，成员为上述两个仓库
	greq := RGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateRGroupRepository(ctx, greq))

	got, err := c.GetRGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetRGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateRGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRGroupRepository(ctx, group)
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

func Test_CreateRHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"r-hosted-%d", time.Now().UnixNano())
	req := RHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateRHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetRHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateRHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRHostedRepository(ctx, name)
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

func Test_CreateRProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"r-proxy-%d", time.Now().UnixNano())
	req := RProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://cloud.r-project.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateRProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetRProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "r", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetRProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateRProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRProxyRepository(ctx, name)
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

func Test_CreateRawHostedRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"raw-hosted-%d", time.Now().UnixNano())
	req := RawHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Raw: RawAttributes{ContentDisposition: "ATTACHMENT"},
	}

	require.NoError(t, c.CreateRawHostedRepository(ctx, req))

	repo, err := c.GetRawHostedRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, repo)
	require.Equal(t, name, repo.Name)
	require.Equal(t, "hosted", repo.Type)
	require.Equal(t, "raw", repo.Format)

	{
		// 获取仓库详情
		repo, err := c.GetRawHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Raw.ContentDisposition = "INLINE"
		// 更新仓库
		err := c.UpdateRawHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRawHostedRepository(ctx, name)
			require.NoError(t, err)
			require.NotNil(t, repo)
			require.Equal(t, req.Raw.ContentDisposition, repo.Raw.ContentDisposition)
		}

		{
			// 删除仓库
			err := c.DeleteRepository(ctx, name)
			require.NoError(t, err)
		}
	}
}

func Test_CreateRubygemsGroupRepository(t *testing.T) {
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

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"rubygems-hosted-%d", time.Now().UnixNano())
	proxy := fmt.Sprintf(TestRepositoriesNamePrefix+"rubygems-proxy-%d", time.Now().UnixNano())
	group := fmt.Sprintf(TestRepositoriesNamePrefix+"rubygems-group-%d", time.Now().UnixNano())

	// 先创建成员仓库：hosted、proxy
	hreq := RubyGemsHostedRepositoryApiRequest{
		Name:   hosted,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
	require.NoError(t, c.CreateRubygemsHostedRepository(ctx, hreq))

	preq := RubyGemsProxyRepositoryApiRequest{
		Name:   proxy,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://rubygems.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}
	require.NoError(t, c.CreateRubygemsProxyRepository(ctx, preq))

	// 创建 group 仓库，成员为上述两个仓库
	greq := RubyGemsGroupRepositoryApiRequest{
		Name:   group,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Group: GroupAttributes{MemberNames: []string{proxy, hosted}},
	}
	require.NoError(t, c.CreateRubygemsGroupRepository(ctx, greq))

	got, err := c.GetRubygemsGroupRepository(ctx, group)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, group, got.Name)
	require.Equal(t, "group", got.Type)
	require.ElementsMatch(t, []string{proxy, hosted}, got.Group.MemberNames)

	{
		// 获取仓库详情
		repo, err := c.GetRubygemsGroupRepository(ctx, group)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, group, repo.Name)
		require.Equal(t, "group", repo.Type)
	}

	{
		greq.Group.MemberNames = []string{hosted}
		// 更新仓库
		err := c.UpdateRubygemsGroupRepository(ctx, group, greq)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRubygemsGroupRepository(ctx, group)
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

func Test_CreateRubygemsHostedRepository(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"rubygems-hosted-%d", time.Now().UnixNano())
	req := RubyGemsHostedRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: HostedStorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}

	// 创建仓库
	err := c.CreateRubygemsHostedRepository(ctx, req)
	require.NoError(t, err)

	{
		// 获取仓库详情
		repo, err := c.GetRubygemsHostedRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "hosted", repo.Type)
	}

	{
		req.Storage.WritePolicy = "allow"
		// 更新仓库
		err := c.UpdateRubygemsHostedRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRubygemsHostedRepository(ctx, name)
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

func Test_CreateRubygemsProxyRepository(t *testing.T) {
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

	name := fmt.Sprintf(TestRepositoriesNamePrefix+"rubygems-proxy-%d", time.Now().UnixNano())
	req := RubyGemsProxyRepositoryApiRequest{
		Name:   name,
		Online: true,
		Storage: StorageAttributes{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: ProxyAttributes{
			RemoteURL:      "https://rubygems.org",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: NegativeCacheAttributes{Enabled: true, TimeToLive: 1440},
		HttpClient:    HttpClientAttributes{Blocked: false, AutoBlock: true},
	}

	err := c.CreateRubygemsProxyRepository(ctx, req)
	require.NoError(t, err)

	got, err := c.GetRubygemsProxyRepository(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
	require.Equal(t, "proxy", got.Type)
	require.Equal(t, "rubygems", got.Format)
	require.Equal(t, req.Proxy.RemoteURL, got.Proxy.RemoteURL)

	{
		// 获取仓库详情
		repo, err := c.GetRubygemsProxyRepository(ctx, name)
		require.NoError(t, err)
		require.NotNil(t, repo)
		require.Equal(t, name, repo.Name)
		require.Equal(t, "proxy", repo.Type)
	}

	{
		req.Proxy.ContentMaxAge = 60
		// 更新仓库
		err := c.UpdateRubygemsProxyRepository(ctx, name, req)
		require.NoError(t, err)

		{
			// 获取仓库详情
			repo, err := c.GetRubygemsProxyRepository(ctx, name)
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
