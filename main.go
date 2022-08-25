package main

import (
	"fmt"
	_ "gin-blog/docs"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	//router.Run(fmt.Sprintf(":%d", setting.HTTPPort))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	// 使用endless实现热更新  https://segmentfault.com/a/1190000013757098
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	logging.Info("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	logging.Error("Server err: %v", err)
	//}
}
