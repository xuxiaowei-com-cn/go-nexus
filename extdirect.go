package nexus

import (
	"math/rand"
	"net/http"
)

type ExtDirectService struct {
	client *Client
}

type PostExtDirectRequest struct {
	Action string              `json:"action,omitempty"`
	Method string              `json:"method,omitempty"`
	Data   []PostExtDirectData `json:"data,omitempty"`
	Type   string              `json:"type,omitempty"`
	Tid    int                 `json:"tid,omitempty"`
}

type PostExtDirectData struct {
	RepositoryName string `json:"repositoryName,omitempty"`
	Node           string `json:"node,omitempty"`
}

type ExtDirectResultData struct {
	Id          string `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
	Type        string `json:"type,omitempty"`
	Leaf        bool   `json:"leaf,omitempty"`
	ComponentId string `json:"componentId,omitempty"`
	AssetId     string `json:"assetId,omitempty"`
	PackageUrl  string `json:"packageUrl,omitempty"`
}

type ExtDirectResult struct {
	Success bool                  `json:"success,omitempty"`
	Data    []ExtDirectResultData `json:"data,omitempty"`
}

type ExtDirect struct {
	Tid    int             `json:"tid,omitempty"`
	Action string          `json:"action,omitempty"`
	Method string          `json:"method,omitempty"`
	Result ExtDirectResult `json:"result,omitempty"`
	Type   string          `json:"type,omitempty"`
}

// PostExtDirect 搜索
func (s *ExtDirectService) PostExtDirect(requestBody *PostExtDirectRequest) (*ExtDirect, *Response, error) {

	u := "service/extdirect"

	if requestBody.Tid == 0 {
		requestBody.Tid = rand.Intn(10000)
	}

	req, err := s.client.NewRequest(http.MethodPost, u, nil, requestBody)
	if err != nil {
		return nil, nil, err
	}

	var extDirect *ExtDirect
	resp, err := s.client.Do(req, &extDirect)
	if err != nil {
		return nil, resp, err
	}

	return extDirect, resp, nil
}
