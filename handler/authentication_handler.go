package handler

import (
	"go-project/middleware"
	"go-project/model"
	"go-project/repository"
	"go-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authenticationHandler struct {
	service service.AuthenticationService
}

func InitAuthenticationHandler(db *gorm.DB, e *gin.Engine) {

	externalServiceRepo := repository.NewExternalServiceRepositoryDB(db)
	service := service.NewAuthenticationService(externalServiceRepo)
	handler := authenticationHandler{service}

	v1 := e.Group("/v1")
	v1.GET("/authentications/verify-google-id-token", handler.verifyGoogleIdToken)
	v1.POST("/authentications/sign-in", handler.signIn)
	v1.GET("/authentications/sign-out", middleware.AuthorizationMiddleware, handler.signOut)
}

func (h authenticationHandler) verifyGoogleIdToken(c *gin.Context) {

	idToken := c.Query("idToken")

	response, err := h.service.VerifyGoogleIdToken(idToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, response)
}

func (h authenticationHandler) signIn(c *gin.Context) {

	request := model.DefaultPayload[model.SignInRequest]{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	response, err := h.service.SignIn(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, response)
}

func (h authenticationHandler) signOut(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "done"})
}
