package main

import (
	"fmt"
	jww "github.com/spf13/jwalterweatherman"
	"google.golang.org/grpc"
	"net"
	ur "tiny-url/internal/app/repository/psql"
	uu "tiny-url/internal/app/usecase"
	ud "tiny-url/internal/app/delivery/grpc"
	"gorm.io/gorm"
	"os"
	"time"
	config "tiny-url/internal/app"
	"gorm.io/driver/postgres"
	"tiny-url/internal/models"
	desc "tiny-url/pkg/tiny-url-api"
)

//Необходимо реализовать сервис, который должен предоставлять API по созданию сокращённых ссылок следующего формата:
//- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
//- Ссылка должна быть длинной 10 символов
//- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)

//Сервис должен быть написан на Go и принимать следующие запросы по gRPC:
//1. Метод Create, который будет сохранять оригинальный URL в базе и возвращать сокращённый
//2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL

//Решение должно быть предоставлено в «конечном виде», а именно: Сервис должен быть распространён в виде
//Docker-образа В качестве хранилища можно использовать in-memory решение или postgresql.
//API должно быть описано в proto файле
//Покрыть реализованный функционал Unit-тестами
//

func getPostgres() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		jww.FATAL.Println(err)
	}
	db.AutoMigrate(&models.Url{})
	return db
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		jww.ERROR.Println("Config error")
		return
	}

	timeoutContext := 2 * time.Second
	db := getPostgres()

	urlRepo := ur.NewUserRepository(db)
	urlUsecase := uu.NewUrlUsecase(urlRepo, timeoutContext)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.ServiceGrpcPort))
	if err != nil {
		jww.ERROR.Println("Can't listen port")
	}

	serverGrpc := grpc.NewServer()

	desc.RegisterUrlServiceServer(serverGrpc, ud.NewServer(urlUsecase))

	jww.INFO.Printf("Service started at port :%d", cfg.ServiceGrpcPort)

	err = serverGrpc.Serve(lis)
	if err != nil {
		jww.FATAL.Println(err)
	}
}