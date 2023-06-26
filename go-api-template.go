package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"go-api-template/config"
	"go-api-template/middleware"
	"go-api-template/model"
	"go-api-template/pkg/logger"
	"go-api-template/pkg/prometheus"
	"go-api-template/pkg/swagger"
	"go-api-template/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate goctl api plugin -p gengin -api go-api-template.api -dir .
//go:generate goctl api plugin -p "goctl-swagger swagger -filename asset/swagger/swagger.json" -api go-api-template.api -dir .

func main() {
	flag.Parse()

	configFile := fmt.Sprintf("etc/go-api-template-%s.yaml", config.Env)
	e := gin.New()

	e.Use(
		middleware.RequestId,
		middleware.RequestLog,
		gin.Recovery(),
		middleware.CorsMiddleware,
		middleware.ApiHitRecord,
	)

	config.Setup(configFile)
	logger.Setup()
	prometheus.Setup(e)
	//redis.Setup()
	model.Setup()
	//query.SetDefault(model.DB())
	routes.Setup(e)
	swagger.Setup(e)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ServerConf.Host, config.ServerConf.Port),
		Handler: e,
	}
	go func() {
		fmt.Println("listen to ", server.Addr)
		server.ListenAndServe()
	}()

	wait := config.RegisterCloseFn(func() {
		defer logrus.Warning("closed api server")

		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()
		server.Shutdown(ctx)
	})
	wait()
}
