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

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate gengin go-api-template.api
//go:generate gen-swagger --local_api= go-api-template.api

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

	pprof.Register(e)
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
