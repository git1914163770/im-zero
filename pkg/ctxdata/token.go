package ctxdata

import "github.com/golang-jwt/jwt/v4"

const Identify = "uid"

// GenJwtToken 生成JWT令牌。
// 参数:
//
//	secretKey: 用于签名的密钥。
//	iat: 令牌发行时间。
//	seconds: 令牌有效期的秒数。
//	uid: 用户标识符。
func GenJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
