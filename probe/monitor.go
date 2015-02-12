package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/onsi/gomega/format"
	"github.com/pivotal-cf-experimental/veritas/say"
)

func MonitorCommand() Command {
	var (
		route         string
		pollInterval  time.Duration
		batchInterval time.Duration
	)

	flagSet := flag.NewFlagSet("monitor", flag.ExitOnError)
	flagSet.StringVar(&route, "route", "http://grace.10.244.0.34.xip.io", "route to grace")
	flagSet.DurationVar(&pollInterval, "pollInterval", 30*time.Millisecond, "interval to pull")
	flagSet.DurationVar(&batchInterval, "batchInterval", 5*time.Second, "interval to batch collect indices")

	return Command{
		Name:        "monitor",
		Description: "Monitor Grace",
		FlagSet:     flagSet,
		Run: func(args []string) {
			monitor(route, pollInterval, batchInterval)
		},
	}
}

func monitor(route string, pollInterval, batchInterval time.Duration) {
	say.Println(0, "Monitoring %s every %s", say.Green(route), say.Green("%s", pollInterval))
	// http.DefaultClient.Timeout = 200 * time.Millisecond

	ticker := time.NewTicker(pollInterval)

	startTime := time.Now()
	roundTime := time.Now()
	indices := map[int]int{}
	requests := 0
	succesfulRequests := 0
	for {
		<-ticker.C
		requests++
		resp, err := http.Get(route + "/index")

		if err != nil {
			say.Println(0, "%s: %s", say.Yellow("%s", time.Since(startTime)), say.Red(fmt.Sprintf("Error: %s", err.Error())))
			continue
		}

		if resp.StatusCode != http.StatusOK {
			say.Println(0, "%s: %s", say.Yellow("%s", time.Since(startTime)), say.Red(fmt.Sprintf("Invalid Status Code: %d", resp.StatusCode)))
			say.Println(1, say.Red(format.Object(resp.Header, 0)))
			continue
		}

		succesfulRequests++
		indexStr, _ := ioutil.ReadAll(resp.Body)
		index, _ := strconv.Atoi(string(indexStr))

		indices[index]++
		resp.Body.Close()

		if time.Since(roundTime) >= batchInterval {
			say.Println(0, "%s: %d/%d %s", say.Yellow("%s", time.Since(startTime)), succesfulRequests, requests, sortedIndices(indices))
			indices = map[int]int{}
			requests = 0
			succesfulRequests = 0
			roundTime = time.Now()
		}
	}
}

func sortedIndices(counts map[int]int) string {
	indices := []int{}
	for index := range counts {
		indices = append(indices, index)
	}

	sort.Ints(indices)

	out := []string{}

	for _, index := range indices {
		out = append(out, fmt.Sprintf("%d: %d", index, counts[index]))
	}

	return strings.Join(out, ", ")
}
