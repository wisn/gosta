package gosta

import (
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v7"
)

// New returns a Gosta Client struct
func New(cfg Config) (*Client, error) {
	esCfg := elasticsearch.Config{
		Addresses: []string{
			cfg.Host + ":" + strconv.Itoa(cfg.Port),
		},
	}

	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return nil, fmt.Errorf("Error creating the ElasticSearch client: %s", err)
	}

	ret := Client{
		Es: es,
	}

	return &ret, nil
}
