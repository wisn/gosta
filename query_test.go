package gosta

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	cfg := Config{
		Host: "http://shroudoftheavatar.com",
		Port: 9200,
	}

	g, err := New(cfg)
	if err != nil {
		t.Error(err)
	}

	t.Run("TestRawQuery", RawQueryFn(g))
	t.Run("TestRawQueryStr", RawQueryStrFn(g))
}

func RawQueryFn(g *Client) func(*testing.T) {
	return func(t *testing.T) {
		q := map[string]interface{}{
			"size": 2,
			"sort": []map[string]interface{}{
				map[string]interface{}{
					"@timestamp": map[string]interface{}{
						"order":         "desc",
						"unmapped_type": "boolean",
					},
				},
			},
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						map[string]interface{}{
							"query_string": map[string]interface{}{
								"analyze_wildcard": true,
								"query":            "LocationEvent:*",
							},
						},
					},
				},
			},
		}

		_, err := g.RawQuery(q)
		if err != nil {
			t.Error(err)
		}
	}
}

func RawQueryStrFn(g *Client) func(*testing.T) {
	return func(t *testing.T) {
		q := `
			{
				"size": 2,
				"sort": [
					{
						"@timestamp": {
							"order": "desc",
							"unmapped_type": "boolean"
						}
					}
				],
				"query": {
					"bool": {
						"must": [
							{
								"query_string": {
									"analyze_wildcard": true,
									"query": "LocationEvent:ItemGained_*"
								}
							}
						]
					}
				}
			}
		`

		res, err := g.RawQueryStr(q)
		if err != nil {
			t.Error(err)
		}

		log.Println(res)

		json, err := g.JSON(res)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(json)
	}
}
