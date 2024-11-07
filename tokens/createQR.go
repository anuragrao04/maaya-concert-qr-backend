package tokens

import (
	"errors"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/anuragrao04/maaya-concert-qr/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
)

var JWT_SECRET []byte

func CreateQR(user *models.User) (qrFilePath string, err error) {
	// first we create a JWT
	JWTString, err := CreateJWT(user)
	if err != nil {
		return "", err
	}
	// encode this into a QR with bullshita's colors
	qrFilePath = "./tempTickets/" + uuid.New().String() + ".png"
	err = qrcode.WriteColorFile(JWTString, qrcode.Low, 5120, color.RGBA{10, 18, 58, 255}, color.NRGBA{15, 213, 178, 255}, qrFilePath)
	// low error correction, to reduce the number of pixels in the qr code
	// there is 0 chance of qr code being damaged since it's on a screen.
	// so we practically don't need any error correction
	return qrFilePath, err
}

func CreateJWT(user *models.User) (JWTString string, err error) {
	ttl := time.Hour * 24 * 3 // 3 days
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(ttl).Unix(),
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
