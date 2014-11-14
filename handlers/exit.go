package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Exit struct {
}

func (p *Exit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	exitCode := r.URL.Query().Get(":code")
	code, err := strconv.Atoi(exitCode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("See ya!")
		os.Exit(code)
	}()
}
