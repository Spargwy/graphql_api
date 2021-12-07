package resolvers

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
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

	err = token.Claims.(jwt.MapClaims).Valid()

	if err != nil {
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID

	tokenString, err = token.SignedString(secret)
	if err != nil {
		return
	}

	return
}

func (r *Resolver) SendCode(phoneNumber string) error {
	codeLength := 4
	code := generateRandomCode(codeLength)

	if os.Getenv("TEST") == "true" {
		log.Print(code)
		usersCodes[phoneNumber] = code

		return nil
	}

	usersCodes[phoneNumber] = code
	client := r.RestClient

	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(code)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}

	return nil
}
