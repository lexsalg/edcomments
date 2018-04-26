package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Configuration clase de configuración
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration obtiene la configuración de BD.
func getConfiguration() Configuration {

	var c Configuration

	file, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c

}

// GetConnection devuelve la conexión a la BD
func GetConnection() *gorm.DB {
	c := getConfiguration()
	//data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
