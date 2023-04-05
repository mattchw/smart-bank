package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	token_interfaces "github.com/mattchw/smart-bank/internal/token/interfaces"
	"github.com/mattchw/smart-bank/util"
)

const (
	AUTHORIZATION_HEADER_KEY  = "authorization"
	AUTHORIZATION_TYPE_BEARER = "bearer"
	AUTHORIZATION_PAYLOAD_KEY = "authorization_payload"
)

// AuthMiddleware is a middleware that checks if the request has a valid access token.
// If the access token is valid, it will call the next handler.
// Otherwise, it will return an error.
func AuthMiddleware(tokenMaker token_interfaces.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION_HEADER_KEY)
		if authorizationHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(errors.New("authorization header is empty")))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(errors.New("authorization header is invalid")))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AUTHORIZATION_TYPE_BEARER {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(errors.New("authorization type is not bearer")))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		ctx.Set(AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Next()
	}
}
