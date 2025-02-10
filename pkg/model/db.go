package model

import (
	"sync"

	"github.com/glebarez/sqlite"
	"github.com/alkaidchen/invitation-card/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var lock sync.Mutex

func NewMySQL(dsn string) (*gorm.DB, error) {
	return NewDB(mysql.Open(dsn))
}

func NewSQLite(path string) (*gorm.DB, error) {
	return NewDB(sqlite.Open(path))
}

func NewDB(dial gorm.Dialector) (*gorm.DB, error) {
	db, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Error("connect db failed, ", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&Item{},
	)
	if err != nil {
		log.Error("migrate database failed, ", err)
		return nil, err
	}
	return db, nil
}

func BeginTx() *gorm.DB {
	return db.Begin()
}
