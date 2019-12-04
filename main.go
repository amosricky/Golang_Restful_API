package main

import (
	"Golang_Restful_API/pkg/models"
	"Golang_Restful_API/pkg/routers"
	"Golang_Restful_API/pkg/setting"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"
)

func setLogrus()  {

	if setting.ServerSetting.RunMode == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	for{
		filePath := fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.LogSetting.LogSavePath)
		fileName := fmt.Sprintf("%s%s.%s",
			setting.LogSetting.LogPrefix,
			time.Now().Format(setting.LogSetting.TimeFormat),
			setting.LogSetting.LogFileExtension,
		)
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil{
			logrus.Info("Failed to log to file, using default stderr")
			break
		}
		file, err := os.OpenFile(path.Join(filePath, fileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			logrus.SetOutput(file)
		} else {
			logrus.Info("Failed to log to file, using default stderr")
		}
		break
	}
}

func main() {
	setting.Setup()
	models.Setup()
	setLogrus()

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
			logrus.Infof("Listen And Server: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	logrus.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown: %v", err)
	}

	models.CloseDB()
	logrus.Infof("Server exiting")
}