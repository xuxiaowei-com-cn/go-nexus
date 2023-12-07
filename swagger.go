package nexus

import (
	"net/http"
)

type Swagger struct {
	Swagger  string `json:"swagger,omitempty"`
	Info     Info   `json:"info,omitempty"`
	BasePath string `json:"basePath,omitempty"`
}

type Info struct {
	Version string `json:"version,omitempty"`
	Title   string `json:"title,omitempty"`
}

type SwaggerService struct {
	client *Client
}

// GetSwagger 获取 Swagger 配置
// 接口不需要凭证
func (s *SwaggerService) GetSwagger() (*Swagger, *Response, error) {

	u := "swagger.json"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var swagger *Swagger
	resp, err := s.client.Do(req, &swagger)
	if err != nil {
		return nil, resp, err
	}

	return swagger, resp, nil
}
