package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/onsi/grace/handlers"
	"github.com/onsi/grace/routes"
	"github.com/onsi/grace/vcap_application_parser"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/rata"
)

func main() {
	var chatty bool
	flag.BoolVar(&chatty, "chatty", false, "make grace chatty")
	flag.Parse()

	logger := lager.NewLogger("grace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	logger.Info("hello", lager.Data{"port": os.Getenv("PORT")})
	handler, err := rata.NewRouter(routes.Routes, handlers.New(logger))
	if err != nil {
		logger.Fatal("router.creation.failed", err)
	}

	if chatty {
		index, err := vcap_application_parser.GetIndex()

		go func() {
			t := time.NewTicker(time.Second)
			for {
				<-t.C
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to fetch index: %s\n", err.Error())
				} else {
					fmt.Printf("This is Grace on index: %d\n", index)
				}
			}
		}()
	}

	server := ifrit.Envoke(http_server.New(":"+os.Getenv("PORT"), handler))
	err = <-server.Wait()
	if err != nil {
		logger.Error("farewell", err)
	}
	logger.Info("farewell")
}
