package nexus

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type BrowseRepository struct {
	Href string `json:"href,omitempty"`
	Url  string `json:"url,omitempty"`
	Text string `json:"text,omitempty"`
	Path string `json:"path,omitempty"`
	Type string `json:"type,omitempty"`
}

// GetBrowseRepository 浏览仓库
func (s *ExtDirectService) GetBrowseRepository(repository string, path string, options ...RequestOptionFunc) ([]BrowseRepository, *Response, error) {

	u := fmt.Sprintf("service/rest/repository/browse/%s/%s", repository, path)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, nil, options)
	if err != nil {
		return nil, nil, err
	}

	// 注意：此处响应的是 html 页面，第二个参数不能为 nil，否则将无结果
	resp, err := s.client.Do(req, "")
	if err != nil {
		return nil, resp, err
	}

	reader := strings.NewReader(resp.body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	var browseRepositories []BrowseRepository

	var repositoryUrl = s.client.baseURL.String() + u

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		href, _ := s.Attr("href")

		if href != "../" {
			browseRepository := BrowseRepository{
				Href: href,
				Text: text,
			}

			if browseRepository.Href[len(browseRepository.Href)-1] == '/' {
				browseRepository.Url = repositoryUrl + browseRepository.Href
				browseRepository.Type = "folder"
				browseRepository.Path = path + href
			} else {
				browseRepository.Url = browseRepository.Href
				browseRepository.Type = "file"
				browseRepository.Path = path + text
			}

			browseRepositories = append(browseRepositories, browseRepository)
		}
	})

	return browseRepositories, resp, nil
}
