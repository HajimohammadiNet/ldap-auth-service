package controllers

import (
	"github.com/gin-gonic/gin"
	"ldap-auth/services"
	"net/http"
	"strings"
)

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	jwttoken, authenticated, err := services.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil || !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials or unauthorized access"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authenticated successfully", "token": jwttoken, "username": credentials.Username})
}

func ValidateTokenHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	// Extract token from Authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Token validation
	claims, err := services.DecodeJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": claims.Username,
		"role":     claims.Role,
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Statue": "Ok"})
}
