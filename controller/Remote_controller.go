package controller

import (
	"../config"
	"../service"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)



//向网页中写入html传入的状态值
func WriteControlSituationRedis(ctx iris.Context){
	ControlStateStr :=""
	Keyname := "ControlSituation"
	checkCode := ctx.PostValue("checkcode")
	KeyValue  := ctx.PostValue("ControlSituation")
	if checkCode==config.CheckIDCode{
		service.RedisWrite(Keyname,KeyValue)
		fmt.Println("Success WriteRedis")
		ControlStateStr="控制信息发送成功"
	}else{
		fmt.Println("writeError")
		ControlStateStr="提交失败，请检查校验码是否正确或者联系管理员"

	}

	ctx.ViewData("ControlState", ControlStateStr)
	ctx.View("RemoteControl.html")



}

//建立WebSocket连接并与现场服务器进行通信
func WebScoketRedisToOPC(c websocket.Connection){
	c.OnMessage(func(data []byte) {

		message := string(data)
		if message==config.CheckIDCode{
			returnMsg := service.QueryRedisAndCleanSituation("ControlSituation")

			c.To(websocket.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + returnMsg)) // broadcast to all clients except this
			c.EmitMessage([]byte("Server:"+returnMsg)) // writes to itself
		}else{
			returnMsg := "token请求不正确，请校验后发送"
			c.To(websocket.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + returnMsg)) // broadcast to all clients except this
			c.EmitMessage([]byte(returnMsg)) // writes to itself
		}

	})

	c.OnDisconnect(func() {
		fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
	})

}



//建立WebSocket连接并与现场服务器进行通信
func WebScoketUnityToRedis(c websocket.Connection,StituationValue string,checkCode string){
	c.OnMessage(func(data []byte) {

		message := string(data)
		if message==checkCode{
			service.RedisWrite("ControlSituation",StituationValue)

			c.To(websocket.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " )) // broadcast to all clients except this
			c.EmitMessage([]byte("Me: " +"ValueWriteOK")) // writes to itself
		}else{
			returnMsg := "token请求不正确，请校验后发送"
			c.To(websocket.Broadcast).EmitMessage([]byte("Message from: " + c.ID() + "-> " + returnMsg)) // broadcast to all clients except this
			c.EmitMessage([]byte("Me: " + returnMsg)) // writes to itself
		}

	})
	c.OnDisconnect(func() {
		fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
	})

}
