package validator

import (
	"errors"
	"regexp"
)

func Validatephonenumbr(con_num string) error {
	re := regexp.MustCompile(`^\d{10}$`)
	if !re.MatchString(con_num) {
		return errors.New("contact number should be of 10 digit ")
	}
	return nil
}

func Validateisbn(isbn string) error {
	re := regexp.MustCompile(`^\d{13}$`)
	if !re.MatchString(isbn) {
		return errors.New("isbn should be of 13 digit")

	}
	return nil
}
