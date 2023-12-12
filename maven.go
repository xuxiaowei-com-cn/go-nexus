package nexus

import (
	"fmt"
	"net/http"
)

type MavenService struct {
	client *Client
}

// GetMavenGroupRepository 获取 Maven 分组 存储库
func (s *MavenService) GetMavenGroupRepository(repositoryName string, options ...RequestOptionFunc) (*SimpleApiGroupRepository, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/repositories/maven/group/%s", repositoryName)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
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

// GetMavenHostedRepository 获取 Maven 宿主 存储库
func (s *MavenService) GetMavenHostedRepository(repositoryName string, options ...RequestOptionFunc) (*MavenHostedApiRepository, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/repositories/maven/hosted/%s", repositoryName)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mavenHostedApiRepository *MavenHostedApiRepository
	resp, err := s.client.Do(req, &mavenHostedApiRepository)
	if err != nil {
		return nil, resp, err
	}

	return mavenHostedApiRepository, resp, nil
}

// GetMavenProxyRepository 获取 Maven 代理 存储库
func (s *MavenService) GetMavenProxyRepository(repositoryName string, options ...RequestOptionFunc) (*MavenProxyApiRepository, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/repositories/maven/proxy/%s", repositoryName)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mavenProxyApiRepository *MavenProxyApiRepository
	resp, err := s.client.Do(req, &mavenProxyApiRepository)
	if err != nil {
		return nil, resp, err
	}

	return mavenProxyApiRepository, resp, nil
}
