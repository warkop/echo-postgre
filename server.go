package main

import (

	//"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

//func UsersModel() *pg.DB {
//	db := pg.Connect(&pg.Options{
//		Database: viper.GetString("DB_NAME"),
//		User: viper.GetString("DB_USER"),
//		Password: viper.GetString("DB_PASS"),
//	})
//
//	defer db.Close()
//	return db
//}

//Users is pattern from table users
type Users struct {
	gorm.Model
	id       int `gorm:"primary_key"`
	roleID   int
	name     string
	username string
	password string
}

func main() {
	//db := UsersModel()
	e := echo.New()
	var users Users

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(context echo.Context) error {
		db, err := gorm.Open("postgres", "host="+viper.GetString("DB_HOST")+" port="+viper.GetString("DB_PORT")+" user="+viper.GetString("DB_USER")+" dbname="+viper.GetString("DB_NAME")+" password="+viper.GetString("DB_PASS")+" sslmode=disable")
		defer db.Close()

		if err != nil {
			panic(err.Error())
		}

		return context.JSON(http.StatusOK, db.First(&users))
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("SERVER_PORT")))
}
