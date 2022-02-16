/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2021-09-29 14:16:31
 * @LastEditTime: 2021-12-30 18:27:49
 */
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	"xiaoshuo/internal/routers"
	"xiaoshuo/pkg/log/logrus"

	_ "go.uber.org/automaxprocs"
)

func main() {
	logrus.NewLogger()
	logrus.NewLeavMessage()
	logrus.NewAskBook()
	httpPort := "9999"
	addr := fmt.Sprintf(":%s", httpPort)
	readTimeout := time.Second * 5
	writeTimeout := time.Second * 5

	s := &http.Server{
		Addr:           addr,
		Handler:        routers.InitRouter(),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
