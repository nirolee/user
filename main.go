package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/nirolee/user.git/domain/repository"
	"github.com/nirolee/user.git/handler"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("latest"),
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
	userDataService := service2.

	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
