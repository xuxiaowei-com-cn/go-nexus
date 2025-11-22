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

	maven2 := assets.Maven2
	if maven2 != nil {
		if maven2.GroupId == "" || maven2.ArtifactId == "" || maven2.Version == "" {
			return fmt.Errorf("groupId, artifactId and version are required")
		}

		_ = w.WriteField("maven2.groupId", maven2.GroupId)
		_ = w.WriteField("maven2.artifactId", maven2.ArtifactId)
		_ = w.WriteField("maven2.version", maven2.Version)

		if maven2.GeneratePom != nil {
			_ = w.WriteField("maven2.generate-pom", strconv.FormatBool(*maven2.GeneratePom))
		}
		if maven2.Packaging != "" {
			_ = w.WriteField("maven2.packaging", maven2.Packaging)
		}

		if maven2.Asset1Classifier != "" {
			_ = w.WriteField("maven2.asset1.classifier", maven2.Asset1Classifier)
		}
		if maven2.Asset1Extension != "" {
			_ = w.WriteField("maven2.asset1.extension", maven2.Asset1Extension)
		}

		if maven2.Asset2Classifier != "" {
			_ = w.WriteField("maven2.asset2.classifier", maven2.Asset2Classifier)
		}
		if maven2.Asset2Extension != "" {
			_ = w.WriteField("maven2.asset2.extension", maven2.Asset2Extension)
		}

		if maven2.Asset3Classifier != "" {
			_ = w.WriteField("maven2.asset3.classifier", maven2.Asset3Classifier)
		}
		if maven2.Asset3Extension != "" {
			_ = w.WriteField("maven2.asset3.extension", maven2.Asset3Extension)
		}

		if maven2.Asset1 != nil {
			fn := maven2.ArtifactId + "-" + maven2.Version + "." + maven2.Asset1Extension
			fw, err := w.CreateFormFile("maven2.asset1", fn)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, maven2.Asset1); err != nil {
				return err
			}
		}

		if maven2.Asset2 != nil {
			fn := maven2.ArtifactId + "-" + maven2.Version + "." + maven2.Asset2Extension
			fw, err := w.CreateFormFile("maven2.asset2", fn)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, maven2.Asset2); err != nil {
				return err
			}
		}

		if maven2.Asset3 != nil {
			fn := maven2.ArtifactId + "-" + maven2.Version + "." + maven2.Asset3Extension
			fw, err := w.CreateFormFile("maven2.asset3", fn)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, maven2.Asset3); err != nil {
				return err
			}
		}
	}

	yum := assets.Yum
	if yum != nil {
		if yum.AssetFilename == "" {
			return fmt.Errorf("assetFilename are required")
		}

		_ = w.WriteField("yum.asset.filename", yum.AssetFilename)

		if yum.Directory != "" {
			_ = w.WriteField("yum.directory", yum.Directory)
		}

		if yum.Asset != nil {
			fw, err := w.CreateFormFile("yum.asset", yum.AssetFilename)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, yum.Asset); err != nil {
				return err
			}
		}
	}

	raw := assets.Raw
	if raw != nil {
		if raw.Asset1Filename == "" {
			return fmt.Errorf("asset1Filename are required")
		}

		_ = w.WriteField("raw.asset1.filename", raw.Asset1Filename)

		if raw.Directory != "" {
			_ = w.WriteField("raw.directory", raw.Directory)
		}

		if raw.Asset1 != nil {
			fw, err := w.CreateFormFile("raw.asset1", raw.Asset1Filename)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, raw.Asset1); err != nil {
				return err
			}
		}

		if raw.Asset2Filename != "" {
			_ = w.WriteField("raw.asset2.filename", raw.Asset2Filename)
		}
		if raw.Asset2 != nil {
			fw, err := w.CreateFormFile("raw.asset2", raw.Asset2Filename)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, raw.Asset2); err != nil {
				return err
			}
		}

		if raw.Asset3Filename != "" {
			_ = w.WriteField("raw.asset3.filename", raw.Asset3Filename)
		}
		if raw.Asset3 != nil {
			fw, err := w.CreateFormFile("raw.asset3", raw.Asset3Filename)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fw, raw.Asset3); err != nil {
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
