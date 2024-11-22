package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"ldap-auth/config"
	"ldap-auth/repositories"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(username, role string) (string, error) {
	jwtKey := []byte(config.C.JWTKey)
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthenticateUser(username, password string) (string, bool, error) {
	userDN, err := repositories.GetUserDN(username)
	if err != nil {
		fmt.Println("error 1")
		return "", false, fmt.Errorf("failed to find user: %v", err)
	}

	ldapConn, err := repositories.ConnectToLDAP()
	if err != nil {
		fmt.Println("error 2")
		return "", false, fmt.Errorf("failed to connect to LDAP: %v", err)
	}
	defer ldapConn.Close()

	if err := ldapConn.Bind(userDN, password); err != nil {
		fmt.Println("error 3")
		return "", false, fmt.Errorf("authentication failed: %v", err)
	}

	role, err := repositories.GetUserRole(username)
	if err != nil {
		fmt.Println("error 5")
		return "", false, fmt.Errorf("failed to get user role: %v", err)
	}

	jwtToken, err := GenerateJWT(username, role)
	if err != nil {
		fmt.Println("error 4")
		return "", false, err
	}

	return jwtToken, true, nil
}

func DecodeJWT(tokenString string) (*Claims, error) {
	jwtKey := []byte(config.C.JWTKey)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
