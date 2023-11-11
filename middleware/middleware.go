package middleware

import (
	"go-project/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {

	authorization := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(authorization, "Bearer ")

	if err := util.ValidateToken(token); err != nil {
		util.ErrorHandler(c, http.StatusUnauthorized, err)
		return
	}
}
