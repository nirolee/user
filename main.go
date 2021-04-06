package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nirolee/user.git/domain/repository"
	service2 "github.com/nirolee/user.git/domain/service"
	"github.com/nirolee/user.git/proto/user"

	//"github.com/nirolee/user.git/domain/repository"
	"github.com/nirolee/user.git/handler"

)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		)
	//初始化服务
	srv.Init()
	// Register handler

	//创建数据库链接
	db, err := gorm.Open("mysql", "root:123456@micro?charset=utf8mb4")
	if err!= nil{
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	//数据表初始化 只执行一次

	//rp := repository.NewUserRepository(db)
	//rp.InitTable()
	userDataService := service2.NewsUserDataService(repository.NewUserRepository(db))

	err = user.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
