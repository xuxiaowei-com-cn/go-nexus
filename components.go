package nexus

import (
	"fmt"
	"net/http"
)

type ComponentsService struct {
	client *Client
}

type ListComponentsRequest struct {
	Repository        string `json:"repository,omitempty" url:"repository,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty" url:"continuationToken,omitempty"`
}

// ListComponents 列出组件
func (s *ComponentsService) ListComponents(requestQuery *ListComponentsRequest) (*PageComponentXO, *Response, error) {

	u := "service/rest/v1/components"

	req, err := s.client.NewRequest(http.MethodGet, u, requestQuery, nil)
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

// GetComponents 获取组件详细信息
func (s *ComponentsService) GetComponents(id string) (*ComponentXO, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/components/%s", id)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil)
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
