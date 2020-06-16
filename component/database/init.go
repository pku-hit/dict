package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/micro/go-micro/util/log"
	"github.com/pku-hit/dict/model/entity"
)

var db *gorm.DB

var host string
var user string
var password string
var dbname string
var schema string

func registerEntity() {
	err := db.AutoMigrate(&entity.DictInfo{}).Error
	log.Infof("registering result %v.", err)
}

func init() {

	host = os.Getenv("opensvc.dict.dbHost")
	user = os.Getenv("opensvc.dict.dbUser")
	password = os.Getenv("opensvc.dict.dbPassword")
	dbname = os.Getenv("opensvc.dict.dbName")
	schema = os.Getenv("opensvc.dict.dbSchema")

	var err error
	url := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password)
	log.Infof("db url %s.", url)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		log.Errorf("connection error %v", err)
		panic(err)
	} else {
		log.Info(db)
	}

	registerEntity()
}
