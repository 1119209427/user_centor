package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
type EmailClaims struct {
	UserID uint `json:"user_id"`
	Email string `json:"email"`
	PassWord string `json:"pass_word"`
	OperationType uint `json:"operation_type"`
	jwt.StandardClaims
}
/*生产jwt的步骤
1.先定义一个standerclaims的结构体，可以含有自己需要的属性
2.然后在封装产生的方法中，初始化这个结构体
3.调用jwt.NewWithClaims方法
4.最后调用 上一个方法生成的对象中的signedstring方法
*/

func GenerateEmailToken(email,password string,userId,operationType uint)(string,error){
	nowTime:=time.Now()
	expireTime:=nowTime.Add(15*time.Minute)

	claims:=EmailClaims{
		UserID: userId,
		Email: email,
		PassWord: password,
		OperationType: operationType,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jing_dong",

		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token,err:=withClaims.SignedString(jwtSecret)
	return token,err

}
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
