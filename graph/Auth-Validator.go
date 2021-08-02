package graph

import (
	"errors"

	"github.com/Az3z3l/GQL-SERVER/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret = []byte("NICMy0Wmzk9xyoYmV6um")
)


// CreateTokenEndpoint new jwt token
func CreateTokenEndpoint(s *model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": s.Username,
		"email":    s.Email,
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		Logger(err)
	}
	return tokenString
}

// errors.New("not found")

// ParseToken check if valid token
func ParseToken(tokenStr string) (string, string, error) {

	if tokenStr == "" {
		return "", "", errors.New("Authorization token must be present")
	}
	token, damn := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("There was an error")
		}
		return jwtSecret, nil
	})
	if damn != nil {
		return "", "", errors.New("Authorization token must be present")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, err := claims["username"].(string)
		email, err1 := claims["email"].(string)
		if !err || !err1 {
			return "", "", nil
		}
		return username, email, nil
	} else {
		return "", "", errors.New("Invalid authorization token")
	}

}
