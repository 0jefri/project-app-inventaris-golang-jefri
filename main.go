package main

import (
	"fmt"
	config "inventaris/config"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// if err := viper.Unmarshal(&Cfg); err != nil {
	// 	panic(err)
	// }

	db, err := config.InitDB()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	routes := config.InitRoute(db)
	config.RunServer(routes)
}
