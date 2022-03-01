package database

import (
	"fmt"

	"github.com/martinsmessias/go_rest_api/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func loadConfig(path string) (config models.Config, err error) {
	fmt.Println("Loading config...")
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func ConnectDatabase() {
	config, err := loadConfig(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	connectionString := "host=" + config.DBHost + " port=" + config.DBPort + " user=" + config.DBUser + " dbname=" + config.DBName + " password=" + config.DBPassword + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic(err)
	}
}
