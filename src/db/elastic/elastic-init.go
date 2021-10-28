package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticInstance struct {
	Client *elasticsearch.Client
}

var Instance ElasticInstance

const elasticURI = "http://localhost:9200"

func ConfigureAndConnect() error {
	config := elasticsearch.Config{
		Addresses: []string{
			elasticURI,
		},
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		return err
	}

	Instance = ElasticInstance{
		Client: es,
	}

	return nil
}
