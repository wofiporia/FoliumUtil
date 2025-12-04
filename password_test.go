package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wofiporia/foliumutil/fpassword"
	"github.com/wofiporia/foliumutil/frandom"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := frandom.RandomString(6)

	hashedPassword1, err := fpassword.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = fpassword.CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := frandom.RandomString(6)
	err = fpassword.CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := fpassword.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
