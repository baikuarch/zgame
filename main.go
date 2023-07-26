package main

import (
	"github.com/aceld/zinx/znet"
	"github.com/baikuarch/zgame/db"
	"github.com/baikuarch/zgame/msg"
	"github.com/baikuarch/zgame/router/login"
)

func main() {

	//	初始化DB
	newDb := db.NewDb()
	newDb.InitDb([]int8{db.ENUM_Redis})

	//	创建服务器
	server := znet.NewServer()

	//	添加路由
	server.AddRouter(uint32(msg.MsgIDEnum_LOGIN_REQ.Number()), &login.LoginRouter{})

	//	启动服务
	server.Serve()

}
