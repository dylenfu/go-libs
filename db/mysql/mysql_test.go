package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/big"
	"testing"
	"time"
)

type Email struct {
	ID     int      `gorm:"column:id;							primary_key"`
	UserID *big.Int `gorm:"column:user_id;	type:bigint(20);	index"`
	Email  string   `gorm:"column:email;	type:varchar(100)"`
}

func TestCreateEmailModel(t *testing.T) {
	var (
		tablePrefix = "db"
		user        = "root"
		pwd         = "111111"
		host        = "127.0.0.1"
		port        = "6379"
		dbname      = "test"

		connMaxLifeTime    = time.Duration(100 * time.Second)
		maxIdleConnections = 10
		maxOpenConnections = 100
		debug              = true
	)

	// create database
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//url := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, pwd, host, port, dbname)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(fmt.Sprintf("mysql connection error:%s", err.Error()))
	}

	db.DB().SetConnMaxLifetime(connMaxLifeTime)
	db.DB().SetMaxIdleConns(maxIdleConnections)
	db.DB().SetMaxOpenConns(maxOpenConnections)

	db = db.LogMode(debug)

	// append tables
	tables := []interface{}{}
	tables = append(tables, &Email{})

	for _, t := range tables {
		if ok := db.HasTable(t); !ok {
			if err := db.CreateTable(t).Error; err != nil {
				panic(fmt.Sprintf("create mysql table error:%s", err.Error()))
			}
		}
	}

	// auto migrate to keep schema update to date
	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing column's type or delete unused columns to protect your data
	if err := db.AutoMigrate(tables...).Error; err != nil {
		panic(fmt.Sprintf("auto migrate table error:%s", err.Error()))
	}

	// TODO(fukun): insert and query
}
