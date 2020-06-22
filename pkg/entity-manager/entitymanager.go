package entitymanager

import (
	"errors"
	"log"
	"os"

	environmentloader "github.com/MarianoArias/ApiGo/pkg/environment-loader"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var client *gorm.DB

func init() {
	environmentloader.Load()

	connectionString := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("\033[97;41m%s\033[0m\n", "### MySQL connection error: "+err.Error()+" ###")
	} else {
		client = conn
		log.Printf("\033[97;42m%s\033[0m\n", "### MySQL connection established ###")
	}
}

func Ping() error {
	err := client.DB().Ping()

	if err != nil {
		return errors.New("Could not ping")
	}

	return nil
}

func GetClient() *gorm.DB {
	return client
}
