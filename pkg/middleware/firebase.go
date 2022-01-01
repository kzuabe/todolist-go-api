package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const CONTEXT_TOKEN_KEY = "token"

type FirebaseAuthMiddleware struct {
	Client *auth.Client
}

type FirebaseAuthMiddlewareInterface interface {
	MiddlewareFunc() gin.HandlerFunc
}

func (middleware *FirebaseAuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken, ok := extractTokenFromAuthHeader(c.Request.Header.Get("Authorization"))
		if !ok {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "need token"})
			c.Abort()
			return
		}

		token, err := middleware.Client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}
		c.Set(CONTEXT_TOKEN_KEY, token)
		c.Next()
	}
}

// 参考: https://github.com/go-kit/kit/blob/a073a093d1ee02b920ab78db0fb5600cef24a10e/auth/jwt/transport.go#L78-L85
func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], "bearer") {
		return "", false
	}

	return authHeaderParts[1], true
}
