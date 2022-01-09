package component

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/olivere/elastic/v7"
)

var esClient *elastic.Client
var esContext context.Context

func InitES() {
	g.Log().Info("init elastic search")
	host := g.Cfg().GetString("elasticsearch.host")
	port := g.Cfg().GetString("elasticsearch.port")
	username := g.Cfg().GetString("elasticsearch.user")
	password := g.Cfg().GetString("elasticsearch.pass")
	url := host + ":" + port
	esContext = context.Background()
	var err error
	esClient, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetHealthcheck(false), elastic.SetBasicAuth(username, password))
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := esClient.Ping(url).Do(esContext)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Infof("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esVersion, err := esClient.ElasticsearchVersion(url)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Infof("Elasticsearch version %s", esVersion)
}

func GetESClient() *elastic.Client {
	return esClient
}

func GetESContext() context.Context {
	return esContext
}
