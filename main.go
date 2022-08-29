package main

import (
	"fmt"
	_ "gin-blog/docs"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	router := routers.InitRouter()
	//router.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	// 使用endless实现热更新  https://segmentfault.com/a/1190000013757098
	//endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
