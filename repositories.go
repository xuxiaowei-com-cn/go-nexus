package nexus

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// DeleteRepository 删除仓库（Swagger: DELETE /v1/repositories/{repositoryName}）
// Delete repository (Swagger: DELETE /v1/repositories/{repositoryName})
func (c *Client) DeleteRepository(ctx context.Context, repositoryName string) error {
	url := c.BaseURL + "/service/rest/v1/repositories/" + repositoryName
	_, _, err := c.do(ctx, http.MethodDelete, url, nil)
	return err
}

// GetRepository 获取仓库详情（Swagger: GET /v1/repositories/{repositoryName}）
// Get repository details (Swagger: GET /v1/repositories/{repositoryName})
func (c *Client) GetRepository(ctx context.Context, repositoryName string) (*Repository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo Repository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// ListRepositories 列出仓库（Swagger: GET /v1/repositories）
// List repositories (Swagger: GET /v1/repositories)
//
// 说明/Notes:
// - 基于 Nexus REST 接口 `BaseURL + /service/rest/v1/repositories`
func (c *Client) ListRepositories(ctx context.Context) ([]Repository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories"
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repos []Repository
	if err := json.Unmarshal(b, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// CreateAptHostedRepository 创建 Apt hosted 仓库（Swagger: POST /v1/repositories/apt/hosted）
// Create Apt hosted repository (Swagger: POST /v1/repositories/apt/hosted)
func (c *Client) CreateAptHostedRepository(ctx context.Context, req AptHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/apt/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateAptProxyRepository 创建 Apt proxy 仓库（Swagger: POST /v1/repositories/apt/proxy）
// Create Apt proxy repository (Swagger: POST /v1/repositories/apt/proxy)
func (c *Client) CreateAptProxyRepository(ctx context.Context, req AptProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/apt/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetAptHostedRepository 获取 Apt hosted 仓库（Swagger: GET /v1/repositories/apt/hosted/{repositoryName}）
// Get Apt hosted repository (Swagger: GET /v1/repositories/apt/hosted/{repositoryName})
func (c *Client) GetAptHostedRepository(ctx context.Context, repositoryName string) (*AptHostedApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/apt/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo AptHostedApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetAptProxyRepository 获取 Apt proxy 仓库（Swagger: GET /v1/repositories/apt/proxy/{repositoryName}）
// Get Apt proxy repository (Swagger: GET /v1/repositories/apt/proxy/{repositoryName})
func (c *Client) GetAptProxyRepository(ctx context.Context, repositoryName string) (*AptProxyApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/apt/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo AptProxyApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateAptHostedRepository 更新 Apt hosted 仓库（Swagger: PUT /v1/repositories/apt/hosted/{repositoryName}）
// Update Apt hosted repository (Swagger: PUT /v1/repositories/apt/hosted/{repositoryName})
func (c *Client) UpdateAptHostedRepository(ctx context.Context, repositoryName string, req AptHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/apt/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateAptProxyRepository 更新 Apt proxy 仓库（Swagger: PUT /v1/repositories/apt/proxy/{repositoryName}）
// Update Apt proxy repository (Swagger: PUT /v1/repositories/apt/proxy/{repositoryName})
func (c *Client) UpdateAptProxyRepository(ctx context.Context, repositoryName string, req AptProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/apt/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateCargoGroupRepository 创建 Cargo group 仓库（Swagger: POST /v1/repositories/cargo/group）
// Create cargo group repository (Swagger: POST /v1/repositories/cargo/group)
func (c *Client) CreateCargoGroupRepository(ctx context.Context, req CargoGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateCargoHostedRepository 创建 Cargo hosted 仓库（Swagger: POST /v1/repositories/cargo/hosted）
// Create cargo hosted repository (Swagger: POST /v1/repositories/cargo/hosted)
func (c *Client) CreateCargoHostedRepository(ctx context.Context, req CargoHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateCargoProxyRepository 创建 Cargo proxy 仓库（Swagger: POST /v1/repositories/cargo/proxy）
// Create cargo proxy repository (Swagger: POST /v1/repositories/cargo/proxy)
func (c *Client) CreateCargoProxyRepository(ctx context.Context, req CargoProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetCargoGroupRepository 获取 Cargo group 仓库（Swagger: GET /v1/repositories/cargo/group/{repositoryName}）
// Get cargo group repository (Swagger: GET /v1/repositories/cargo/group/{repositoryName})
func (c *Client) GetCargoGroupRepository(ctx context.Context, repositoryName string) (*CargoGroupApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/group/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo CargoGroupApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetCargoHostedRepository 获取 Cargo hosted 仓库（Swagger: GET /v1/repositories/cargo/hosted/{repositoryName}）
// Get cargo hosted repository (Swagger: GET /v1/repositories/cargo/hosted/{repositoryName})
func (c *Client) GetCargoHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetCargoProxyRepository 获取 Cargo proxy 仓库（Swagger: GET /v1/repositories/cargo/proxy/{repositoryName}）
// Get cargo proxy repository (Swagger: GET /v1/repositories/cargo/proxy/{repositoryName})
func (c *Client) GetCargoProxyRepository(ctx context.Context, repositoryName string) (*CargoProxyApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo CargoProxyApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateCargoGroupRepository 更新 Cargo group 仓库（Swagger: PUT /v1/repositories/cargo/group/{repositoryName}）
// Update cargo group repository (Swagger: PUT /v1/repositories/cargo/group/{repositoryName})
func (c *Client) UpdateCargoGroupRepository(ctx context.Context, repositoryName string, req CargoGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateCargoHostedRepository 更新 Cargo hosted 仓库（Swagger: PUT /v1/repositories/cargo/hosted/{repositoryName}）
// Update cargo hosted repository (Swagger: PUT /v1/repositories/cargo/hosted/{repositoryName})
func (c *Client) UpdateCargoHostedRepository(ctx context.Context, repositoryName string, req CargoHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateCargoProxyRepository 更新 Cargo proxy 仓库（Swagger: PUT /v1/repositories/cargo/proxy/{repositoryName}）
// Update cargo proxy repository (Swagger: PUT /v1/repositories/cargo/proxy/{repositoryName})
func (c *Client) UpdateCargoProxyRepository(ctx context.Context, repositoryName string, req CargoProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cargo/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateCocoapodsProxyRepository 创建 Cocoapods proxy 仓库（Swagger: POST /v1/repositories/cocoapods/proxy）
// Create Cocoapods proxy repository (Swagger: POST /v1/repositories/cocoapods/proxy)
func (c *Client) CreateCocoapodsProxyRepository(ctx context.Context, req CocoapodsProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cocoapods/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetCocoapodsProxyRepository 获取 Cocoapods proxy 仓库（Swagger: GET /v1/repositories/cocoapods/proxy/{repositoryName}）
// Get Cocoapods proxy repository (Swagger: GET /v1/repositories/cocoapods/proxy/{repositoryName})
func (c *Client) GetCocoapodsProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/cocoapods/proxy/" + repositoryName
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

// UpdateCocoapodsProxyRepository 更新 Cocoapods proxy 仓库（Swagger: PUT /v1/repositories/cocoapods/proxy/{repositoryName}）
// Update Cocoapods proxy repository (Swagger: PUT /v1/repositories/cocoapods/proxy/{repositoryName})
func (c *Client) UpdateCocoapodsProxyRepository(ctx context.Context, repositoryName string, req CocoapodsProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/cocoapods/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateComposerProxyRepository 创建 Composer proxy 仓库（Swagger: POST /v1/repositories/composer/proxy）
// Create Composer proxy repository (Swagger: POST /v1/repositories/composer/proxy)
func (c *Client) CreateComposerProxyRepository(ctx context.Context, req ComposerProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/composer/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetComposerProxyRepository 获取 Composer proxy 仓库（Swagger: GET /v1/repositories/composer/proxy/{repositoryName}）
// Get Composer proxy repository (Swagger: GET /v1/repositories/composer/proxy/{repositoryName})
func (c *Client) GetComposerProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/composer/proxy/" + repositoryName
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

// UpdateComposerProxyRepository 更新 Composer proxy 仓库（Swagger: PUT /v1/repositories/composer/proxy/{repositoryName}）
// Update Composer proxy repository (Swagger: PUT /v1/repositories/composer/proxy/{repositoryName})
func (c *Client) UpdateComposerProxyRepository(ctx context.Context, repositoryName string, req ComposerProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/composer/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateConanGroupRepository 创建 Conan group 仓库（Swagger: POST /v1/repositories/conan/group）
// Create Conan group repository (Swagger: POST /v1/repositories/conan/group)
func (c *Client) CreateConanGroupRepository(ctx context.Context, req ConanGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateConanHostedRepository 创建 Conan hosted 仓库（Swagger: POST /v1/repositories/conan/hosted）
// Create Conan hosted repository (Swagger: POST /v1/repositories/conan/hosted)
func (c *Client) CreateConanHostedRepository(ctx context.Context, req ConanHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateConanProxyRepository 创建 Conan proxy 仓库（Swagger: POST /v1/repositories/conan/proxy）
// Create Conan proxy repository (Swagger: POST /v1/repositories/conan/proxy)
func (c *Client) CreateConanProxyRepository(ctx context.Context, req ConanProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetConanGroupRepository 获取 Conan group 仓库（Swagger: GET /v1/repositories/conan/group/{repositoryName}）
// Get Conan group repository (Swagger: GET /v1/repositories/conan/group/{repositoryName})
func (c *Client) GetConanGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupDeployRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/group/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiGroupDeployRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetConanHostedRepository 获取 Conan hosted 仓库（Swagger: GET /v1/repositories/conan/hosted/{repositoryName}）
// Get Conan hosted repository (Swagger: GET /v1/repositories/conan/hosted/{repositoryName})
func (c *Client) GetConanHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetConanProxyRepository 获取 Conan proxy 仓库（Swagger: GET /v1/repositories/conan/proxy/{repositoryName}）
// Get Conan proxy repository (Swagger: GET /v1/repositories/conan/proxy/{repositoryName})
func (c *Client) GetConanProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/proxy/" + repositoryName
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

// UpdateConanGroupRepository 更新 Conan group 仓库（Swagger: PUT /v1/repositories/conan/group/{repositoryName}）
// Update Conan group repository (Swagger: PUT /v1/repositories/conan/group/{repositoryName})
func (c *Client) UpdateConanGroupRepository(ctx context.Context, repositoryName string, req ConanGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateConanHostedRepository 更新 Conan hosted 仓库（Swagger: PUT /v1/repositories/conan/hosted/{repositoryName}）
// Update Conan hosted repository (Swagger: PUT /v1/repositories/conan/hosted/{repositoryName})
func (c *Client) UpdateConanHostedRepository(ctx context.Context, repositoryName string, req ConanHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateConanProxyRepository 更新 Conan proxy 仓库（Swagger: PUT /v1/repositories/conan/proxy/{repositoryName}）
// Update Conan proxy repository (Swagger: PUT /v1/repositories/conan/proxy/{repositoryName})
func (c *Client) UpdateConanProxyRepository(ctx context.Context, repositoryName string, req ConanProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/conan/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateDockerGroupRepository 创建 Docker group 仓库（Swagger: POST /v1/repositories/docker/group）
// Create Docker group repository (Swagger: POST /v1/repositories/docker/group)
func (c *Client) CreateDockerGroupRepository(ctx context.Context, req DockerGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateDockerHostedRepository 创建 Docker hosted 仓库（Swagger: POST /v1/repositories/docker/hosted）
// Create Docker hosted repository (Swagger: POST /v1/repositories/docker/hosted)
func (c *Client) CreateDockerHostedRepository(ctx context.Context, req DockerHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateDockerProxyRepository 创建 Docker proxy 仓库（Swagger: POST /v1/repositories/docker/proxy）
// Create Docker proxy repository (Swagger: POST /v1/repositories/docker/proxy)
func (c *Client) CreateDockerProxyRepository(ctx context.Context, req DockerProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetDockerGroupRepository 获取 Docker group 仓库（Swagger: GET /v1/repositories/docker/group/{repositoryName}）
// Get Docker group repository (Swagger: GET /v1/repositories/docker/group/{repositoryName})
func (c *Client) GetDockerGroupRepository(ctx context.Context, repositoryName string) (*DockerGroupApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/group/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo DockerGroupApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetDockerHostedRepository 获取 Docker hosted 仓库（Swagger: GET /v1/repositories/docker/hosted/{repositoryName}）
// Get Docker hosted repository (Swagger: GET /v1/repositories/docker/hosted/{repositoryName})
func (c *Client) GetDockerHostedRepository(ctx context.Context, repositoryName string) (*DockerHostedApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo DockerHostedApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetDockerProxyRepository 获取 Docker proxy 仓库（Swagger: GET /v1/repositories/docker/proxy/{repositoryName}）
// Get Docker proxy repository (Swagger: GET /v1/repositories/docker/proxy/{repositoryName})
func (c *Client) GetDockerProxyRepository(ctx context.Context, repositoryName string) (*DockerProxyApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo DockerProxyApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateDockerGroupRepository 更新 Docker group 仓库（Swagger: PUT /v1/repositories/docker/group/{repositoryName}）
// Update Docker group repository (Swagger: PUT /v1/repositories/docker/group/{repositoryName})
func (c *Client) UpdateDockerGroupRepository(ctx context.Context, repositoryName string, req DockerGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateDockerHostedRepository 更新 Docker hosted 仓库（Swagger: PUT /v1/repositories/docker/hosted/{repositoryName}）
// Update Docker hosted repository (Swagger: PUT /v1/repositories/docker/hosted/{repositoryName})
func (c *Client) UpdateDockerHostedRepository(ctx context.Context, repositoryName string, req DockerHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateDockerProxyRepository 更新 Docker proxy 仓库（Swagger: PUT /v1/repositories/docker/proxy/{repositoryName}）
// Update Docker proxy repository (Swagger: PUT /v1/repositories/docker/proxy/{repositoryName})
func (c *Client) UpdateDockerProxyRepository(ctx context.Context, repositoryName string, req DockerProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/docker/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateGitlfsHostedRepository 创建 Git LFS hosted 仓库（Swagger: POST /v1/repositories/gitlfs/hosted）
// Create Git LFS hosted repository (Swagger: POST /v1/repositories/gitlfs/hosted)
func (c *Client) CreateGitlfsHostedRepository(ctx context.Context, req GitLfsHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/gitlfs/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetGitlfsHostedRepository 获取 Git LFS hosted 仓库（Swagger: GET /v1/repositories/gitlfs/hosted/{repositoryName}）
// Get Git LFS hosted repository (Swagger: GET /v1/repositories/gitlfs/hosted/{repositoryName})
func (c *Client) GetGitlfsHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/gitlfs/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateGitlfsHostedRepository 更新 Git LFS hosted 仓库（Swagger: PUT /v1/repositories/gitlfs/hosted/{repositoryName}）
// Update Git LFS hosted repository (Swagger: PUT /v1/repositories/gitlfs/hosted/{repositoryName})
func (c *Client) UpdateGitlfsHostedRepository(ctx context.Context, repositoryName string, req GitLfsHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/gitlfs/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateGoGroupRepository 创建 Go group 仓库（Swagger: POST /v1/repositories/go/group）
// Create Go group repository (Swagger: POST /v1/repositories/go/group)
func (c *Client) CreateGoGroupRepository(ctx context.Context, req GolangGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/go/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateGoProxyRepository 创建 Go proxy 仓库（Swagger: POST /v1/repositories/go/proxy）
// Create Go proxy repository (Swagger: POST /v1/repositories/go/proxy)
func (c *Client) CreateGoProxyRepository(ctx context.Context, req GolangProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/go/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetGoGroupRepository 获取 Go group 仓库（Swagger: GET /v1/repositories/go/group/{repositoryName}）
// Get Go group repository (Swagger: GET /v1/repositories/go/group/{repositoryName})
func (c *Client) GetGoGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/go/group/" + repositoryName
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

// GetGoProxyRepository 获取 Go proxy 仓库（Swagger: GET /v1/repositories/go/proxy/{repositoryName}）
// Get Go proxy repository (Swagger: GET /v1/repositories/go/proxy/{repositoryName})
func (c *Client) GetGoProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/go/proxy/" + repositoryName
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

// UpdateGoGroupRepository 更新 Go group 仓库（Swagger: PUT /v1/repositories/go/group/{repositoryName}）
// Update Go group repository (Swagger: PUT /v1/repositories/go/group/{repositoryName})
func (c *Client) UpdateGoGroupRepository(ctx context.Context, repositoryName string, req GolangGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/go/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateGoProxyRepository 更新 Go proxy 仓库（Swagger: PUT /v1/repositories/go/proxy/{repositoryName}）
// Update Go proxy repository (Swagger: PUT /v1/repositories/go/proxy/{repositoryName})
func (c *Client) UpdateGoProxyRepository(ctx context.Context, repositoryName string, req GolangProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/go/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateHelmHostedRepository 创建 Helm hosted 仓库（Swagger: POST /v1/repositories/helm/hosted）
// Create Helm hosted repository (Swagger: POST /v1/repositories/helm/hosted)
func (c *Client) CreateHelmHostedRepository(ctx context.Context, req HelmHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/helm/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateHelmProxyRepository 创建 Helm proxy 仓库（Swagger: POST /v1/repositories/helm/proxy）
// Create Helm proxy repository (Swagger: POST /v1/repositories/helm/proxy)
func (c *Client) CreateHelmProxyRepository(ctx context.Context, req HelmProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/helm/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetHelmHostedRepository 获取 Helm hosted 仓库（Swagger: GET /v1/repositories/helm/hosted/{repositoryName}）
// Get Helm hosted repository (Swagger: GET /v1/repositories/helm/hosted/{repositoryName})
func (c *Client) GetHelmHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/helm/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetHelmProxyRepository 获取 Helm proxy 仓库（Swagger: GET /v1/repositories/helm/proxy/{repositoryName}）
// Get Helm proxy repository (Swagger: GET /v1/repositories/helm/proxy/{repositoryName})
func (c *Client) GetHelmProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/helm/proxy/" + repositoryName
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

// UpdateHelmHostedRepository 更新 Helm hosted 仓库（Swagger: PUT /v1/repositories/helm/hosted/{repositoryName}）
// Update Helm hosted repository (Swagger: PUT /v1/repositories/helm/hosted/{repositoryName})
func (c *Client) UpdateHelmHostedRepository(ctx context.Context, repositoryName string, req HelmHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/helm/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateHelmProxyRepository 更新 Helm proxy 仓库（Swagger: PUT /v1/repositories/helm/proxy/{repositoryName}）
// Update Helm proxy repository (Swagger: PUT /v1/repositories/helm/proxy/{repositoryName})
func (c *Client) UpdateHelmProxyRepository(ctx context.Context, repositoryName string, req HelmProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/helm/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateHuggingfaceProxyRepository 创建 HuggingFace proxy 仓库（Swagger: POST /v1/repositories/huggingface/proxy）
// Create HuggingFace proxy repository (Swagger: POST /v1/repositories/huggingface/proxy)
func (c *Client) CreateHuggingfaceProxyRepository(ctx context.Context, req HuggingFaceProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/huggingface/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetHuggingfaceProxyRepository 获取 HuggingFace proxy 仓库（Swagger: GET /v1/repositories/huggingface/proxy/{repositoryName}）
// Get HuggingFace proxy repository (Swagger: GET /v1/repositories/huggingface/proxy/{repositoryName})
func (c *Client) GetHuggingfaceProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/huggingface/proxy/" + repositoryName
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

// UpdateHuggingfaceProxyRepository 更新 HuggingFace proxy 仓库（Swagger: PUT /v1/repositories/huggingface/proxy/{repositoryName}）
// Update HuggingFace proxy repository (Swagger: PUT /v1/repositories/huggingface/proxy/{repositoryName})
func (c *Client) UpdateHuggingfaceProxyRepository(ctx context.Context, repositoryName string, req HuggingFaceProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/huggingface/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateMavenGroupRepository 创建 Maven group 仓库（Swagger: POST /v1/repositories/maven/group）
// Create Maven group repository (Swagger: POST /v1/repositories/maven/group)
func (c *Client) CreateMavenGroupRepository(ctx context.Context, req MavenGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateMavenHostedRepository 创建 Maven hosted 仓库（Swagger: POST /v1/repositories/maven/hosted）
// Create Maven hosted repository (Swagger: POST /v1/repositories/maven/hosted)
func (c *Client) CreateMavenHostedRepository(ctx context.Context, req MavenHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateMavenProxyRepository 创建 Maven proxy 仓库（Swagger: POST /v1/repositories/maven/proxy）
// Create Maven proxy repository (Swagger: POST /v1/repositories/maven/proxy)
func (c *Client) CreateMavenProxyRepository(ctx context.Context, req MavenProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetMavenGroupRepository 获取 Maven group 仓库（Swagger: GET /v1/repositories/maven/group/{repositoryName}）
// Get Maven group repository (Swagger: GET /v1/repositories/maven/group/{repositoryName})
func (c *Client) GetMavenGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/group/" + repositoryName
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

// GetMavenHostedRepository 获取 Maven hosted 仓库（Swagger: GET /v1/repositories/maven/hosted/{repositoryName}）
// Get Maven hosted repository (Swagger: GET /v1/repositories/maven/hosted/{repositoryName})
func (c *Client) GetMavenHostedRepository(ctx context.Context, repositoryName string) (*MavenHostedApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo MavenHostedApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetMavenProxyRepository 获取 Maven proxy 仓库（Swagger: GET /v1/repositories/maven/proxy/{repositoryName}）
// Get Maven proxy repository (Swagger: GET /v1/repositories/maven/proxy/{repositoryName})
func (c *Client) GetMavenProxyRepository(ctx context.Context, repositoryName string) (*MavenProxyApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo MavenProxyApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateMavenGroupRepository 更新 Maven group 仓库（Swagger: PUT /v1/repositories/maven/group/{repositoryName}）
// Update Maven group repository (Swagger: PUT /v1/repositories/maven/group/{repositoryName})
func (c *Client) UpdateMavenGroupRepository(ctx context.Context, repositoryName string, req MavenGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateMavenHostedRepository 更新 Maven hosted 仓库（Swagger: PUT /v1/repositories/maven/hosted/{repositoryName}）
// Update Maven hosted repository (Swagger: PUT /v1/repositories/maven/hosted/{repositoryName})
func (c *Client) UpdateMavenHostedRepository(ctx context.Context, repositoryName string, req MavenHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateMavenProxyRepository 更新 Maven proxy 仓库（Swagger: PUT /v1/repositories/maven/proxy/{repositoryName}）
// Update Maven proxy repository (Swagger: PUT /v1/repositories/maven/proxy/{repositoryName})
func (c *Client) UpdateMavenProxyRepository(ctx context.Context, repositoryName string, req MavenProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/maven/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateNpmGroupRepository 创建 NPM group 仓库（Swagger: POST /v1/repositories/npm/group）
// Create NPM group repository (Swagger: POST /v1/repositories/npm/group)
func (c *Client) CreateNpmGroupRepository(ctx context.Context, req NpmGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateNpmHostedRepository 创建 NPM hosted 仓库（Swagger: POST /v1/repositories/npm/hosted）
// Create NPM hosted repository (Swagger: POST /v1/repositories/npm/hosted)
func (c *Client) CreateNpmHostedRepository(ctx context.Context, req NpmHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateNpmProxyRepository 创建 NPM proxy 仓库（Swagger: POST /v1/repositories/npm/proxy）
// Create NPM proxy repository (Swagger: POST /v1/repositories/npm/proxy)
func (c *Client) CreateNpmProxyRepository(ctx context.Context, req NpmProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetNpmGroupRepository 获取 NPM group 仓库（Swagger: GET /v1/repositories/npm/group/{repositoryName}）
// Get NPM group repository (Swagger: GET /v1/repositories/npm/group/{repositoryName})
func (c *Client) GetNpmGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupDeployRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/group/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiGroupDeployRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetNpmHostedRepository 获取 NPM hosted 仓库（Swagger: GET /v1/repositories/npm/hosted/{repositoryName}）
// Get NPM hosted repository (Swagger: GET /v1/repositories/npm/hosted/{repositoryName})
func (c *Client) GetNpmHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetNpmProxyRepository 获取 NPM proxy 仓库（Swagger: GET /v1/repositories/npm/proxy/{repositoryName}）
// Get NPM proxy repository (Swagger: GET /v1/repositories/npm/proxy/{repositoryName})
func (c *Client) GetNpmProxyRepository(ctx context.Context, repositoryName string) (*NpmProxyApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo NpmProxyApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateNpmGroupRepository 更新 NPM group 仓库（Swagger: PUT /v1/repositories/npm/group/{repositoryName}）
// Update NPM group repository (Swagger: PUT /v1/repositories/npm/group/{repositoryName})
func (c *Client) UpdateNpmGroupRepository(ctx context.Context, repositoryName string, req NpmGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateNpmHostedRepository 更新 NPM hosted 仓库（Swagger: PUT /v1/repositories/npm/hosted/{repositoryName}）
// Update NPM hosted repository (Swagger: PUT /v1/repositories/npm/hosted/{repositoryName})
func (c *Client) UpdateNpmHostedRepository(ctx context.Context, repositoryName string, req NpmHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateNpmProxyRepository 更新 NPM proxy 仓库（Swagger: PUT /v1/repositories/npm/proxy/{repositoryName}）
// Update NPM proxy repository (Swagger: PUT /v1/repositories/npm/proxy/{repositoryName})
func (c *Client) UpdateNpmProxyRepository(ctx context.Context, repositoryName string, req NpmProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/npm/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateNugetGroupRepository 创建 NuGet group 仓库（Swagger: POST /v1/repositories/nuget/group）
// Create NuGet group repository (Swagger: POST /v1/repositories/nuget/group)
func (c *Client) CreateNugetGroupRepository(ctx context.Context, req NugetGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateNugetHostedRepository 创建 NuGet hosted 仓库（Swagger: POST /v1/repositories/nuget/hosted）
// Create NuGet hosted repository (Swagger: POST /v1/repositories/nuget/hosted)
func (c *Client) CreateNugetHostedRepository(ctx context.Context, req NugetHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateNugetProxyRepository 创建 NuGet proxy 仓库（Swagger: POST /v1/repositories/nuget/proxy）
// Create NuGet proxy repository (Swagger: POST /v1/repositories/nuget/proxy)
func (c *Client) CreateNugetProxyRepository(ctx context.Context, req NugetProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetNugetGroupRepository 获取 NuGet group 仓库（Swagger: GET /v1/repositories/nuget/group/{repositoryName}）
// Get NuGet group repository (Swagger: GET /v1/repositories/nuget/group/{repositoryName})
func (c *Client) GetNugetGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/group/" + repositoryName
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

// GetNugetHostedRepository 获取 NuGet hosted 仓库（Swagger: GET /v1/repositories/nuget/hosted/{repositoryName}）
// Get NuGet hosted repository (Swagger: GET /v1/repositories/nuget/hosted/{repositoryName})
func (c *Client) GetNugetHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetNugetProxyRepository 获取 NuGet proxy 仓库（Swagger: GET /v1/repositories/nuget/proxy/{repositoryName}）
// Get NuGet proxy repository (Swagger: GET /v1/repositories/nuget/proxy/{repositoryName})
func (c *Client) GetNugetProxyRepository(ctx context.Context, repositoryName string) (*NugetProxyApiRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/proxy/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo NugetProxyApiRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateNugetGroupRepository 更新 NuGet group 仓库（Swagger: PUT /v1/repositories/nuget/group/{repositoryName}）
// Update NuGet group repository (Swagger: PUT /v1/repositories/nuget/group/{repositoryName})
func (c *Client) UpdateNugetGroupRepository(ctx context.Context, repositoryName string, req NugetGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateNugetHostedRepository 更新 NuGet hosted 仓库（Swagger: PUT /v1/repositories/nuget/hosted/{repositoryName}）
// Update NuGet hosted repository (Swagger: PUT /v1/repositories/nuget/hosted/{repositoryName})
func (c *Client) UpdateNugetHostedRepository(ctx context.Context, repositoryName string, req NugetHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateNugetProxyRepository 更新 NuGet proxy 仓库（Swagger: PUT /v1/repositories/nuget/proxy/{repositoryName}）
// Update NuGet proxy repository (Swagger: PUT /v1/repositories/nuget/proxy/{repositoryName})
func (c *Client) UpdateNugetProxyRepository(ctx context.Context, repositoryName string, req NugetProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/nuget/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateP2ProxyRepository 创建 P2 proxy 仓库（Swagger: POST /v1/repositories/p2/proxy）
// Create P2 proxy repository (Swagger: POST /v1/repositories/p2/proxy)
func (c *Client) CreateP2ProxyRepository(ctx context.Context, req P2ProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/p2/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetP2ProxyRepository 获取 P2 proxy 仓库（Swagger: GET /v1/repositories/p2/proxy/{repositoryName}）
// Get P2 proxy repository (Swagger: GET /v1/repositories/p2/proxy/{repositoryName})
func (c *Client) GetP2ProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/p2/proxy/" + repositoryName
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

// UpdateP2ProxyRepository 更新 P2 proxy 仓库（Swagger: PUT /v1/repositories/p2/proxy/{repositoryName}）
// Update P2 proxy repository (Swagger: PUT /v1/repositories/p2/proxy/{repositoryName})
func (c *Client) UpdateP2ProxyRepository(ctx context.Context, repositoryName string, req P2ProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/p2/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreatePypiHostedRepository 创建 PyPI hosted 仓库（Swagger: POST /v1/repositories/pypi/hosted）
// Create PyPI hosted repository (Swagger: POST /v1/repositories/pypi/hosted)
func (c *Client) CreatePypiHostedRepository(ctx context.Context, req PypiHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreatePypiProxyRepository 创建 PyPI proxy 仓库（Swagger: POST /v1/repositories/pypi/proxy）
// Create PyPI proxy repository (Swagger: POST /v1/repositories/pypi/proxy)
func (c *Client) CreatePypiProxyRepository(ctx context.Context, req PypiProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreatePypiGroupRepository 创建 PyPI group 仓库（Swagger: POST /v1/repositories/pypi/group）
// Create PyPI group repository (Swagger: POST /v1/repositories/pypi/group)
func (c *Client) CreatePypiGroupRepository(ctx context.Context, req PypiGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetPypiHostedRepository 获取 PyPI hosted 仓库（Swagger: GET /v1/repositories/pypi/hosted/{repositoryName}）
// Get PyPI hosted repository (Swagger: GET /v1/repositories/pypi/hosted/{repositoryName})
func (c *Client) GetPypiHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetPypiProxyRepository 获取 PyPI proxy 仓库（Swagger: GET /v1/repositories/pypi/proxy/{repositoryName}）
// Get PyPI proxy repository (Swagger: GET /v1/repositories/pypi/proxy/{repositoryName})
func (c *Client) GetPypiProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/proxy/" + repositoryName
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

// GetPypiGroupRepository 获取 PyPI group 仓库（Swagger: GET /v1/repositories/pypi/group/{repositoryName}）
// Get PyPI group repository (Swagger: GET /v1/repositories/pypi/group/{repositoryName})
func (c *Client) GetPypiGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupDeployRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/group/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiGroupDeployRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdatePypiHostedRepository 更新 PyPI hosted 仓库（Swagger: PUT /v1/repositories/pypi/hosted/{repositoryName}）
// Update PyPI hosted repository (Swagger: PUT /v1/repositories/pypi/hosted/{repositoryName})
func (c *Client) UpdatePypiHostedRepository(ctx context.Context, repositoryName string, req PypiHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdatePypiProxyRepository 更新 PyPI proxy 仓库（Swagger: PUT /v1/repositories/pypi/proxy/{repositoryName}）
// Update PyPI proxy repository (Swagger: PUT /v1/repositories/pypi/proxy/{repositoryName})
func (c *Client) UpdatePypiProxyRepository(ctx context.Context, repositoryName string, req PypiProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdatePypiGroupRepository 更新 PyPI group 仓库（Swagger: PUT /v1/repositories/pypi/group/{repositoryName}）
// Update PyPI group repository (Swagger: PUT /v1/repositories/pypi/group/{repositoryName})
func (c *Client) UpdatePypiGroupRepository(ctx context.Context, repositoryName string, req PypiGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/pypi/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateRGroupRepository 创建 R group 仓库（Swagger: POST /v1/repositories/r/group）
// Create R group repository (Swagger: POST /v1/repositories/r/group)
func (c *Client) CreateRGroupRepository(ctx context.Context, req RGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/r/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateRHostedRepository 创建 R hosted 仓库（Swagger: POST /v1/repositories/r/hosted）
// Create R hosted repository (Swagger: POST /v1/repositories/r/hosted)
func (c *Client) CreateRHostedRepository(ctx context.Context, req RHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/r/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateRProxyRepository 创建 R proxy 仓库（Swagger: POST /v1/repositories/r/proxy）
// Create R proxy repository (Swagger: POST /v1/repositories/r/proxy)
func (c *Client) CreateRProxyRepository(ctx context.Context, req RProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/r/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetRGroupRepository 获取 R group 仓库（Swagger: GET /v1/repositories/r/group/{repositoryName}）
// Get R group repository (Swagger: GET /v1/repositories/r/group/{repositoryName})
func (c *Client) GetRGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/r/group/" + repositoryName
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

// GetRHostedRepository 获取 R hosted 仓库（Swagger: GET /v1/repositories/r/hosted/{repositoryName}）
// Get R hosted repository (Swagger: GET /v1/repositories/r/hosted/{repositoryName})
func (c *Client) GetRHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/r/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetRProxyRepository 获取 R proxy 仓库（Swagger: GET /v1/repositories/r/proxy/{repositoryName}）
// Get R proxy repository (Swagger: GET /v1/repositories/r/proxy/{repositoryName})
func (c *Client) GetRProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/r/proxy/" + repositoryName
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

// UpdateRGroupRepository 更新 R group 仓库（Swagger: PUT /v1/repositories/r/group/{repositoryName}）
// Update R group repository (Swagger: PUT /v1/repositories/r/group/{repositoryName})
func (c *Client) UpdateRGroupRepository(ctx context.Context, repositoryName string, req RGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/r/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateRHostedRepository 更新 R hosted 仓库（Swagger: PUT /v1/repositories/r/hosted/{repositoryName}）
// Update R hosted repository (Swagger: PUT /v1/repositories/r/hosted/{repositoryName})
func (c *Client) UpdateRHostedRepository(ctx context.Context, repositoryName string, req RHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/r/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateRProxyRepository 更新 R proxy 仓库（Swagger: PUT /v1/repositories/r/proxy/{repositoryName}）
// Update R proxy repository (Swagger: PUT /v1/repositories/r/proxy/{repositoryName})
func (c *Client) UpdateRProxyRepository(ctx context.Context, repositoryName string, req RProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/r/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateRawHostedRepository 创建 Raw hosted 仓库（Swagger: POST /v1/repositories/raw/hosted）
// Create raw hosted repository (Swagger: POST /v1/repositories/raw/hosted)
func (c *Client) CreateRawHostedRepository(ctx context.Context, req RawHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/raw/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetRawHostedRepository 获取 Raw hosted 仓库（Swagger: GET /v1/repositories/raw/hosted/{repositoryName}）
// Get raw hosted repository (Swagger: GET /v1/repositories/raw/hosted/{repositoryName})
func (c *Client) GetRawHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/raw/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// UpdateRawHostedRepository 更新 Raw hosted 仓库（Swagger: PUT /v1/repositories/raw/hosted/{repositoryName}）
// Update raw hosted repository (Swagger: PUT /v1/repositories/raw/hosted/{repositoryName})
func (c *Client) UpdateRawHostedRepository(ctx context.Context, repositoryName string, req RawHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/raw/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// CreateRubygemsGroupRepository 创建 RubyGems group 仓库（Swagger: POST /v1/repositories/rubygems/group）
// Create RubyGems group repository (Swagger: POST /v1/repositories/rubygems/group)
func (c *Client) CreateRubygemsGroupRepository(ctx context.Context, req RubyGemsGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/group"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateRubygemsHostedRepository 创建 RubyGems hosted 仓库（Swagger: POST /v1/repositories/rubygems/hosted）
// Create RubyGems hosted repository (Swagger: POST /v1/repositories/rubygems/hosted)
func (c *Client) CreateRubygemsHostedRepository(ctx context.Context, req RubyGemsHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/hosted"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// CreateRubygemsProxyRepository 创建 RubyGems proxy 仓库（Swagger: POST /v1/repositories/rubygems/proxy）
// Create RubyGems proxy repository (Swagger: POST /v1/repositories/rubygems/proxy)
func (c *Client) CreateRubygemsProxyRepository(ctx context.Context, req RubyGemsProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/proxy"
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	return err
}

// GetRubygemsGroupRepository 获取 RubyGems group 仓库（Swagger: GET /v1/repositories/rubygems/group/{repositoryName}）
// Get RubyGems group repository (Swagger: GET /v1/repositories/rubygems/group/{repositoryName})
func (c *Client) GetRubygemsGroupRepository(ctx context.Context, repositoryName string) (*SimpleApiGroupRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/group/" + repositoryName
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

// GetRubygemsHostedRepository 获取 RubyGems hosted 仓库（Swagger: GET /v1/repositories/rubygems/hosted/{repositoryName}）
// Get RubyGems hosted repository (Swagger: GET /v1/repositories/rubygems/hosted/{repositoryName})
func (c *Client) GetRubygemsHostedRepository(ctx context.Context, repositoryName string) (*SimpleApiHostedRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/hosted/" + repositoryName
	b, _, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	var repo SimpleApiHostedRepository
	if err := json.Unmarshal(b, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// GetRubygemsProxyRepository 获取 RubyGems proxy 仓库（Swagger: GET /v1/repositories/rubygems/proxy/{repositoryName}）
// Get RubyGems proxy repository (Swagger: GET /v1/repositories/rubygems/proxy/{repositoryName})
func (c *Client) GetRubygemsProxyRepository(ctx context.Context, repositoryName string) (*SimpleApiProxyRepository, error) {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/proxy/" + repositoryName
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

// UpdateRubygemsGroupRepository 更新 RubyGems group 仓库（Swagger: PUT /v1/repositories/rubygems/group/{repositoryName}）
// Update RubyGems group repository (Swagger: PUT /v1/repositories/rubygems/group/{repositoryName})
func (c *Client) UpdateRubygemsGroupRepository(ctx context.Context, repositoryName string, req RubyGemsGroupRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/group/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateRubygemsHostedRepository 更新 RubyGems hosted 仓库（Swagger: PUT /v1/repositories/rubygems/hosted/{repositoryName}）
// Update RubyGems hosted repository (Swagger: PUT /v1/repositories/rubygems/hosted/{repositoryName})
func (c *Client) UpdateRubygemsHostedRepository(ctx context.Context, repositoryName string, req RubyGemsHostedRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/hosted/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}

// UpdateRubygemsProxyRepository 更新 RubyGems proxy 仓库（Swagger: PUT /v1/repositories/rubygems/proxy/{repositoryName}）
// Update RubyGems proxy repository (Swagger: PUT /v1/repositories/rubygems/proxy/{repositoryName})
func (c *Client) UpdateRubygemsProxyRepository(ctx context.Context, repositoryName string, req RubyGemsProxyRepositoryApiRequest) error {
	url := c.BaseURL + "/service/rest/v1/repositories/rubygems/proxy/" + repositoryName
	body, _ := json.Marshal(req)
	_, _, err := c.do(ctx, http.MethodPut, url, strings.NewReader(string(body)))
	return err
}
