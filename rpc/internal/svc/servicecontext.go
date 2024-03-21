package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"zero/rpc/dao/query"
	"zero/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Query  *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.Mysql.DSN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("mysql connection err: " + err.Error())
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Query:  query.Use(db),
	}
}
