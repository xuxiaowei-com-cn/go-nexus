package nexus

type FileService struct {
	client *Client
}

// Download 文件下载
func (s *FileService) Download(method string, url string, path string, requestQuery interface{}, requestBody interface{}, options ...RequestOptionFunc) (*Response, error) {

	req, err := s.client.NewRequest(method, url, requestQuery, requestBody, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoDownload(req, path)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
