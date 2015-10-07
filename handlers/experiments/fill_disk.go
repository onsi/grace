package experiments

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
)

func init() {
	Experiments["fill_disk"] = Experiment{
		Name:        "Fill the disk",
		Description: "Writes data to the working directory forever, reporting the amount of data written",
		Handler:     http.HandlerFunc(FillDisk),
	}
}

func FillDisk(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("./bloat")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to create /dev/shm: %s", err.Error())))
		return
	}

	c := 1024 * 1024
	b := make([]byte, c)

	fmt.Println(bytes.Equal(b, make([]byte, c)))

	w.Write([]byte(`<html><head></head><body>`))

	total := 0
	for {
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		w.Write([]byte("<div>Writing another 1MB..."))
		n, err := f.Write(b)
		if err != nil {
			w.Write([]byte(fmt.Sprintf(" failed to write! %s", err.Error())))
			return
		}
		err = f.Sync()
		if err != nil {
			w.Write([]byte(fmt.Sprintf(" failed to sync! %s", err.Error())))
			return
		}

		total += n
		w.Write([]byte(fmt.Sprintf(" total is at: %.3f MB</div>", float64(total)/1024.0/1024.0)))
		f, ok := w.(http.Flusher)
		if ok {
			f.Flush()
		}
	}
}
