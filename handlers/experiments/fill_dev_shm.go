package experiments

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func init() {
	Experiments["fill_dev_shm"] = Experiment{
		Name:        "Fill /dev/shm",
		Description: "Writes data to /dev/shm forever, reporting the amount of data written",
		Handler:     http.HandlerFunc(FillDevShm),
	}
}

func FillDevShm(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("/dev/shm/fill_er_up")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create /dev/shm: %s", err.Error())))
		return
	}

	sixtyFourK := []byte(strings.Repeat("7", 1024*1024))
	w.Write([]byte(`<html><head></head><body>`))

	total := 1
	for {
		w.Write([]byte("<div>Writing another 1MB..."))
		f.Write(sixtyFourK)
		f.Sync()
		w.Write([]byte(fmt.Sprintf(" total is at: %dMB</div>", total)))
		total += 1
	}
}
