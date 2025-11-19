package nexus

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// Client Nexus REST SDK 客户端
// Nexus REST SDK client
//
// 字段说明/Fields:
// - BaseURL：Nexus 基地址（不含 basePath），例如 `http://localhost:8081` / Nexus base URL (without basePath), e.g. `http://localhost:8081`
// - Username：基本认证用户名 / Basic authentication username
// - Password：基本认证密码 / Basic authentication password
// - HTTPClient：底层 HTTP 客户端（可自定义超时、传输等）/ Underlying HTTP client (customizable timeout/transport)
// - MaxRetries：重试最大次数 / Maximum retry attempts
// - BackoffBase：退避最小等待时间 / Minimum backoff duration
// - BackoffMax：退避最大等待时间 / Maximum backoff duration
type Client struct {
	BaseURL     string
	Username    string
	Password    string
	HTTPClient  *http.Client
	MaxRetries  int
	BackoffBase time.Duration
	BackoffMax  time.Duration
}

// New 创建 SDK 客户端实例
func New(baseURL string) *Client {
	return &Client{
		BaseURL:     strings.TrimRight(baseURL, "/"),
		HTTPClient:  &http.Client{Timeout: 30 * time.Second},
		MaxRetries:  3,
		BackoffBase: 200 * time.Millisecond,
		BackoffMax:  2 * time.Second,
	}
}

// retryable HTTP 客户端（含退避与重试）
// Retryable HTTP client with backoff and retry
func (c *Client) retryClient() *retryablehttp.Client {
	rc := retryablehttp.NewClient()
	rc.Logger = nil
	rc.RetryMax = c.MaxRetries
	rc.RetryWaitMin = c.BackoffBase
	rc.RetryWaitMax = c.BackoffMax
	rc.HTTPClient = c.HTTPClient
	rc.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if err != nil {
			return true, nil
		}
		if resp == nil {
			return false, nil
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			return true, nil
		}
		if resp.StatusCode >= 500 {
			return true, nil
		}
		return false, nil
	}
	return rc
}

// 发送 HTTP 请求并处理响应，返回响应体字节
// Perform HTTP request and return response body bytes
func (c *Client) do(ctx context.Context, method, url string, body io.Reader) ([]byte, *http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, nil, err
	}
	if c.Username != "" || c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	rreq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.retryClient().Do(rreq)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode != http.StatusTooManyRequests && resp.StatusCode < 500 {
			if len(b) > 0 {
				return nil, resp, errors.New(string(b))
			}
		}
		return nil, resp, errors.New(resp.Status)
	}
	return b, resp, nil
}

func (c *Client) doWithHeaders(ctx context.Context, method, url string, body io.Reader, headers map[string]string) ([]byte, *http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, nil, err
	}
	if c.Username != "" || c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	} else if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	rreq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.retryClient().Do(rreq)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode != http.StatusTooManyRequests && resp.StatusCode < 500 {
			if len(b) > 0 {
				return nil, resp, errors.New(string(b))
			}
		}
		return nil, resp, errors.New(resp.Status)
	}
	return b, resp, nil
}
