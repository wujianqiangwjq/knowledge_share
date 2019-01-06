package elastic

import (
	"log"
	"os"

	"context"
	//	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
)

const (
	host   = "http://localhost:9200"
	user   = "elastic"
	passwd = "elas123456tic"
)

type Item struct {
	//	Id           string   `json:"id"`
	Title        string `json:"title"`
	Os           string `json:"os"`
	Lico_Version string `json:"lico_version"`
	Description  string `json:"description"`
	Solution     string `json:"solution"`
	Feedback     string `json:"feedback"`
}

var client *elastic.Client
var err error

func init() {
	errorlog := log.New(os.Stdout, "elastic", log.LstdFlags)
	client, err = elastic.NewClient(elastic.SetURL(host),
		elastic.SetBasicAuth(user, passwd),
		elastic.SetErrorLog(errorlog))
	if err != nil {
		panic(err)
	}

}

func SearchData(data string) []interface{} {
	rdata := []interface{}{}
	if data != "" {
		//		sql := fmt.Sprintf("title:%s+description:%s+lico_version:%s+os:%s", data, data, data, data)
		//		fmt.Println(sql)
		q := elastic.NewQueryStringQuery(data)
		q.Field("lico_version")
		q.Field("os")
		q.Field("title")
		q.Field("description")

		res, qer := client.Search().Index("issues").Type("doc").Query(q).Pretty(true).Do(context.Background())
		if qer != nil {
			fmt.Println(qer)
		}
		if res.Hits.TotalHits > 0 {
			for _, hit := range res.Hits.Hits {
				rdata = append(rdata, string(*hit.Source))
			}
		}
	} else {
		res, _ := client.Search().Index("issues").Type("doc").Do(context.Background())
		if res.Hits.TotalHits > 0 {
			for _, hit := range res.Hits.Hits {
				rdata = append(rdata, string(*hit.Source))
			}
		}
	}
	return rdata

}

func (item *Item) AddData() bool {
	fmt.Println(item.Solution)
	_, er := client.Index().Index("issues").Type("doc").BodyJson(item).Do(context.Background())
	if er != nil {
		fmt.Println(er)
		return false
	}
	return true

}
