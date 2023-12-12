package nexus

import (
	"fmt"
	"net/http"
)

type ComponentsService struct {
	client *Client
}

type ListComponentsQuery struct {
	Repository        string `json:"repository,omitempty" url:"repository,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty" url:"continuationToken,omitempty"`
}

// ListComponents 列出组件
func (s *ComponentsService) ListComponents(requestQuery *ListComponentsQuery, options ...RequestOptionFunc) (*PageComponentXO, *Response, error) {

	u := "service/rest/v1/components"

	req, err := s.client.NewRequest(http.MethodGet, u, requestQuery, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var pageComponentXO *PageComponentXO
	resp, err := s.client.Do(req, &pageComponentXO)
	if err != nil {
		return nil, resp, err
	}

	return pageComponentXO, resp, nil
}

// DeleteComponents 删除组件
func (s *ComponentsService) DeleteComponents(id string, options ...RequestOptionFunc) (*Response, error) {

	u := fmt.Sprintf("service/rest/v1/components/%s", id)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, nil, options)
	if err != nil {
		return nil, err
	}

	var pageComponentXO *PageComponentXO
	resp, err := s.client.Do(req, &pageComponentXO)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetComponents 获取组件详细信息
func (s *ComponentsService) GetComponents(id string, options ...RequestOptionFunc) (*ComponentXO, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/components/%s", id)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var componentXO *ComponentXO
	resp, err := s.client.Do(req, &componentXO)
	if err != nil {
		return nil, resp, err
	}

	return componentXO, resp, nil
}
