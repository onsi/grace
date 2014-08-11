package experiments

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func init() {
	Experiments["fill_inodes"] = Experiment{
		Name:        "Fill inodes",
		Description: "Writes 201,000 tiny files, reporting the number of files written",
		Handler:     http.HandlerFunc(FillInodes),
	}
}

func FillInodes(w http.ResponseWriter, r *http.Request) {
	oneByte := []byte("1")
	w.Write([]byte(`<html><head></head><body>`))

	total := 1
	for outer := 0; outer < 201; outer++ {
		w.Write([]byte("<div>Writing another 1000 files..."))
		for i := 0; i < 1000; i++ {
			f, err := os.Create(fmt.Sprintf("%d", total))
			if err != nil {
				w.Write([]byte(fmt.Sprintf("<div>Failed to write file #%d</div>", total)))
			}
			f.Write(oneByte)
			f.Close()
			total += 1
		}
		time.Sleep(100 * time.Millisecond)

		w.Write([]byte(fmt.Sprintf(" total is at: %d</div>", total)))
	}
	w.Write([]byte(`</body></html>`))
}
