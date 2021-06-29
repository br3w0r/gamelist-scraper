package format

import (
	"strconv"
	"strings"
)

func SinglePlatform(uPlatform string) string {
	uPlatform = uPlatform[41:]
	breakIndex := strings.Index(uPlatform, "\n")

	return uPlatform[:breakIndex]
}

func YearReleased(uDate string) (int, error) {
	val, err := strconv.ParseInt(uDate[len(uDate)-4:], 0, 0)
	return int(val), err
}
