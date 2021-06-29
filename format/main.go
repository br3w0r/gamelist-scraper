package format

import "strings"

func SinglePlatform(uPlatform string) string {
	uPlatform = uPlatform[41:]
	breakIndex := strings.Index(uPlatform, "\n")

	return uPlatform[:breakIndex]
}

func YearReleased(uDate string) string {
	return uDate[len(uDate)-4:]
}
