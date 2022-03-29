package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const AuthTokenKey = "token"

type Client *auth.Client

func NewClient() (Client, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewAuthorizer(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken, ok := extractTokenFromAuthHeader(c.Request.Header.Get("Authorization"))
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(AuthTokenKey, token)
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
