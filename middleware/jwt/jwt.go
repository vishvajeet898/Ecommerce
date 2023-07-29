package jwt

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	"reflect"
	"strings"
	"time"
)

const (
	// JWTSecret TODO Fetch this from env
	JWTSecret  = "123ACD"
	JwtExpHour = 1
	JwtExpMin  = 0
	JwtExpSec  = 30
	UserScope  = "user"
	AdminScope = "admin"
)

type customClaims struct {
	UserId   string `json:"userID"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

func NewAuthMiddleware(scopes []string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(c context.Context, request interface{}) (interface{}, error) {

			//Extracting token from interface
			requestValue := reflect.ValueOf(request)
			tokenString := requestValue.FieldByName("JWT").String()

			token := strings.Split(tokenString, " ")
			if len(token) != 2 {
				return nil, fmt.Errorf("invlid length")
			}
			if token[0] != "Bearer" {
				return nil, fmt.Errorf("invlid berarer")
			}

			jwtToken, err := ParseToken(token[1])
			if err != nil {
				return nil, err
			}

			claims, err := GetClaims(jwtToken)
			if err != nil {
				return nil, err
			}

			userID, _ := claims["userID"].(string)
			userType, _ := claims["userType"].(string)

			fmt.Printf("\nFromMiddleWare \nUSERID : %v", userID)
			//adding userID to context to retrieve in makeEndpoint and pass forward to service
			c = context.WithValue(c, "userID", userID)

			//Checking if user is allowed to access this api or not
			for _, scope := range scopes {
				if scope == userType {
					return next(c, request)
				}
			}

			return nil, fmt.Errorf("NOT ENOUGH PERMISSION")
		}
	}
}

func NewToken(userID string, userType string) (string, error) {
	customClaim := customClaims{
		UserId:   userID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(JwtExpHour) + time.Minute*time.Duration(JwtExpMin) + time.Second*time.Duration(JwtExpSec)).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)
	sToken, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}
	return sToken, nil
}

// ParseToken parse a token
func ParseToken(tokenString string) (*jwt.Token, error) {
	var token *jwt.Token
	var err error
	token, err = parseHS256(tokenString, token)

	//TODO ErrTokenExpired
	if err != nil && err.Error() != "Token is expired" {
		token, err = parseHS256(tokenString, token)
	}

	return token, err
}

// GetClaims get claims information
func GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}
	err := token.Claims.(jwt.MapClaims).Valid()
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

func parseHS256(tokenString string, token *jwt.Token) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	return token, err
}
