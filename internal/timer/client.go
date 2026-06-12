package timer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/th3shadowbroker/hymetric/internal/cache"
	"github.com/th3shadowbroker/hymetric/internal/config"
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
		cache:     cache.NewMemoryCache[string, Response](config.Active.Cache.Ttl, config.Active.Cache.CleanupInterval),
	}
}

func (c *Client) GetTimer(def config.Timer) (*Response, error) {
	if cached, ok := c.cache.Get(def.Name); ok {
		return cached, nil
	}

	return c.fetchTimer(def)
}

func (c *Client) fetchTimer(def config.Timer) (*Response, error) {
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

func (c *Client) TryFetchAll(definitions []config.Timer) {
	for _, def := range definitions {
		if _, err := c.GetTimer(def); err != nil {
			log.Panicf("Failed to fetch timer %s: %s", err, def.Name)
		}
	}
}
