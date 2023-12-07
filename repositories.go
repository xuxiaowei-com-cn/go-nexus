package nexus

import (
	"net/http"
)

type RepositoryService struct {
	client *Client
}

// ListRepository 列出存储库
func (s *RepositoryService) ListRepository() ([]*RepositoryXO, *Response, error) {

	u := "service/rest/v1/repositories"

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
