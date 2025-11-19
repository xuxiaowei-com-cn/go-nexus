package nexus

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// DeleteAssetByID 删除单个资产（Swagger: DELETE /v1/assets/{id}）
// Delete a single asset (Swagger: DELETE /v1/assets/{id})
func (c *Client) DeleteAssetByID(ctx context.Context, id string) error {
	u := c.BaseURL + "/service/rest/v1/assets/" + id
	_, _, err := c.do(ctx, http.MethodDelete, u, nil)
	return err
}

// GetAssetByID 获取单个资产（Swagger: GET /v1/assets/{id}）
// Get a single asset (Swagger: GET /v1/assets/{id})
func (c *Client) GetAssetByID(ctx context.Context, id string) (*Asset, error) {
	u := c.BaseURL + "/service/rest/v1/assets/" + id
	b, _, err := c.do(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	var asset Asset
	if err := json.Unmarshal(b, &asset); err != nil {
		return nil, err
	}
	return &asset, nil
}

// ListAssets 列出资产（Swagger: GET /v1/assets）
// List assets (Swagger: GET /v1/assets)
//
// 查询参数/Query params:
// - repository：必填 / required
// - continuationToken：可选 / optional
// 返回体/Response:
// - PageAsset（Swagger: PageAssetXO），包含 items 与 continuationToken
func (c *Client) ListAssets(ctx context.Context, repository string, continuationToken string) (*PageAsset, error) {
	q := url.Values{}
	q.Set("repository", repository)
	if continuationToken != "" {
		q.Set("continuationToken", continuationToken)
	}
	u := c.BaseURL + "/service/rest/v1/assets" + "?" + q.Encode()
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
