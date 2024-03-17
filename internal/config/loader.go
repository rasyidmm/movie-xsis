package config

import (
	"encoding/json"
	"fmt"
	"github.com/rasyidmm/movie-xsis/internal/config/db"
	"github.com/rasyidmm/movie-xsis/internal/config/server"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Server   server.ServerList
	Database db.DatabaseList
}

var cfg config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(dir + "/internal/config/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error server config file: %w \n", err))
	}

	viper.AddConfigPath(dir + "/internal/config/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error db config file: %w \n", err))
	}
	viper.UnmarshalExact(&cfg)

	fmt.Println("=============================")
	dataByte, _ := json.Marshal(cfg)
	fmt.Println(string(dataByte))
	fmt.Println("=============================")
}

func GetConfig() *config {
	return &cfg
}
