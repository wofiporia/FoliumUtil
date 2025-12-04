package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wofiporia/foliumutil/fvalidator"
)

func TestValidateUsername(t *testing.T) {
	// Test valid usernames
	validUsernames := []string{
		"user",
		"User123",
		"test_user",
		"a123456789012345",
	}

	for _, username := range validUsernames {
		err := fvalidator.ValidateUsername(username)
		require.NoError(t, err, "username %s should be valid", username)
	}

	// Test invalid usernames
	invalidUsernames := []string{
		"123user",                          // starts with number
		"_user",                            // starts with underscore
		"u",                                // too short
		"user-with-dash",                   // contains invalid character (-)
		"user with space",                  // contains space
		"verylongusernamethatexceedslimit", // too long
		"",                                 // empty
	}

	for _, username := range invalidUsernames {
		err := fvalidator.ValidateUsername(username)
		require.Error(t, err, "username %s should be invalid", username)
	}
}

func TestValidatePassword(t *testing.T) {
	// Test valid passwords
	validPasswords := []string{
		"Password123",
		"TestPass456",
		"ValidPass789",
		"MySecure123",
		"Admin@Pass123",
	}

	for _, password := range validPasswords {
		err := fvalidator.ValidatePassword(password)
		require.NoError(t, err, "password %s should be valid", password)
	}

	// Test invalid passwords
	invalidPasswords := []string{
		"short",              // too short
		"alllowercase123",    // no uppercase
		"ALLUPPERCASE123",    // no lowercase
		"NoDigitsHere",       // no numbers
		"Short1",             // too short even with all requirements
		"toolongpassword123", // too long
		"",                   // empty
		"Password123#",       // invalid character (not in allowed set)
	}

	for _, password := range invalidPasswords {
		err := fvalidator.ValidatePassword(password)
		require.Error(t, err, "password %s should be invalid", password)
	}
}

func TestValidateEmail(t *testing.T) {
	// Test valid emails
	validEmails := []string{
		"user@example.com",
		"test.email@domain.co.uk",
		"user+tag@example.org",
		"firstname.lastname@company.com",
		"a@b.co",
	}

	for _, email := range validEmails {
		err := fvalidator.ValidateEmail(email)
		require.NoError(t, err, "email %s should be valid", email)
	}

	// Test invalid emails
	invalidEmails := []string{
		"invalid-email",          // no @
		"@domain.com",            // no local part
		"user@",                  // no domain
		"user@.com",              // invalid domain
		"user..name@example.com", // consecutive dots
		"user@domain..com",       // consecutive dots in domain
		"",                       // empty
		"a",                      // too short
	}

	for _, email := range invalidEmails {
		err := fvalidator.ValidateEmail(email)
		require.Error(t, err, "email %s should be invalid", email)
	}
}

func TestValidateString(t *testing.T) {
	// Test valid strings
	testCases := []struct {
		value     string
		minLength int
		maxLength int
		expectErr bool
	}{
		{"hello", 3, 10, false},           // valid
		{"hi", 3, 10, true},               // too short
		{"this is too long", 3, 10, true}, // too long
		{"", 1, 5, true},                  // empty
		{"abc", 3, 3, false},              // exact min and max
	}

	for _, tc := range testCases {
		err := fvalidator.ValidateString(tc.value, tc.minLength, tc.maxLength)
		if tc.expectErr {
			require.Error(t, err, "value '%s' with length %d-%d should be invalid", tc.value, tc.minLength, tc.maxLength)
		} else {
			require.NoError(t, err, "value '%s' with length %d-%d should be valid", tc.value, tc.minLength, tc.maxLength)
		}
	}
}

func TestValidateMultipleRules(t *testing.T) {
	// Test combination of validation rules
	username := "testuser123"
	password := "ValidPass123"
	email := "testuser@example.com"

	// All should be valid
	require.NoError(t, fvalidator.ValidateUsername(username))
	require.NoError(t, fvalidator.ValidatePassword(password))
	require.NoError(t, fvalidator.ValidateEmail(email))

	// Test with invalid data
	invalidUsername := "123user"
	invalidPassword := "weakpass"
	invalidEmail := "not-an-email"

	require.Error(t, fvalidator.ValidateUsername(invalidUsername))
	require.Error(t, fvalidator.ValidatePassword(invalidPassword))
	require.Error(t, fvalidator.ValidateEmail(invalidEmail))
}
