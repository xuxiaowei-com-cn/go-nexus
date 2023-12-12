package nexus

import (
	"math/rand"
	"net/http"
	"time"
)

type ExtDirectService struct {
	client *Client
}

type PostExtDirectReadBody struct {
	Action string                  `json:"action,omitempty"`
	Method string                  `json:"method,omitempty"`
	Data   []PostExtDirectReadData `json:"data,omitempty"`
	Type   string                  `json:"type,omitempty"`
	Tid    int                     `json:"tid,omitempty"`
}

type PostExtDirectReadData struct {
	RepositoryName string `json:"repositoryName,omitempty"`
	Node           string `json:"node,omitempty"`
}

type PostExtDirectReadResultData struct {
	Id          string `json:"id,omitempty"`
	Text        string `json:"text,omitempty"`
	Type        string `json:"type,omitempty"`
	Leaf        bool   `json:"leaf,omitempty"`
	ComponentId string `json:"componentId,omitempty"`
	AssetId     string `json:"assetId,omitempty"`
	PackageUrl  string `json:"packageUrl,omitempty"`
}

type PostExtDirectReadResult struct {
	Success bool                          `json:"success,omitempty"`
	Data    []PostExtDirectReadResultData `json:"data,omitempty"`
}

type PostExtDirectRead struct {
	Tid    int                     `json:"tid,omitempty"`
	Action string                  `json:"action,omitempty"`
	Method string                  `json:"method,omitempty"`
	Result PostExtDirectReadResult `json:"result,omitempty"`
	Type   string                  `json:"type,omitempty"`
}

// PostExtDirectRead 搜索资产
func (s *ExtDirectService) PostExtDirectRead(repository string, node string) (*PostExtDirectRead, *Response, error) {

	u := "service/extdirect"

	requestBody := &PostExtDirectReadBody{
		Action: "coreui_Browse",
		Method: "read",
		Type:   "rpc",
		Tid:    rand.Intn(10000),
		Data: []PostExtDirectReadData{
			{
				RepositoryName: repository,
				Node:           node,
			},
		},
	}

	req, err := s.client.NewRequest(http.MethodPost, u, nil, requestBody)
	if err != nil {
		return nil, nil, err
	}

	var postExtDirectRead *PostExtDirectRead
	resp, err := s.client.Do(req, &postExtDirectRead)
	if err != nil {
		return nil, resp, err
	}

	return postExtDirectRead, resp, nil
}

type PostExtDirectReadAssetBody struct {
	Action string   `json:"action"`
	Method string   `json:"method"`
	Data   []string `json:"data"`
	Type   string   `json:"type"`
	Tid    int      `json:"tid"`
}

type Checksum struct {
	Sha1   string `json:"sha1"`
	Sha512 string `json:"sha512"`
	Sha256 string `json:"sha256"`
	Md5    string `json:"md5"`
}

type Maven2 struct {
	Extension   string `json:"extension"`
	ArtifactId  string `json:"artifactId"`
	BaseVersion string `json:"baseVersion"`
	Version     string `json:"version"`
	AssetKind   string `json:"asset_kind"`
	GroupId     string `json:"groupId"`
}

type Cache struct {
	LastVerified time.Time `json:"last_verified"`
}

type Provenance struct {
	HashesNotVerified bool `json:"hashes_not_verified"`
}

type Content struct {
	LastModified time.Time `json:"last_modified"`
	Etag         string    `json:"etag"`
}

type Attributes struct {
	Checksum   Checksum   `json:"checksum"`
	Maven2     Maven2     `json:"maven2"`
	Cache      Cache      `json:"cache"`
	Provenance Provenance `json:"provenance"`
	Content    Content    `json:"content"`
}

type PostExtDirectReadAssetResultData struct {
	Id                       string     `json:"id"`
	Name                     string     `json:"name"`
	Format                   string     `json:"format"`
	ContentType              string     `json:"contentType"`
	Size                     int        `json:"size"`
	RepositoryName           string     `json:"repositoryName"`
	ContainingRepositoryName string     `json:"containingRepositoryName"`
	BlobCreated              time.Time  `json:"blobCreated"`
	BlobUpdated              time.Time  `json:"blobUpdated"`
	LastDownloaded           time.Time  `json:"lastDownloaded"`
	BlobRef                  string     `json:"blobRef"`
	ComponentId              string     `json:"componentId"`
	CreatedBy                string     `json:"createdBy"`
	CreatedByIp              string     `json:"createdByIp"`
	Attributes               Attributes `json:"attributes"`
}

type PostExtDirectReadAssetResult struct {
	Success bool                             `json:"success"`
	Data    PostExtDirectReadAssetResultData `json:"data"`
}

type PostExtDirectReadAsset struct {
	Tid    int                          `json:"tid"`
	Action string                       `json:"action"`
	Method string                       `json:"method"`
	Result PostExtDirectReadAssetResult `json:"result"`
	Type   string                       `json:"type"`
}

// PostExtDirectReadAsset 查看资产
func (s *ExtDirectService) PostExtDirectReadAsset(repository string, assetId string) (*PostExtDirectReadAsset, *Response, error) {

	u := "service/extdirect"

	requestBody := &PostExtDirectReadAssetBody{
		Action: "coreui_Component",
		Method: "readAsset",
		Type:   "rpc",
		Tid:    rand.Intn(10000),
		Data: []string{
			assetId,
			repository,
		},
	}

	req, err := s.client.NewRequest(http.MethodPost, u, nil, requestBody)
	if err != nil {
		return nil, nil, err
	}

	var postExtDirectReadAsset *PostExtDirectReadAsset
	resp, err := s.client.Do(req, &postExtDirectReadAsset)
	if err != nil {
		return nil, resp, err
	}

	return postExtDirectReadAsset, resp, nil
}
