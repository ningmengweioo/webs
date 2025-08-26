package main

//根据自身条件,创建对应的go.mod
import (
	"demo_web3/webs/go_task1/database"
	"demo_web3/webs/go_task1/model"
	"demo_web3/webs/go_task1/task3/gormfile"
)

func main() {
	database.InitDB(&model.User{})
	gormfile.Run(database.DB)
}
