package tokens

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/anuragrao04/maaya-concert-qr/models"
	"github.com/golang-jwt/jwt/v5"
	qrcode "github.com/skip2/go-qrcode"
)

var JWT_SECRET []byte

func CreateQR(user *models.User) (qr []byte, err error) {
	// first we create a JWT

	JWTString, err := CreateJWT(user)
	if err != nil {
		return nil, err
	}
	// encode this into a QR with bullshita's colors

	qr, err = qrcode.Encode(JWTString, qrcode.Highest, 1024) // highest error recovery, and a 1024x1024 resolution
	return
}

func CreateJWT(user *models.User) (JWTString string, err error) {
	ttl := time.Hour * 24 * 3 // 3 days
	claims := jwt.MapClaims{
		"name":  user.Name,
		"srn":   user.SRN,
		"prn":   user.PRN,
		"email": user.Email,
		"exp":   time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Println("error signing token", err)
		return "", err
	}
	return tokenString, nil
}

func LoadPrivateKey() error {
	if os.Getenv("JWT_SECRET") == "" {
		return errors.New("JWT_SECRET not found in ENV")
	}
	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
	return nil
}
