package validation

import "regexp"

var re = regexp.MustCompile(`^\+[\d]{12}$`)

func ValidatePhoneNumber(phone string) bool {
	return re.MatchString(phone)
}
