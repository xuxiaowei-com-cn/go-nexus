package nexus

import (
	"fmt"
	"net/http"
)

type AssetsService struct {
	client *Client
}

type ListAssetsQuery struct {
	Repository        string `json:"repository,omitempty" url:"repository,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty" url:"continuationToken,omitempty"`
}

// GetAssets 获取资产详细信息
func (s *AssetsService) GetAssets(id string, options ...RequestOptionFunc) (*AssetXO, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/assets/%s", id)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var assetXO *AssetXO
	resp, err := s.client.Do(req, &assetXO)
	if err != nil {
		return nil, resp, err
	}

	return assetXO, resp, nil
}

// DeleteAssets 删除资产
func (s *AssetsService) DeleteAssets(id string, options ...RequestOptionFunc) (*Response, error) {

	u := fmt.Sprintf("service/rest/v1/assets/%s", id)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, nil, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListAssets 列出资产
func (s *AssetsService) ListAssets(requestQuery *ListAssetsQuery, options ...RequestOptionFunc) (*PageAssetXO, *Response, error) {

	u := "service/rest/v1/assets"

	req, err := s.client.NewRequest(http.MethodGet, u, requestQuery, nil, options)
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
