package utils

import (
	"strings"
)

func TrimEmailAddress(email string) string {
	s := strings.Split(email, "@")
	f := strings.ToLower(s[0])
	f = strings.Replace(f, ".", "", 1000)
	return f + "@" + s[1]
}
