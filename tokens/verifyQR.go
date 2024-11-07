package tokens

import "github.com/golang-jwt/jwt/v5"

func VerifyQR(jwtString string) (valid bool, claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return false, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims, nil
	} else {
		return false, nil, nil
	}
}
