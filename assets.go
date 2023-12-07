package nexus

import (
	"net/http"
)

type AssetsService struct {
	client *Client
}

type ListAssetsRequest struct {
	Repository        string `json:"repository,omitempty" url:"repository,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty" url:"continuationToken,omitempty"`
}

// ListAssets 列出资产
func (s *AssetsService) ListAssets(requestQuery *ListAssetsRequest) (*PageAssetXO, *Response, error) {

	u := "service/rest/v1/assets"

	req, err := s.client.NewRequest(http.MethodGet, u, requestQuery, nil)
	if err != nil {
		return nil, nil, err
	}

	var pageAssetXO *PageAssetXO
	resp, err := s.client.Do(req, &pageAssetXO)
	if err != nil {
		return nil, resp, err
	}

	return pageAssetXO, resp, nil
}
