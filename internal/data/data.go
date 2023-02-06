package data

import (
	"context"
	"douyin-simple/internal/conf"
	"douyin-simple/internal/data/ent"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"entgo.io/ent/dialect/sql"
)

type Data struct {
	db *ent.Client
}

func NewData(conf *conf.Data) *Data {
	var data = new(Data)
	drv, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		panic("init database driver err: " + err.Error())
	}
	drv.DB().SetConnMaxLifetime(time.Duration(conf.Database.ConnMaxLifetime))
	drv.DB().SetConnMaxIdleTime(time.Duration(conf.Database.MaxIdleConns))
	drv.DB().SetMaxIdleConns(conf.Database.MaxIdleConns)
	data.db = ent.NewClient(ent.Driver(drv))
	ctx := context.Background()
	if err := data.db.Schema.Create(ctx); err != nil {
		panic("database create err: " + err.Error())
	}
	return data
}
