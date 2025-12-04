package fvalidator

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername    = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{3,15}$`).MatchString
	hasLowercase       = regexp.MustCompile(`[a-z]`).MatchString
	hasUppercase       = regexp.MustCompile(`[A-Z]`).MatchString
	hasDigit           = regexp.MustCompile(`\d`).MatchString
	validPasswordChars = regexp.MustCompile(`^[A-Za-z\d@$!%*?&]{8,16}$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if !isValidUsername(value) {

		return fmt.Errorf("must contain only letters, numbers, and underscores and be between 4-16 characters long")

	}
	return nil
}

func ValidatePassword(value string) error {
	if !validPasswordChars(value) {
		return fmt.Errorf("must be between 8-16 characters long and contain only letters, numbers, and special characters @$!%%*?&")
	}
	if !hasLowercase(value) {
		return fmt.Errorf("must contain at least one lowercase letter")
	}
	if !hasUppercase(value) {
		return fmt.Errorf("must contain at least one uppercase letter")
	}
	if !hasDigit(value) {
		return fmt.Errorf("must contain at least one number")
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 320); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}
