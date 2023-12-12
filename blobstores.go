package nexus

import "net/http"

type BlobStoresService struct {
	client *Client
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
