package validate

import (
	"regexp"
)

// credits to this -> https://ihateregex.io/expr/

// holds the strings that will be passed to the func that
// validates the value
const (
	RegexEmail    = `[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`
	RegexUsername = `^[a-z0-9_-]{3,15}$`
)

func MatchesRegex(regex string, value string) bool {
	isValid := regexp.MustCompile(regex).MatchString(value)
	return isValid
}
