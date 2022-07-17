package utils

import (
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// used to generate JWTs
var jwtKey = []byte("supersecretkey")

// struct payload of JWT 
type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role	 string `json:"role"`
	jwt.StandardClaims
}

// generate token with HS256 Signing, expiration 1 hour
func GenerateJWT(email string, username string, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims:= &JWTClaim{
		Email: email,
		Username: username,
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// validate token, check if expired
func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return claims, err
}

func CheckUserType(c *gin.Context, role string)(err error){
	
	userRole := c.GetString("role")
	err = nil

	// gives admin access to all
	if userRole == "admin" {
		return err 
	}

	if userRole != role {
		err = errors.New("Unauthorized to access this resource")
        return err
    }

	return err

}