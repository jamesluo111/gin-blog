package main

import (
	"context"
	"fmt"
	"github.com/jamesluo111/gin-blog/models"
	"github.com/jamesluo111/gin-blog/pkg/gredis"
	"github.com/jamesluo111/gin-blog/pkg/logging"
	"github.com/jamesluo111/gin-blog/pkg/setting"
	"github.com/jamesluo111/gin-blog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//endless.DefaultReadTimeOut = setting.ReadTimeOut
	//endless.DefaultWriteTimeOut = setting.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endpoint := fmt.Sprintf(":%d", setting.HttpPort)
	//server := endless.NewServer(endpoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err:%v", err)
	//}
	setting.Setup()
	models.Setup()
	logging.Setup()
	redisErr := gredis.SetUp()
	if redisErr != nil {
		panic(redisErr)
	}
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("ShutDown Server...")

	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cannel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
