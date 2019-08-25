package main

import (
	"./config"
	"./controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)


func main() {
    app := iris.Default()

    //与现场服务器进行通信
	wsOPC := websocket.New(websocket.Config{})
	app.Get(config.OPCWebScocketURL, wsOPC.Handler())
	wsOPC.OnConnection(func(c websocket.Connection) {
		controller.WebScoketRedisToOPC(c)
	})

	//与Unity进行通信
	wsUnity := websocket.New(websocket.Config{})
	app.Get(config.UnityWebScocketURL, wsOPC.Handler())
	wsUnity.OnConnection(func(c websocket.Connection) {
		controller.WebScoketUnityToRedis(c,config.UnityOrder,config.UnityCheckCode)
	})


    app.RegisterView(iris.HTML("./templates", ".html")) // 注册 view 目录
    app.StaticWeb("static", "./static")     // 初始化静态页面目录

    //进入主页
    app.Get("/", func(ctx iris.Context) {
        ctx.ViewData("name", "wx")
        ctx.View("base.html")
    })

    //CPS系统
    app.Get(config.CPSWebGLURL, func(ctx iris.Context) {
        ctx.View("index_WebGL.html")
    })

    //相机监测画面
    app.Get(config.RemoteControlURL, func(ctx iris.Context) {
        ctx.View("RemoteControl.html")
    })

    //远程监控
    app.Get(config.RealtimeVideo, func(ctx iris.Context) {
        ctx.View("camera.html")
    })

    //校验账号密码
    app.Post(config.UserIDCheckPost, func(ctx iris.Context){
    	controller.CheckUserInformation(ctx)
	})

	//遇到状态控制提交值后写入Redis
	app.Post(config.RemoteControlPost, func(ctx iris.Context){
		controller.WriteControlSituationRedis(ctx)
	})






    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(config.IrisAddr))
}




