package main

import (
	"fmt"
	"log"
	"task4/config"
	"task4/router"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func init() {
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}

	// 初始化数据库
	_, db_err := config.InitDB(config.GetConf())
	if db_err != nil {
		log.Fatalf("%+v", errors.WithStack(db_err))
	}

	fmt.Println("----配置加载成功----")
}

func main() {

	//加载配置
	r := gin.Default()
	router.SetRouters(r)
	err := r.Run(":8090")
	if err != nil {
		panic("Failed to start Gin server: " + err.Error())
	}
}
