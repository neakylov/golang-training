package NotVerySecure

import "regexp"

func alphanumeric(str string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")

	return r.MatchString(str)
}
