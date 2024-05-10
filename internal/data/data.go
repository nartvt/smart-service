package data

import (
	"github.com/nartvt/smart-service/ent"
	"github.com/nartvt/smart-service/internal/conf"
	"github.com/nartvt/smart-service/internal/infra"

	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	client *ent.Client
}

// NewData .
func NewData(conf *conf.Data) (*Data, func(), error) {
	infra.InitDataClient(conf)
	d := &Data{
		client: infra.Postgres,
	}
	return d, infra.Close, nil
}
