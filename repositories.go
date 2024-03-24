package nexus

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type RepositoryService struct {
	client *Client
}

// ListRepository 列出存储库
func (s *RepositoryService) ListRepository(options ...RequestOptionFunc) ([]*RepositoryXO, *Response, error) {

	u := "service/rest/v1/repositories"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var repositoryXO []*RepositoryXO
	resp, err := s.client.Do(req, &repositoryXO)
	if err != nil {
		return nil, resp, err
	}

	return repositoryXO, resp, nil
}

func (s *RepositoryService) UploadFolder(folder string, repositoryName string, options ...RequestOptionFunc) error {

	mavenHostedApiRepository, _, err := s.client.Maven.GetMavenHostedRepository(repositoryName)
	if err != nil {
		return err
	}

	var versionPolicy = mavenHostedApiRepository.Maven.VersionPolicy

	u := "repository/" + repositoryName

	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// 跳过文件夹
			return nil
		}

		fileName := filepath.Base(path)

		if fileName == "desktop.ini" || fileName == "_remote.repositories" ||
			strings.HasSuffix(fileName, ".lastUpdated") || strings.HasSuffix(fileName, ".sha1") {
			return nil
		}

		if versionPolicy != "SNAPSHOT" && strings.Contains(fileName, "snapshots") || strings.Contains(fileName, "SNAPSHOTS") {
			return nil
		}

		filePath := path[len(folder)+1:]
		//tmp := filePath[:len(filePath)-len(fileName)-1]

		log.Printf("上传 文件 %s 开始", path)

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		url := u + "/" + strings.ReplaceAll(filePath, "\\", "/")

		req, err := s.client.UploadRequest(http.MethodPut, url, file, fileName, options)
		if err != nil {
			return err
		}

		resp, err := s.client.Do(req, nil)
		log.Println(resp.Status)
		if err != nil {

			req, err = s.client.NewRequest(http.MethodGet, url, nil, file, options)
			if err != nil {
				return err
			}

			resp, err = s.client.Do(req, nil)
			if err != nil {
				return err
			}

			if resp.StatusCode == 200 {
				log.Println("文件已存在")
				return nil
			}

			return err
		}

		log.Printf("上传 文件 %s 结束", path)

		return nil
	})

	return err
}
