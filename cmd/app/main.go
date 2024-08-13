package main

import (
	"Avito/internal/repository"
	"Avito/internal/repository/postgres"
	"Avito/internal/service"
	"Avito/internal/transport"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("START")
	err := InitConfig()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	db, err := postgres.NewConnPostgres(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})
	if err != nil {
		logrus.Println(err)
		return
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	//serv := service.NewAuthService(repos.Authorization)
	//token, err := serv.GenerateToken("nurmag@")
	//fmt.Println("mykey:", []byte("myKey"))
	//if err != nil {
	//	logrus.Println(err)
	//}
	//email, err := serv.UserIdentity(token)
	//if err != nil {
	//	logrus.Println(err)
	//}
	//fmt.Println(email)
	handlers := transport.NewHandler(services)
	http.ListenAndServe("localhost:8080", handlers.InitRoutes())
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
