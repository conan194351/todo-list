package http

import (
	"bytes"
	"encoding/json"
	"github.com/conan194351/todo-list.git/internal/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	client  *http.Client
	baseUrl string
	header  map[string]string
	cfg     config.Config
}

func NewHttpClient(
	client *http.Client,
	cfg config.Config,
	baseUrl string,
) *HttpClient {
	return &HttpClient{
		client:  client,
		cfg:     cfg,
		baseUrl: baseUrl,
	}
}

func (c *HttpClient) Get(path string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.Path(path, params), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *HttpClient) SetHeader(header map[string]string) {
	c.header = header
}

func (c *HttpClient) Post(path string, payload []byte, param map[string]string) (*http.Response, error) {
	var body io.Reader
	if payload != nil {
		body = bytes.NewBuffer(payload)
	}
	req, err := http.NewRequest(http.MethodPost, c.Path(path, param), body)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *HttpClient) Put(path string, payload []byte) (*http.Response, error) {
	url := c.Path(path, nil)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *HttpClient) Delete(path string) (*http.Response, error) {
	url := c.Path(path, nil)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	for k, v := range c.header {
		req.Header.Add(k, v)
	}
	return c.client.Do(req)
}

func (c *HttpClient) Path(uri string, params map[string]string) string {
	queryString := ""

	if params != nil {
		queryString = "?"
		for k, v := range params {
			encodedValue := url.QueryEscape(v)
			queryString += k + "=" + encodedValue + "&"
		}
		queryString = strings.TrimRight(queryString, "&")
	}

	base := strings.TrimRight(c.baseUrl, "/")
	if uri == "" {
		return base
	}
	return base + "/" + strings.TrimLeft(uri, "/") + queryString
}

func Decode(r *http.Response, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()
	if err := decoder.Decode(val); err != nil {
		return err
	}
	return nil
}

func Encode(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}
