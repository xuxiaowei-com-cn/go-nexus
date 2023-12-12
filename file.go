package nexus

import (
	"strings"
)

type FileService struct {
	client *Client
}

// Download 文件下载
func (s *FileService) Download(method string, url string, path string, requestQuery interface{}, requestBody interface{}, options ...RequestOptionFunc) (*Response, error) {

	var u = url
	if strings.HasPrefix(url, s.client.BaseURL().String()) {
		u = strings.TrimPrefix(url, s.client.BaseURL().String())
	}

	req, err := s.client.NewRequest(method, u, requestQuery, requestBody, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoDownload(req, path)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
