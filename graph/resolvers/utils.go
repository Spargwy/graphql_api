package resolvers

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateRandomCode(length int) (code string) {
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < length; i++ {
		code += strconv.FormatInt(int64(rand.Int31n(10)), 10)
	}
	return code
}

func generateJWT() (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	var secret []byte
	tokenString, err = token.SignedString(secret)
	if err != nil {
		return
	}
	return
}
