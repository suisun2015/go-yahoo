package auction

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"
)

const (
	baseURL = "http://auctions.yahooapis.jp/AuctionWebService/V2"
)

var (
	version   = "0.0.1"
	userAgent = fmt.Sprintf("YahooAuctionGoClient/%s (%s)", version, runtime.Version())
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	AuthToken string

	Logger *log.Logger
}

func NewClient(authToken string, logger *log.Logger) (*Client, error) {
	var client Client
	var err error

	if len(authToken) == 0 {
		return nil, errors.New("missing AuthToken")
	}

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
	client.AuthToken = authToken
	client.Logger = logger

	return &client, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
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
	req = req.WithContext(ctx)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
