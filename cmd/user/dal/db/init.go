package db

import (
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormtracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormtracing.New()); err != nil {
		panic(err)
	}

	//err = DB.AutoMigrate(&RelationRaw{})
	//if err != nil {
	//	panic(err)
	//}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(constants.MySQLMaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MySQLMaxOpenConns)
	sqlDB.SetConnMaxLifetime(constants.MySQLConnMaxLifetime)

}
