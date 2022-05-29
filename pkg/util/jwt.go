package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jamesluo111/gin-blog/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	//过期时间3小时
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	//生成token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//对token进行签发
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//对token解析
func ParseToken(token string) (*Claims, error) {
	//对签名进行反解析
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		//对token进行验证解析,最后生成Claims
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
