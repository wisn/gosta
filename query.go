package gosta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// RawQuery let us to freely querying SotA ElasticSearch service
func (c *Client) RawQuery(query map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}

	es := c.Es

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		}

		return nil,
			fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
	}

	var ret map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}

	return ret, nil
}

// RawQueryStr use JSON in the string form as the input argument
func (c *Client) RawQueryStr(str string) (map[string]interface{}, error) {
	var query map[string]interface{}
	if err := json.Unmarshal([]byte(str), &query); err != nil {
		return nil, fmt.Errorf("Error parsing input argument: %s", err)
	}

	return c.RawQuery(query)
}
