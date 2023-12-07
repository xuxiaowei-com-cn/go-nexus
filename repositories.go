package nexus

import (
	"net/http"
)

type RepositoryXO struct {
	Name       string      `json:"name,omitempty"`
	Format     string      `json:"format,omitempty"`
	Type       string      `json:"type,omitempty"`
	Url        string      `json:"url,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

type RepositoryService struct {
	client *Client
}

// ListRepository 列出存储库
func (s *RepositoryService) ListRepository() ([]*RepositoryXO, *Response, error) {

	u := "v1/repositories"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil)
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
