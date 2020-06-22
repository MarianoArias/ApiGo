package elasticsearch

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	environmentloader "github.com/MarianoArias/ApiGo/pkg/environment-loader"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var url string

func init() {
	environmentloader.Load()

	url = os.Getenv("ELASTICSEARCH_SCHEME") + "://" + os.Getenv("ELASTICSEARCH_HOST") + ":" + os.Getenv("ELASTICSEARCH_PORT")

	elasticClient, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(os.Getenv("ELASTICSEARCH_USER"), os.Getenv("ELASTICSEARCH_PASSWORD")),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(true),
	)

	if elastic.IsConnErr(err) {
		log.Fatalf("\033[97;41m%s\033[0m\n", "### Elasticsearch connection error: "+err.Error()+" ###")
	} else {
		client = elasticClient
		log.Printf("\033[97;42m%s\033[0m\n", "### Elasticsearch connection established ###")
	}
}

func GetResults(index string, query elastic.Query, from int, size int) ([]*elastic.SearchHit, int64, error) {
	searchResult, err := client.Search().
		Index(index).
		Query(query).
		From(from).Size(size).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, 0, err
	}

	return searchResult.Hits.Hits, searchResult.TotalHits(), nil
}

func Ping() error {
	_, code, err := client.
		Ping(url).
		Do(context.Background())

	if code != http.StatusOK {
		return errors.New("No Elasticsearch node available")
	}

	if err != nil {
		return err
	}

	return nil
}

func GetClient() *elastic.Client {
	return client
}
