package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func SimpleOrm() {
	db, err := gorm.Open("mysql", "root:111111@/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	hasEvent := db.HasTable("event")
	log.Println(hasEvent)
	log.Println("=======end")
}

