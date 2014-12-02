package helpers

import (
	"os"
	"strconv"
)

func FetchIndex() (int, error) {
	index := os.Getenv("INSTANCE_INDEX")
	return strconv.Atoi(index)
}
