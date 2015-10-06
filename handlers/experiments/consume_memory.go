package experiments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var bag map[int64][]byte
var fmLock *sync.Mutex

func init() {
	Experiments["consume_memory"] = Experiment{
		Name:        "Consume Memory",
		Description: "Consumes 32MB of memory.  Frees it after 15 seconds.",
		Handler:     http.HandlerFunc(FillMemory),
	}

	bag = map[int64][]byte{}
	fmLock = &sync.Mutex{}
}

func FillMemory(w http.ResponseWriter, r *http.Request) {
	thirtyTwoM := []byte(strings.Repeat("7", 1024*1024*32))
	id := time.Now().UnixNano()

	fmLock.Lock()
	bag[id] = thirtyTwoM
	size := len(bag)
	fmLock.Unlock()

	fmt.Printf("Allocated %d: size is now %d\n", id, size*32)

	t := time.NewTimer(15 * time.Second)

	go func() {
		<-t.C
		fmLock.Lock()
		delete(bag, id)
		size := len(bag)
		fmLock.Unlock()
		fmt.Printf("Deallocated %d: size is now %d\n", id, size*32)
		runtime.GC()
		debug.FreeOSMemory()
	}()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Fprintf(w, "<html><body>Allocated %d: size is now %d<br>", id, size*32)
	body, _ := json.MarshalIndent(memStats, "<br>", "&nbsp;&nbsp;&nbsp;&nbsp;")
	w.Write(body)
	fmt.Fprintf(w, "</body></html>")
}
