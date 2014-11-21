package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/onsi/grace/handlers"
	"github.com/onsi/grace/helpers"
	"github.com/onsi/grace/routes"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/rata"
)

func main() {
	var chatty bool
	var upFile string
	var exitAfter time.Duration
	var exitAfterCode int

	flag.BoolVar(&chatty, "chatty", false, "make grace chatty")
	flag.StringVar(&upFile, "upFile", "", "a file to write to (lives under /tmp)")
	flag.DurationVar(&exitAfter, "exitAfter", 0, "if set, grace will exit after this duration")
	flag.IntVar(&exitAfterCode, "exitAfterCode", 0, "exit code to emit with exitAfter")

	flag.Parse()

	logger := lager.NewLogger("grace")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("hello", lager.Data{"port": port})
	handler, err := rata.NewRouter(routes.Routes, handlers.New(logger))
	if err != nil {
		logger.Fatal("router.creation.failed", err)
	}

	if chatty {
		index, err := helpers.FetchIndex()

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

	if upFile != "" {
		f, err := os.Create("/tmp/" + upFile)
		if err != nil {
			logger.Fatal("upfile.creation.failed", err)
		}
		_, err = f.WriteString("Grace is up")
		if err != nil {
			logger.Fatal("upfile.creation.failed", err)
		}
	}

	if exitAfter != 0 {
		go func() {
			time.Sleep(exitAfter)
			fmt.Println("timebomb... farewell")
			os.Exit(exitAfterCode)
		}()
	}

	server := ifrit.Envoke(http_server.New(":"+port, handler))
	err = <-server.Wait()
	if err != nil {
		logger.Error("farewell", err)
	}
	logger.Info("farewell")
}
