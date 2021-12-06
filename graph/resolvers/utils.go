package resolvers

import (
	"math/rand"
	"os"
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

func generateJWT(userID int) (tokenString string, err error) {
	var secret = []byte(os.Getenv("SECRET_FOR_JWT"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID

	tokenString, err = token.SignedString(secret)
	if err != nil {
		return
	}
	return
}
