package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Step 1: get JWT from headers
		token := getJWT(ctx)
		// Step 2: Parse JWT
		parsedJWT, err := parseToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// Step 3: valid session
		if !isValidToken(parsedJWT) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Step 4: Get claims and set them to the context.
		setClaims(ctx, parsedJWT)
	}
}

func setClaims(ctx *gin.Context, parsedJWT *jwt.Token) {
	claims := parsedJWT.Claims.(jwt.MapClaims)
	ctx.Set("user", claims["user"])
	ctx.Set("userID", claims["userId"])
}

func isValidToken(jwt *jwt.Token) bool {
	return jwt.Valid
}

func parseToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("superdupersecurepass"), nil
	})
	if err != nil {
		return nil, errors.New("Failed to unmarshal jwt")
	}
	return jwtToken, nil
}

func getJWT(ctx *gin.Context) string {
	authHeader := ctx.Request.Header.Get("Authorization")
	// Bearer BHJBJHBS512313 for example
	return strings.TrimPrefix(authHeader, "Bearer ")
}
