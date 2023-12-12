package nexus

import (
	"fmt"
	"net/http"
)

type BlobStoresService struct {
	client *Client
}

// GetBlobStoresQuotaStatus 获取给定blob存储区的配额状态
func (s *BlobStoresService) GetBlobStoresQuotaStatus(name string, options ...RequestOptionFunc) (*BlobStoreQuotaResultXO, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/blobstores/%s/quota-status", name)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var blobStoreQuotaResultXO *BlobStoreQuotaResultXO
	resp, err := s.client.Do(req, &blobStoreQuotaResultXO)
	if err != nil {
		return nil, resp, err
	}

	return blobStoreQuotaResultXO, resp, nil
}

// ListBlobStores 列出blob存储区
func (s *BlobStoresService) ListBlobStores(options ...RequestOptionFunc) ([]*GenericBlobStoreApiResponse, *Response, error) {

	u := "service/rest/v1/blobstores"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var genericBlobStoreApiResponses []*GenericBlobStoreApiResponse
	resp, err := s.client.Do(req, &genericBlobStoreApiResponses)
	if err != nil {
		return nil, resp, err
	}

	return genericBlobStoreApiResponses, resp, nil
}

// GetBlobStoresFile 按名称获取文件blob存储区配置
func (s *BlobStoresService) GetBlobStoresFile(name string, options ...RequestOptionFunc) (*FileBlobStoreApiModel, *Response, error) {

	u := fmt.Sprintf("service/rest/v1/blobstores/file/%s", name)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var fileBlobStoreApiModel *FileBlobStoreApiModel
	resp, err := s.client.Do(req, &fileBlobStoreApiModel)
	if err != nil {
		return nil, resp, err
	}

	return fileBlobStoreApiModel, resp, nil
}
