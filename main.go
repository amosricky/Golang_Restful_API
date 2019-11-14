package main

import (
	"Golang_Restful_API/pkg/models"
	"Golang_Restful_API/pkg/routers"
	"Golang_Restful_API/pkg/setting"
	"fmt"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}