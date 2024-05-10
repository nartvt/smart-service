package infra

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nartvt/smart-service/ent"
	"github.com/nartvt/smart-service/internal/conf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	_ "github.com/lib/pq"
)

var Postgres *ent.Client

// NewData .
func InitDataClient(conf *conf.Data) {
	log := log.NewHelper(log.DefaultLogger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to db: %v", err)
		panic(err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		panic(err)
	}

	Postgres = client
}

func Close() {
	log.Info("message", "closing the data resources")
	if err := Postgres.Close(); err != nil {
		log.Error(err)
	}
}
