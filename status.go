package nexus

import "net/http"

type StatusService struct {
	client *Client
}

type GetStatusCheck struct {
	AvailableCPUs              Result `json:"Available CPUs"`
	BlobStoresQuota            Result `json:"Blob Stores Quota"`
	BlobStoresReady            Result `json:"Blob Stores Ready"`
	CoordinateContentSelectors Result `json:"Coordinate Content Selectors"`
	DefaultAdminCredentials    Result `json:"Default Admin Credentials"`
	DefaultRoleRealm           Result `json:"DefaultRoleRealm"`
	FileBlobStoresPath         Result `json:"File Blob Stores Path"`
	FileDescriptors            Result `json:"File Descriptors"`
	LifecyclePhase             Result `json:"Lifecycle Phase"`
	NuGetV2Repositories        Result `json:"NuGet V2 repositories"`
	ReadOnlyDetector           Result `json:"Read-Only Detector"`
	Scheduler                  Result `json:"Scheduler"`
	Scripting                  Result `json:"Scripting"`
	ThreadDeadlockDetector     Result `json:"Thread Deadlock Detector"`
	Transactions               Result `json:"Transactions"`
}

// GetStatusCheck 返回系统状态检查结果的运行状况检查终结点
func (s *StatusService) GetStatusCheck(options ...RequestOptionFunc) (*GetStatusCheck, *Response, error) {

	u := "service/rest/v1/status/check"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var getStatusCheck *GetStatusCheck
	resp, err := s.client.Do(req, &getStatusCheck)
	if err != nil {
		return nil, resp, err
	}

	return getStatusCheck, resp, nil
}

// GetStatus 验证服务器是否可以响应读取请求的运行状况检查终结点
func (s *StatusService) GetStatus(options ...RequestOptionFunc) (*Response, error) {

	u := "service/rest/v1/status"

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
