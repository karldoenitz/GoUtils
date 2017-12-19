package search

import (
	"github.com/rtt/Go-Solr"
	"fmt"
	"reflect"
)

const solrHost = "10.98.16.215"
const port = 9160
const core = "mitv/search/movie"

func Test_solr()  {
	solrClient, err := solr.Init(solrHost, port, core)
	if err != nil {
		panic(err)
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"*:*"},
			"fq": []string{"id:123221"},
			"wt": []string{"json"},
			"fl": []string{"*"},
		},
		Rows: 10,
		Sort: "title+ASC",
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
		fmt.Println("id:", collection.Fields["id"])
		fmt.Println("title:", collection.Fields["title"])
		titleNorm := reflect.ValueOf(collection.Fields["title_norm"])
		//index := []int{0}
		fmt.Println("title_norm:", titleNorm.Index(0))
		//titleNorm := collection.Fields["title_norm"]
		//fmt.Println("title_norm", titleNorm)
		fmt.Println("---------------------------------")
	}
}
