package experiments

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	Experiments["write_temp"] = Experiment{
		Name:        "Write temp",
		Description: "Writes to /home/vcap/tmp",
		Handler:     http.HandlerFunc(WriteTemp),
	}
}

func WriteTemp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><head></head><body>`))
	defer w.Write([]byte(`</body></html>`))

	err := ioutil.WriteFile("/home/vcap/tmp/foo", []byte("HELLO"), 0777)
	if err != nil {
		fmt.Fprintf(w, "WRITE ERROR:%s", err.Error())
		return
	}

	data, err := ioutil.ReadFile("/home/vcap/tmp/foo")
	if err != nil {
		fmt.Fprintf(w, "READ ERROR:%s", err.Error())
		return
	}

	fmt.Fprintf(w, "DATA: %s", string(data))

}
