package nexus

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ListComponents 列出组件（Swagger: GET /v1/components）
// List components (Swagger: GET /v1/components)
//
// 查询参数/Query params:
// - repository：必填 / required
// - continuationToken：可选 / optional
// 返回体/Response:
// - PageComponent（Swagger: PageComponentXO），包含 items 与 continuationToken
func (c *Client) ListComponents(ctx context.Context, repository string, continuationToken string) (*PageComponent, error) {
	q := url.Values{}
	q.Set("repository", repository)
	if continuationToken != "" {
		q.Set("continuationToken", continuationToken)
	}
	u := c.BaseURL + "/service/rest/v1/components" + "?" + q.Encode()
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

// GetComponentByID 获取单个组件（Swagger: GET /v1/components/{id}）
// Get a single component (Swagger: GET /v1/components/{id})
func (c *Client) GetComponentByID(ctx context.Context, id string) (*Component, error) {
	u := c.BaseURL + "/service/rest/v1/components/" + id
	b, _, err := c.do(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	var comp Component
	if err := json.Unmarshal(b, &comp); err != nil {
		return nil, err
	}
	return &comp, nil
}

// DeleteComponentByID 删除单个组件（Swagger: DELETE /v1/components/{id}）
// Delete a single component (Swagger: DELETE /v1/components/{id})
func (c *Client) DeleteComponentByID(ctx context.Context, id string) error {
	u := c.BaseURL + "/service/rest/v1/components/" + id
	_, _, err := c.do(ctx, http.MethodDelete, u, nil)
	return err
}
