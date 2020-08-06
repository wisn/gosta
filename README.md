# Gosta

A Go wrapper for querying open ElasticSearch (ES) service from
[Shroud of the Avatar: Forsaken Virtues][1] (SotA) events log.

Inspired by [rthompsonj/SotAPublicStatsQuery][2].

## Requirements

- go ^= v1.14
- go-elasticsearch ^= v7.8

## Installation

Add Gosta to your `go.mod` file by adding this line:

```
require github.com/wisn/gosta
```

Or simply run `go get github.com/wisn/gosta` command.

Don't forget to run `go mod tidy` and `go mod verify`.

## Examples

Currently, there are two methods available.
Another method will be coming shortly after I have conducted some research
with the events log data.
After all, we need a library with a good documentation so we can use it with
ease.

### RawQuery() Example

The `RawQuery` method use `map[string]interface{}` as the input and will
returns `(map[string]interface{}, error)` pair. We can freely query ES events
log as long as we understand how to do it. Here is an example for querying
all kind of `LocationEvent` types and sort the records by `@timeline` in the
descending manner.

```go
package main

import (
	"log"

	"github.com/wisn/gosta"
)

func main() {
	cfg := gosta.Config{
		Host: "http://shroudoftheavatar.com",
		Port: 9200,
	}

	gst, err := gosta.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	qry := map[string]interface{}{
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
							"query":            "LocationEvent:*", // here is the query
						},
					},
				},
			},
		},
	}

	res, err := gst.RawQuery(qry)
	if err != nil {
		log.Fatal(err)
	}

	// print as map[string]interface{} type
	log.Println(res)

	// print as a JSON in the string form
	json, err := gst.JSON(res)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(json)
}
```

### RawQueryStr() Example

The `RawQuery` method might be annoying for some people. Hence, there is
`RawQueryStr` method that use JSON format (in the form of `string`) as
the input and will returns `(map[string]interface{}, error)` pair. The
input argument used below is equivalent with the one that used by
`RawQuery` above.

```go
package main

import (
	"log"

	"github.com/wisn/gosta"
)

func main() {
	cfg := gosta.Config{
		Host: "http://shroudoftheavatar.com",
		Port: 9200,
	}

	gst, err := gosta.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	qry := `
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
								"query": "LocationEvent:*"
							}
						}
					]
				}
			}
		}
	`

	res, err := gst.RawQueryStr(qry)
	if err != nil {
		log.Fatal(err)
	}

	// print as map[string]interface{} type
	log.Println(res)

	// print as a JSON in the string form
	json, err := gst.JSON(res)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(json)
}
```

## Result Example

#### without gosta.JSON method

```
map[_shards:map[failed:0 skipped:0 successful:152 total:152] hits:map[hits:[map[_id:NUaNxHMB9DlaFXTY4AYT _index:logstash-2020.08.06 _score:<nil> _source:map[@timestamp:2020-08-06T16:15:00.000Z Archetype:Consumables/Potions/ItemPotion_Haste EconomyGoldDelta:10 GoldValue:10 ItemId:5f2c2c7e4d97a54074436337 LocationEvent:ItemGained_Loot Quantity:1 SceneName:The Rise timestamp:Aug 06 11:15:00 xpos:-1.318 ypos:10.095 zpos:73.121] _type:doc sort:[1.5967305e+12]] map[_id:VUaNxHMB9DlaFXTY5QYc _index:logstash-2020.08.06 _score:<nil> _source:map[@timestamp:2020-08-06T16:15:00.000Z Archetype:Consumables/Misc/ItemConsume_Recall EconomyGoldDelta:15 GoldValue:15 ItemId:5f2c2c854d97a54074436408 LocationEvent:ItemGained_Loot Quantity:1 SceneName:PlayerDungeon_GreyStone timestamp:Aug 06 11:15:00 xpos:75.351 ypos:0.008 zpos:-83.72] _type:doc sort:[1.5967305e+12]]] max_score:<nil> total:1.5974755e+07] timed_out:false took:327]
```

#### with gosta.JSON method

```
{"_shards":{"failed":0,"skipped":0,"successful":152,"total":152},"hits":{"hits":[{"_id":"NUaNxHMB9DlaFXTY4AYT","_index":"logstash-2020.08.06","_score":null,"_source":{"@timestamp":"2020-08-06T16:15:00.000Z","Archetype":"Consumables/Potions/ItemPotion_Haste","EconomyGoldDelta":"10","GoldValue":"10","ItemId":"5f2c2c7e4d97a54074436337","LocationEvent":"ItemGained_Loot","Quantity":"1","SceneName":"The Rise","timestamp":"Aug 06 11:15:00","xpos":"-1.318","ypos":"10.095","zpos":"73.121"},"_type":"doc","sort":[1596730500000]},{"_id":"VUaNxHMB9DlaFXTY5QYc","_index":"logstash-2020.08.06","_score":null,"_source":{"@timestamp":"2020-08-06T16:15:00.000Z","Archetype":"Consumables/Misc/ItemConsume_Recall","EconomyGoldDelta":"15","GoldValue":"15","ItemId":"5f2c2c854d97a54074436408","LocationEvent":"ItemGained_Loot","Quantity":"1","SceneName":"PlayerDungeon_GreyStone","timestamp":"Aug 06 11:15:00","xpos":"75.351","ypos":"0.008","zpos":"-83.72"},"_type":"doc","sort":[1596730500000]}],"max_score":null,"total":15976177},"timed_out":false,"took":325}
```

## Versioning

Gosta following the [Semantic Versioning 2.0.0][3].

## Roadmap

Here is the roadmap that I could thought so far.
Of course this roadmap is subject to change in the future.

#### v1.0.0

Initial release.

#### v1.1.0

First minor release for v1 will adding some features such as:
- `QueryGainedLoot`, targeting `LocationEvent:ItemGained_Loot` event log.
  There will be some arguments for building the query so we can fetch a specific
  records.
- `QueryPvEDeaths`, targeting `LocationEvent:PlayerKilledByMonster` event log.
- `QueryPvEKills`, targeting `LocationEvent:MonsterKilledByPlayer` event log.

#### v1.2.0

This minor release will be focusing on merchant related stuff.
At least two of them will be covered.
- `QueryItemSold`, targeting `LocationEvent:ItemGained_Merchant` event log.
- `QueryItemForSale`, targeting `LocationEvent:ItemDestroyed_Merchant` event log.

#### v1.3.0

This minor release will be focusing on crafting related stuff. Something like:
- `QueryCraftedItem`, targeting `LocationEvent:ItemGained_Crafting` event log.
As for `LocationEvent:ItemDestroyed_Crafting`, I need to conduct a research
before designing the method so this will be updated later on.

#### v1.4.0

This minor release will be focusing on PvP related stuff.
- `QueryPvPDeaths`, targeting `LocationEvent:PlayerKilledByPlayer` event log.
- `QueryPvPKills`, targeting `LocationEvent:PlayerKilledByPlayer` event log.

## Contributing

Since there are a lot of things to do, a pull request will be ignored at
the moment. However, an [issue][4] will be welcomed!

The pull request will be opened after the contributing guideline is ready.

## License

The code itself licensed under [The MIT License](LICENSE).

[1]: https://www.shroudoftheavatar.com/
[2]: https://github.com/rthompsonj/SotAPublicStatsQuery
[3]: https://semver.org/
[4]: https://github.com/wisn/gosta/issues/new/choose