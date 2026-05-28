package timer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/th3shadowbroker/hymetric/internal/cache"
)

type Client struct {
	httpCient *http.Client
	cache     *cache.MemoryCache[string, Response]
}

func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &Client{
		httpCient: httpClient,
		cache:     cache.NewMemoryCache[string, Response](10*time.Second, 20*time.Second),
	}
}

func (c *Client) GetTimer(def Definition) (*Response, error) {
	if cached, ok := c.cache.Get(def.Name); ok {
		return cached, nil
	}

	var response Response
	if req, err := http.NewRequest(http.MethodGet, def.Url, nil); err == nil {
		req.Header.Add("User-Agent", "HyMetric/v1")

		res, err := c.httpCient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code '%d': %s", res.StatusCode, def.Url)
		}

		if bodyBytes, err := io.ReadAll(res.Body); err == nil {
			if err := json.Unmarshal(bodyBytes, &response); err != nil {
				return nil, err
			}
		}

		c.cache.Set(def.Name, response)
		return &response, err
	}

	return nil, fmt.Errorf("could not retrieve data for timer '%s'", def.Name)
}
