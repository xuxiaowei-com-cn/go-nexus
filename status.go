package nexus

import "net/http"

type StatusService struct {
	client *Client
}

// GetStatus 验证服务器是否可以响应读取请求的运行状况检查终结点
func (s *StatusService) GetStatus(options ...RequestOptionFunc) (*Response, error) {

	u := "service/rest/v1/status"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
