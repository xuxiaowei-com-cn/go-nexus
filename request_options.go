package nexus

import "github.com/hashicorp/go-retryablehttp"

func WithHeader(name, value string) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		req.Header.Set(name, value)
		return nil
	}
}
