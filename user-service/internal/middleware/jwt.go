package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		token := ctx.GetHeader("Authorization")
		if token == "" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("token not found").Error())
			ctx.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := verifyJWT(token)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		// Check if the role in the token matches one of the desired roles
		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "invalid token claims")
			ctx.Abort()
			return
		}
		ctx.Set("id", mapClaims["id"].(string))
		ctx.Next()
	}
}

func verifyJWT(tokenString string) (interface{}, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("SECRET_KEY environment variable not set")
	}
	var jwtSecret = []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}
