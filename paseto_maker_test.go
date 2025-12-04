package util_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/wofiporia/foliumutil/frandom"
	"github.com/wofiporia/foliumutil/ftoken"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := ftoken.NewPasetoMaker(frandom.RandomString(32))
	require.NoError(t, err)

	username := frandom.RandomString(10)
	role := "user"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, role, payload.Role)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := ftoken.NewPasetoMaker(frandom.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(frandom.RandomString(10), "user", -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ftoken.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestPasetoWrongTokenType(t *testing.T) {
	maker1, err := ftoken.NewPasetoMaker(frandom.RandomString(32))
	require.NoError(t, err)

	username := frandom.RandomString(10)
	duration := time.Minute

	// 使用maker1创建一个有效的令牌
	validToken, payload, err := maker1.CreateToken(username, "user", duration)
	require.NoError(t, err)
	require.NotEmpty(t, validToken)
	require.NotEmpty(t, payload)

	// 创建第二个Maker，它使用一个完全不同的密钥
	maker2, err := ftoken.NewPasetoMaker(frandom.RandomString(32)) // 密钥不同
	require.NoError(t, err)

	// 尝试使用第二个Maker（错误的密钥）来验证第一个Maker创建的令牌
	// 这应该会导致验证失败，因为密钥不匹配
	payload, err = maker2.VerifyToken(validToken)
	require.Error(t, err)
	// 确认返回的错误是“无效令牌”类型
	require.ErrorIs(t, err, ftoken.ErrInvalidToken)
	require.Nil(t, payload, "当验证密钥错误时，不应返回任何有效载荷")

}
