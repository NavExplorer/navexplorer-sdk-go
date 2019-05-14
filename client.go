package navexplorer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ExplorerClient struct {
	host       string
	network    string
	httpClient *http.Client
}

type Paginator struct {
	CurrentPage int   `json:"currentPage"`
	First       bool  `json:"first"`
	Last        bool  `json:"last"`
	Total       int64 `json:"total"`
	Size        int   `json:"size"`
	Pages       int   `json:"total_pages"`
	Elements    int   `json:"number_of_elements"`
}

func NewClient(host string, network string) (c *ExplorerClient, err error) {
	if len(host) == 0 {
		err = errors.New("bad call missing argument host")
		return
	}

	c = &ExplorerClient{host: host, network: network, httpClient: &http.Client{}}
	return
}

func (c *ExplorerClient) call(method string) (response []byte, paginator Paginator, err error) {
	req, err := http.NewRequest("GET", c.host+method, nil)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Network", string(c.network))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("HTTP error: " + resp.Status)
		return
	}

	paginationHeader := resp.Header.Get("X-Pagination")
	if paginationHeader != "" {
		err = json.Unmarshal([]byte(paginationHeader), &paginator)
		if err != nil {
			return
		}
	}

	return data, paginator, nil
}
