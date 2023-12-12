package nexus

import "net/http"

type UsersService struct {
	client *Client
}

type ListUsersQuery struct {
	UserId string `json:"userId,omitempty" url:"userId,omitempty"`
	Source string `json:"source,omitempty" url:"source,omitempty"`
}

// ListUsers 检索用户列表。请注意，如果 source 不是 default，则响应限制为100个用户。
func (s *UsersService) ListUsers(requestQuery ListUsersQuery, options ...RequestOptionFunc) ([]*ApiUser, *Response, error) {

	u := "service/rest/v1/security/users"

	req, err := s.client.NewRequest(http.MethodGet, u, requestQuery, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var apiUsers []*ApiUser
	resp, err := s.client.Do(req, &apiUsers)
	if err != nil {
		return nil, resp, err
	}

	return apiUsers, resp, nil
}
