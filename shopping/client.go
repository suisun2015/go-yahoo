package shopping

import (
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

const (
	baseURL = "https://shopping.yahooapis.jp/ShoppingWebService/V1"
	AppID   = "<your App ID>"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	AppID      string
	Logger     *log.Logger
}

func NewClient(logger *log.Logger) (*Client, error) {
	var client Client
	var err error

	parsedURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, errors.New("failed to parse url")
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client.URL = parsedURL
	client.HTTPClient = &http.Client{}
	client.AppID = AppID
	client.Logger = logger

	return &client, nil
}

func (c *Client) newRequest(method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	unescapedURL, err := url.QueryUnescape(u.String())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, unescapedURL, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
