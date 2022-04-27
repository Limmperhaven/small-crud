package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.digital-spirit.ru/study/artem_crud/internal"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/models"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/pkg/handler"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/pkg/repository"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/pkg/repository/postgres"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	repos := bindRepository()
	if repos == nil {
		logrus.Fatalf(`error binding repository type: "%s" repository type does not exist`,
			viper.GetString("repository.type"))
	}
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)

	go func() {
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server^: %s", err.Error())
		}
	}()

	logrus.Println("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func bindRepository() *repository.Repository {
	switch viper.GetString("repository.type") {
	case "in-memory":
		return repository.NewInMemoryRepository(make(map[string]models.Record))
	case "postgres":
		db, err := postgres.NewPostgresDB(postgres.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		})
		if err != nil {
			logrus.Fatalf("failed to initialize db: %s", err.Error())
		}

		return repository.NewPostgresRepository(db)
	default:
		return nil
	}
}
