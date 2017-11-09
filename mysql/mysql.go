package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/big"
)

var db *gorm.DB

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

func Initialize() {
	var err error
	db, err = gorm.Open("mysql", "root:111111@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err.Error())
	}
}

type Email struct {
	ID 		int			`gorm:"column:id;							primary_key"`
	UserID 	*big.Int	`gorm:"column:user_id;	type:bigint(20);	index"`
	Email	string 		`gorm:"column:email;	type:varchar(100)"`
}

func CreateEmailModel() {
	//Initialize()
	//db.CreateTable(&Email{})
	str := big.NewRat(1, 33)
	log.Println(str.Float32())
}