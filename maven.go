package nexus

import (
	"fmt"
	"net/http"
)

type MavenService struct {
	client *Client
}

type MavenTypeValue string

const (
	MavenTypeGroup  MavenTypeValue = "group"
	MavenTypeProxy  MavenTypeValue = "proxy"
	MavenTypeHosted MavenTypeValue = "hosted"
)

// GetMavenRepository 获取 Maven 存储库
func (s *MavenService) GetMavenRepository(mavenTypeValue MavenTypeValue, repositoryName string) (*SimpleApiGroupRepository, *Response, error) {

	u := fmt.Sprintf("v1/repositories/maven/%s/%s", mavenTypeValue, repositoryName)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var simpleApiGroupRepository *SimpleApiGroupRepository
	resp, err := s.client.Do(req, &simpleApiGroupRepository)
	if err != nil {
		return nil, resp, err
	}

	return simpleApiGroupRepository, resp, nil
}
