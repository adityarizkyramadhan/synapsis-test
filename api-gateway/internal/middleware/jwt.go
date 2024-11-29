package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		token := ctx.GetHeader("Authorization")
		if token == "" {
			utils.ResponseError(ctx, http.StatusUnauthorized, errors.New("missing token"))
			ctx.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := verifyJWT(token)
		if err != nil {
			utils.ResponseError(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			utils.ResponseError(ctx, http.StatusUnauthorized, errors.New("invalid token"))
			ctx.Abort()
			return
		}
		ctx.Set("id", mapClaims["id"])
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
