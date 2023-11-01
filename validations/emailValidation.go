package validations

import (
	"errors"
	"regexp"
	"strconv"
)


// ValidateEmail checks if the provided email address is valid
func ValidateEmail(email string) (bool, error) {
	// Regular expression for email validation
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(email) {
		return false, errors.New("Invalid email address")
	}
	return true, nil
}


// ValidatePassword checks if the provided password meets complexity requirements
func ValidatePassword(password string) error {
	// Add your password validation logic here
	return nil
}


// ValidateNumericValue checks if the provided string can be converted to a numeric value within a specific range
func ValidateNumericValue(value string, min, max int) error {
	num, err := strconv.Atoi(value)
	if err != nil {
		return errors.New("Invalid numeric value")
	}
	if num < min || num > max {
		return errors.New("Numeric value is out of range")
	}
	return nil
}

// ValidateStringFormat checks if the provided string matches a specific format or pattern
func ValidateStringFormat(input, pattern string) error {
	match, err := regexp.MatchString(pattern, input)
	if err != nil || !match {
		return errors.New("String format validation failed")
	}
	return nil
}

// ValidateURL checks if the provided string is a valid URL
func ValidateURL(url string) error {
	// Add your URL validation logic here
	return nil
}
