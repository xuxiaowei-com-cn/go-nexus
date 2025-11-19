package nexus

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// Search 搜索组件（Swagger: GET /v1/search）
// Search components (Swagger: GET /v1/search)
//
// 查询参数/Query params:
// - q/name/group/version/repository/format 等均可选 / all optional
// 返回体/Response:
// - PageComponent（Swagger: PageComponentXO）
func (c *Client) Search(ctx context.Context, params map[string]string) (*PageComponent, error) {
	q := url.Values{}
	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	u := c.BaseURL + "/service/rest/v1/search"
	if len(q) > 0 {
		u += "?" + q.Encode()
	}
	b, _, err := c.do(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	var page PageComponent
	if err := json.Unmarshal(b, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// SearchAssets 搜索资产（Swagger: GET /v1/search/assets）
// Search assets (Swagger: GET /v1/search/assets)
// 返回体/Response:
// - PageAsset（Swagger: PageAssetXO）
func (c *Client) SearchAssets(ctx context.Context, params map[string]string) (*PageAsset, error) {
	q := url.Values{}
	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	u := c.BaseURL + "/service/rest/v1/search/assets"
	if len(q) > 0 {
		u += "?" + q.Encode()
	}
	b, _, err := c.do(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	var page PageAsset
	if err := json.Unmarshal(b, &page); err != nil {
		return nil, err
	}
	return &page, nil
}
