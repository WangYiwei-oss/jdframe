package configs

import (
	"github.com/WangYiwei-oss/jdframe/src/configinjector"
	"github.com/olivere/elastic/v7"
	"sync"
)

func init() {
	configinjector.RegisterBean("ELASTICSEARCH", NewGormAdapter)
}

var doOnceEs sync.Once

type EsAdapter struct {
	*elastic.Client
}

var es *elastic.Client

func NewEsAdapter() *EsAdapter {
	var err error
	doOnceGorm.Do(func() {
		es, err = elastic.NewClient(
			elastic.SetURL("http://192.168.83.150:31000"),
			elastic.SetSniff(false),
		)
	})
	return &EsAdapter{
		Client: es,
	}
}
