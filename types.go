package gosta

import (
	"github.com/elastic/go-elasticsearch/v7"
)

// Client type as an identifier for the Gosta instance
type Client struct {
	Es *elasticsearch.Client
}

// Config used by Gosta for a connection to ElasticSearch
type Config struct {
	Host string
	Port int
}
