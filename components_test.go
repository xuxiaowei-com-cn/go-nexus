package nexus

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ListComponents(t *testing.T) {
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
		versions := []string{"2.7."}
		extensions := []string{""}
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

	page, err := c.ListComponents(ctx, name, "")
	require.NoError(t, err)
	require.NotNil(t, page)
	num := 0
	for _, asset := range page.Items {
		num++
		a, err := c.GetComponentByID(ctx, asset.ID)
		require.NoError(t, err)
		require.NotNil(t, a)

		err = c.DeleteComponentByID(ctx, asset.ID)
		require.NoError(t, err)

		if num > 3 {
			break
		}
	}
}

func Test_UploadComponents_Maven(t *testing.T) {
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

	{
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
	}

	folder := "./tmp"
	err := os.MkdirAll(folder, 0o755)
	require.NoError(t, err)

	t.Run("pom", func(t *testing.T) {
		groupId := "org.springframework.boot"
		artifactId := "spring-boot-dependencies"
		version := "2.7.14"
		extension := "pom"
		filePath := filepath.Join(folder, fmt.Sprintf("%s-%s.%s", artifactId, version, extension))

		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, extension)
			_, _, err := downloadToFile(u, filePath, "", "")
			require.NoError(t, err)
		}

		defer os.Remove(filePath)
		f, err := os.Open(filePath)
		require.NoError(t, err)
		defer f.Close()

		assets := UploadAssets{
			maven2: &UploadAssetMaven2{
				groupId:         groupId,
				artifactId:      artifactId,
				version:         version,
				asset1Extension: extension,
				asset1:          f,
			},
		}
		require.NoError(t, c.UploadComponents(ctx, hosted, assets))

		page, err := c.ListComponents(ctx, hosted, "")
		require.NoError(t, err)
		require.NotNil(t, page)
		var found *Component
		for i := range page.Items {
			it := page.Items[i]
			if it.Group == groupId && it.Name == artifactId && it.Version == version {
				found = &it
				break
			}
		}
		require.NotNil(t, found)
		component, err := c.GetComponentByID(ctx, found.ID)
		require.NoError(t, err)
		require.NotNil(t, component)

		err = c.DeleteComponentByID(ctx, found.ID)
		require.NoError(t, err)
	})

	t.Run("jar", func(t *testing.T) {
		groupId := "org.springframework.boot"
		artifactId := "spring-boot-starter-web"
		version := "2.7.14"
		extension := "jar"
		filePath := filepath.Join(folder, fmt.Sprintf("%s-%s.%s", artifactId, version, extension))

		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, extension)
			_, _, err := downloadToFile(u, filePath, "", "")
			require.NoError(t, err)
		}

		defer os.Remove(filePath)
		f, err := os.Open(filePath)
		require.NoError(t, err)
		defer f.Close()

		generatePom := true

		assets := UploadAssets{
			maven2: &UploadAssetMaven2{
				groupId:         groupId,
				artifactId:      artifactId,
				version:         version,
				generatePom:     &generatePom,
				asset1Extension: extension,
				asset1:          f,
			},
		}
		require.NoError(t, c.UploadComponents(ctx, hosted, assets))

		page, err := c.ListComponents(ctx, hosted, "")
		require.NoError(t, err)
		require.NotNil(t, page)
		var found *Component
		for i := range page.Items {
			it := page.Items[i]
			if it.Group == groupId && it.Name == artifactId && it.Version == version {
				found = &it
				break
			}
		}
		require.NotNil(t, found)
		component, err := c.GetComponentByID(ctx, found.ID)
		require.NoError(t, err)
		require.NotNil(t, component)

		err = c.DeleteComponentByID(ctx, found.ID)
		require.NoError(t, err)
	})

	t.Run("sources", func(t *testing.T) {
		groupId := "org.springframework.boot"
		artifactId := "spring-boot-starter-web"
		version := "2.7.14"
		extension := "jar"
		classifier := "sources"

		filePath := filepath.Join(folder, fmt.Sprintf("%s-%s-%s.%s", artifactId, version, classifier, extension))

		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, classifier, extension)
			_, _, err := downloadToFile(u, filePath, "", "")
			require.NoError(t, err)
		}

		defer os.Remove(filePath)
		f, err := os.Open(filePath)
		require.NoError(t, err)
		defer f.Close()

		assets := UploadAssets{
			maven2: &UploadAssetMaven2{
				groupId:          groupId,
				artifactId:       artifactId,
				version:          version,
				asset1Extension:  extension,
				asset1Classifier: classifier,
				asset1:           f,
			},
		}
		require.NoError(t, c.UploadComponents(ctx, hosted, assets))

		page, err := c.ListComponents(ctx, hosted, "")
		require.NoError(t, err)
		require.NotNil(t, page)
		var found *Component
		for i := range page.Items {
			it := page.Items[i]
			if it.Group == groupId && it.Name == artifactId && it.Version == version {
				found = &it
				break
			}
		}
		require.NotNil(t, found)
		component, err := c.GetComponentByID(ctx, found.ID)
		require.NoError(t, err)
		require.NotNil(t, component)

		err = c.DeleteComponentByID(ctx, found.ID)
		require.NoError(t, err)
	})

	t.Run("javadoc", func(t *testing.T) {
		groupId := "org.springframework.boot"
		artifactId := "spring-boot-starter-web"
		version := "2.7.14"
		extension := "jar"
		classifier := "javadoc"

		filePath := filepath.Join(folder, fmt.Sprintf("%s-%s-%s.%s", artifactId, version, classifier, extension))

		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, classifier, extension)
			_, _, err := downloadToFile(u, filePath, "", "")
			require.NoError(t, err)
		}

		defer os.Remove(filePath)
		f, err := os.Open(filePath)
		require.NoError(t, err)
		defer f.Close()

		assets := UploadAssets{
			maven2: &UploadAssetMaven2{
				groupId:          groupId,
				artifactId:       artifactId,
				version:          version,
				asset1Extension:  extension,
				asset1Classifier: classifier,
				asset1:           f,
			},
		}
		require.NoError(t, c.UploadComponents(ctx, hosted, assets))

		page, err := c.ListComponents(ctx, hosted, "")
		require.NoError(t, err)
		require.NotNil(t, page)
		var found *Component
		for i := range page.Items {
			it := page.Items[i]
			if it.Group == groupId && it.Name == artifactId && it.Version == version {
				found = &it
				break
			}
		}
		require.NotNil(t, found)
		component, err := c.GetComponentByID(ctx, found.ID)
		require.NoError(t, err)
		require.NotNil(t, component)

		err = c.DeleteComponentByID(ctx, found.ID)
		require.NoError(t, err)
	})

	t.Run("jar-sources-javadoc", func(t *testing.T) {
		groupId := "org.springframework.boot"
		artifactId := "spring-boot-starter-web"
		version := "2.7.10"
		extension := "jar"

		classifierSources := "sources"
		classifierJavadoc := "javadoc"

		filePathJar := filepath.Join(folder, fmt.Sprintf("%s-%s.%s", artifactId, version, extension))
		filePathSources := filepath.Join(folder, fmt.Sprintf("%s-%s-%s.%s", artifactId, version, classifierSources, extension))
		filePathJavadoc := filepath.Join(folder, fmt.Sprintf("%s-%s-%s.%s", artifactId, version, classifierJavadoc, extension))

		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, extension)
			_, _, err := downloadToFile(u, filePathJar, "", "")
			require.NoError(t, err)
		}
		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, classifierSources, extension)
			_, _, err := downloadToFile(u, filePathSources, "", "")
			require.NoError(t, err)
		}
		{
			baseUrl := "https://repo1.maven.org/maven2"
			u := fmt.Sprintf("%s/%s/%s/%s/%s-%s-%s.%s", baseUrl, strings.ReplaceAll(groupId, ".", "/"), artifactId, version, artifactId, version, classifierJavadoc, extension)
			_, _, err := downloadToFile(u, filePathJavadoc, "", "")
			require.NoError(t, err)
		}

		defer os.Remove(filePathJar)
		defer os.Remove(filePathSources)
		defer os.Remove(filePathJavadoc)

		fJar, err := os.Open(filePathJar)
		require.NoError(t, err)
		defer fJar.Close()

		fSources, err := os.Open(filePathSources)
		require.NoError(t, err)
		defer fSources.Close()

		fJavadoc, err := os.Open(filePathJavadoc)
		require.NoError(t, err)
		defer fJavadoc.Close()

		generatePom := true

		assets := UploadAssets{
			maven2: &UploadAssetMaven2{
				groupId:          groupId,
				artifactId:       artifactId,
				version:          version,
				generatePom:      &generatePom,
				asset1Extension:  extension,
				asset1:           fJar,
				asset2Classifier: classifierSources,
				asset2Extension:  extension,
				asset2:           fSources,
				asset3Classifier: classifierJavadoc,
				asset3Extension:  extension,
				asset3:           fJavadoc,
			},
		}
		require.NoError(t, c.UploadComponents(ctx, hosted, assets))

		page, err := c.ListComponents(ctx, hosted, "")
		require.NoError(t, err)
		require.NotNil(t, page)
		var found *Component
		for i := range page.Items {
			it := page.Items[i]
			if it.Group == groupId && it.Name == artifactId && it.Version == version {
				found = &it
				break
			}
		}
		require.NotNil(t, found)
		component, err := c.GetComponentByID(ctx, found.ID)
		require.NoError(t, err)
		require.NotNil(t, component)

		err = c.DeleteComponentByID(ctx, found.ID)
		require.NoError(t, err)
	})

	err = c.DeleteRepository(ctx, hosted)
	require.NoError(t, err)
}

func Test_UploadComponents_Yum(t *testing.T) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	hosted := fmt.Sprintf(TestRepositoriesNamePrefix+"yum-hosted-%d", time.Now().UnixNano())

	{
		hreq := YumHostedRepositoryApiRequest{
			Name:   hosted,
			Online: true,
			Storage: HostedStorageAttributes{
				BlobStoreName:               "default",
				StrictContentTypeValidation: true,
				WritePolicy:                 "allow_once",
			},
			Yum: YumAttributes{
				RepodataDepth: 1,
				DeployPolicy:  "STRICT",
			},
		}
		require.NoError(t, c.CreateYumHostedRepository(ctx, hreq))
	}

	folder := "./tmp"
	err := os.MkdirAll(folder, 0o755)
	require.NoError(t, err)

	fileName := "docker-ce-cli-23.0.0-1.el8.x86_64.rpm"
	filePath := filepath.Join(folder, fileName)

	{
		u := "https://mirrors.aliyun.com/docker-ce/linux/centos/8/x86_64/stable/Packages/" + fileName
		_, _, err := downloadToFile(u, filePath, "", "")
		require.NoError(t, err)
	}

	defer os.Remove(filePath)

	asset, err := os.Open(filePath)
	require.NoError(t, err)
	defer asset.Close()

	assets := UploadAssets{
		yum: &UploadAssetYum{
			directory:     "/test",
			asset:         asset,
			assetFilename: filepath.Base(filePath),
		},
	}
	require.NoError(t, c.UploadComponents(ctx, hosted, assets))

	page, err := c.ListComponents(ctx, hosted, "")
	require.NoError(t, err)
	require.NotNil(t, page)
	require.Len(t, page.Items, 1)
	require.Equal(t, filepath.Join(assets.yum.directory, fileName), page.Items[0].Assets[0].Path)

	err = c.DeleteRepository(ctx, hosted)
	require.NoError(t, err)
}

func downloadToFile(url string, filePath string, username string, password string) (int64, time.Duration, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, 0, err
	}
	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}
	client := &http.Client{Timeout: 300 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, 0, fmt.Errorf("下载失败: %s => 状态码 %d", url, resp.StatusCode)
	}
	f, err := os.Create(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	start := time.Now()
	n, err := io.Copy(f, resp.Body)
	dur := time.Since(start)
	return n, dur, err
}
