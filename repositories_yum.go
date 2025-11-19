package nexus

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func (c *Client) CreateYumGroupRepository(ctx context.Context, req YumGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

func (c *Client) CreateYumHostedRepository(ctx context.Context, req YumHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

func (c *Client) CreateYumProxyRepository(ctx context.Context, req YumProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

func (c *Client) GetYumGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/group/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiGroupRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func (c *Client) GetYumHostedRepository(ctx context.Context, repositoryName string) (*YumHostedApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo YumHostedApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func (c *Client) GetYumProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiProxyRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

func (c *Client) UpdateYumGroupRepository(ctx context.Context, repositoryName string, req YumGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

func (c *Client) UpdateYumHostedRepository(ctx context.Context, repositoryName string, req YumHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

func (c *Client) UpdateYumProxyRepository(ctx context.Context, repositoryName string, req YumProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/yum/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}
