package main

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

//Users is pattern from table users
type Users struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    //column name is `email`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func connect() *gorm.DB {
	db, _ := gorm.Open("postgres", "host="+viper.GetString("DB_HOST")+" port="+viper.GetString("DB_PORT")+" user="+viper.GetString("DB_USER")+" dbname="+viper.GetString("DB_NAME")+" password="+viper.GetString("DB_PASS")+" sslmode=disable")

	return db
}

func closed(db *gorm.DB) {
	defer db.Close()
}

func main() {
	e := echo.New()
	var users []Users

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(context echo.Context) error {
		db := connect()
		db.Table("users").Find(&users)
		closed(db)
		return context.JSON(http.StatusOK, users)
	})

	e.GET("/:id", func(context echo.Context) error {
		id := context.Param("id")
		db := connect()
		db.Table("users").First(&users, id)
		closed(db)
		return context.JSON(http.StatusOK, users)
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("SERVER_PORT")))
}
