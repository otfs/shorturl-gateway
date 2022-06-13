package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/speps/go-hashids/v2"
	"github.com/spf13/viper"
)

var db *sqlx.DB
var hashIds *hashids.HashID

func init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 短链
	salt := viper.GetString("shorturl.salt")
	hashIds, err = hashids.NewWithData(&hashids.HashIDData{
		Alphabet: hashids.DefaultAlphabet,
		Salt:     salt,
	})
	if err != nil {
		panic(err)
	}

	// 数据库
	dbHost := viper.GetString(`db.host`)
	dbPort := viper.GetString(`db.port`)
	dbUser := viper.GetString(`db.user`)
	dbPass := viper.GetString(`db.password`)
	dbName := viper.GetString(`db.database`)
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db = sqlx.MustOpen("mysql", dbConnection)
}
