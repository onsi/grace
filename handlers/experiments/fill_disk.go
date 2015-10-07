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

	total := 1
	for {
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		w.Write([]byte("<div>Writing another 1MB..."))
		f.Write(b)
		f.Sync()
		w.Write([]byte(fmt.Sprintf(" total is at: %dMB</div>", total)))
		total += 1
	}
}
