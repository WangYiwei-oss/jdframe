package configs

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/configinjector"
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"github.com/olivere/elastic/v7"
	"log"
	"sync"
)

func init() {
	configinjector.RegisterBean("ELASTICSEARCH", NewEsAdapter)
}

var doOnceEs sync.Once

type EsAdapter struct {
	*elastic.Client
}

var es *elastic.Client

func NewEsAdapter() *EsAdapter {
	var err error
	doOnceEs.Do(func() {
		config := configparser.GlobalSettings["ELASTICSEARCH_CONFIG"].(map[string]interface{})
		es, err = elastic.NewClient(
			elastic.SetURL(fmt.Sprintf("http://%s:%s", config["IP"].(string), config["PORT"].(string))),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Fatalln("连接ElasticSearch失败", err)
		}
		log.Println("连接ElasticSearch成功,URL: ", config["IP"].(string)+":"+config["PORT"].(string))
	})
	return &EsAdapter{
		Client: es,
	}
}
