package search

import (
	"github.com/rtt/Go-Solr"
	"fmt"
	"encoding/json"
	//"reflect"
)

const solrHost = "127.0.0.1"
const port = 8983
const core = "music"

type Data struct {
	Data []Music
	Length int
}

type Music struct {
	Id int64
	Name string
	Singer []int64
	SearchWorld []string
	Url string
}

func Test_solr()  {
	solrClient, err := solr.Init(solrHost, port, core)
	if err != nil {
		panic(err)
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{"id:1"},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "name+ASC",
	}
	queryString := q.String()
	fmt.Println(queryString)
	res, err := solrClient.Select(&q)
	if err != nil {
		panic(err)
	}
	results := res.Results
	result := []Music{}
	for i := 0; i < results.NumFound; i++ {
		collection := results.Collection[i]
		id := int64(collection.Fields["id"].(float64))
		name := collection.Fields["name"].(string)
		url := collection.Fields["url"].(string)
		singerInterface := collection.Fields["singer"].([]interface{})
		searchWordInterface := collection.Fields["searchWord"].([]interface{})
		singer := make([]int64, len(singerInterface))
		for i := range singerInterface {
			singer[i] = int64(singerInterface[i].(float64))
		}
		searchWord := make([]string, len(searchWordInterface))
		for i := range searchWordInterface {
			searchWord[i] = searchWordInterface[i].(string)
		}
		musicObj := Music{
			Id: id,
			Name: name,
			Url: url,
			Singer: singer,
			SearchWorld: searchWord,
		}
		result = append(result, musicObj)
	}
	data := Data{
		Data:result,
		Length:len(result),
	}
	jsonData, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		panic(jsonErr)
	}
	fmt.Println(string(jsonData))
}
