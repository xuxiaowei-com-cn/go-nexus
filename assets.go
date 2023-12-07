package nexus

import (
	"net/http"
	"time"
)

type AssetsService struct {
	client *Client
}

type ListAssetsRequest struct {
	Repository        string `json:"repository,omitempty" url:"repository,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty" url:"continuationToken,omitempty"`
}

type PageAssetXO struct {
	Items             []AssetXO `json:"items"`
	ContinuationToken string    `json:"continuationToken"`
}

type AssetXO struct {
	DownloadUrl    string      `json:"downloadUrl"`
	Path           string      `json:"path"`
	Id             string      `json:"id"`
	Repository     string      `json:"repository"`
	Format         string      `json:"format"`
	Checksum       interface{} `json:"checksum"`
	ContentType    string      `json:"contentType"`
	LastModified   time.Time   `json:"lastModified"`
	LastDownloaded time.Time   `json:"lastDownloaded"`
	Uploader       string      `json:"uploader"`
	UploaderIp     string      `json:"uploaderIp"`
	FileSize       int         `json:"fileSize"`
	BlobCreated    time.Time   `json:"blobCreated"`
}

// ListAssets 列出资产
func (s *AssetsService) ListAssets(requestQuery *ListAssetsRequest) (*PageAssetXO, *Response, error) {

	u := "service/rest/v1/assets"

	req, err := s.client.NewRequest(http.MethodGet, u, requestQuery, nil)
	if err != nil {
		return nil, nil, err
	}

	var pageAssetXO *PageAssetXO
	resp, err := s.client.Do(req, &pageAssetXO)
	if err != nil {
		return nil, resp, err
	}

	return pageAssetXO, resp, nil
}
