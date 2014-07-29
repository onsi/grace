package main

import (
	"os"

	"github.com/onsi/grace/handlers"
	"github.com/onsi/grace/routes"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/rata"
)

func main() {
	logger := lager.NewLogger("grace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	logger.Info("hello", lager.Data{"port": os.Getenv("PORT")})
	handler, err := rata.NewRouter(routes.Routes, handlers.New(logger))
	if err != nil {
		logger.Fatal("router.creation.failed", err)
	}
	server := ifrit.Envoke(http_server.New(":"+os.Getenv("PORT"), handler))
	err = <-server.Wait()
	if err != nil {
		logger.Error("farewell", err)
	}
	logger.Info("farewell")
}
