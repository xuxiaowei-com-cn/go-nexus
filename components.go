package nexus

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
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

// UploadComponents 上传单个组件（Swagger: POST /v1/components, operationId: uploadComponent）
// Upload a single component (Swagger: POST /v1/components, operationId: uploadComponent)
func (c *Client) UploadComponents(ctx context.Context, repository string, assets UploadAssets) error {

	q := url.Values{}
	q.Set("repository", repository)
	u := c.BaseURL + "/service/rest/v1/components" + "?" + q.Encode()

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	maven2 := assets.maven2

	if maven2 != nil {
		if maven2.groupId == "" || maven2.artifactId == "" || maven2.version == "" {
			return fmt.Errorf("groupId, artifactId and version are required")
		}

		if maven2.groupId != "" {
			_ = w.WriteField("maven2.groupId", maven2.groupId)
		}
		if maven2.artifactId != "" {
			_ = w.WriteField("maven2.artifactId", maven2.artifactId)
		}
		if maven2.version != "" {
			_ = w.WriteField("maven2.version", maven2.version)
		}

		if maven2.generatePom != nil {
			_ = w.WriteField("maven2.generate-pom", strconv.FormatBool(*maven2.generatePom))
		}
		if maven2.packaging != "" {
			_ = w.WriteField("maven2.packaging", maven2.packaging)
		}

		if maven2.asset1Classifier != "" {
			_ = w.WriteField("maven2.asset1.classifier", maven2.asset1Classifier)
		}
		if maven2.asset1Extension != "" {
			_ = w.WriteField("maven2.asset1.extension", maven2.asset1Extension)
		}

		if maven2.asset2Classifier != "" {
			_ = w.WriteField("maven2.asset2.classifier", maven2.asset2Classifier)
		}
		if maven2.asset2Extension != "" {
			_ = w.WriteField("maven2.asset2.extension", maven2.asset2Extension)
		}

		if maven2.asset3Classifier != "" {
			_ = w.WriteField("maven2.asset3.classifier", maven2.asset3Classifier)
		}
		if maven2.asset3Extension != "" {
			_ = w.WriteField("maven2.asset3.extension", maven2.asset3Extension)
		}

		if maven2.asset1 != nil {
			fn := maven2.artifactId + "-" + maven2.version + "." + maven2.asset1Extension
			fw, err := w.CreateFormFile("maven2.asset1", fn)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, maven2.asset1); err != nil {
				return err
			}
		}

		if maven2.asset2 != nil {
			fn := maven2.artifactId + "-" + maven2.version + "." + maven2.asset2Extension
			fw, err := w.CreateFormFile("maven2.asset2", fn)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, maven2.asset2); err != nil {
				return err
			}
		}

		if maven2.asset3 != nil {
			fn := maven2.artifactId + "-" + maven2.version + "." + maven2.asset3Extension
			fw, err := w.CreateFormFile("maven2.asset3", fn)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, maven2.asset3); err != nil {
				return err
			}
		}
	}

	_ = w.Close()

	headers := map[string]string{
		"Content-Type": w.FormDataContentType(),
		"Accept":       "application/json",
	}
	_, _, err := c.doWithHeaders(ctx, http.MethodPost, u, &buf, headers)
	return err
}
