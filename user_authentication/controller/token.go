package controller

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hgfischer/go-otp"
	"golang.org/x/crypto/bcrypt"
)

var letterBytes = "0123456789"

//var sampleSecretKey = []byte(os.Getenv("sampleSecretKey"))

func GenerateJWT(username string) (string, error) {
	var sampleSecretKey = []byte(os.Getenv("sampleSecretKey"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		err := fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(w http.ResponseWriter, r *http.Request) (err error, b bool) {
	var sampleSecretKey = []byte(os.Getenv("sampleSecretKey"))
	b = true
	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		b = false
		return
	}
	token, _ := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return sampleSecretKey, nil
	})

	if token == nil {
		fmt.Fprintf(w, "invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if claims == nil {
		fmt.Fprintf(w, "Error the token is nil")
	}
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return errors.New("token error"), b
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		fmt.Fprintf(w, "token expired")
		return errors.New("token error"), b
	}

	return nil, b
}
func GenerateResetToken(n int) string {
	//var letterBytes = []byte(os.Getenv("letterBytes"))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func ValidateResetToken(tokendb string, token string) error {
	return bcrypt.CompareHashAndPassword([]byte(tokendb), []byte(token))
}

func OtpGenerator() *otp.TOTP {
	secrete := os.Getenv("sampleSecretKey")
	totp := otp.TOTP{Secret: secrete, Length: 6, IsBase32Secret: true}
	return &totp
}
