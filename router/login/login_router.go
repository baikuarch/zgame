package login

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/baikuarch/zgame/db/redis"
	"github.com/baikuarch/zgame/msg"
	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (r *LoginRouter) Handle(request ziface.IRequest) {
	//读取客户端的数据
	//fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	m := &msg.LoginReq{}
	err := proto.Unmarshal(request.GetData(), m)
	if err != nil {

	}
	_, err = redis.SET("name", "1")
	if err != nil {

	}

	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()), ",m.name=", m.Name, ",m.password=", m.Password)

}
