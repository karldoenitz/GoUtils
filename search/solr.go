package search

import (
	"github.com/rtt/Go-Solr"
	"fmt"
	//"reflect"
)

const solrHost = "127.0.0.1"
const port = 8983
const core = "music"

type Music struct {
	id int64
	name string
	singer []int64
	searchWorld []string
	url string
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
		fmt.Println(id, name, url, singer, searchWord)
		music := Music{
			id: id,
			name: name,
			url: url,
			singer: singer,
			searchWorld: searchWord,
		}
		fmt.Println(music)
		//titleNorm := reflect.ValueOf(collection.Fields["title_norm"])
		//fmt.Println("title_norm:", titleNorm.Index(0))
	}
}
