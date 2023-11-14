package handler

import (
	"go-project/internal/middleware"
	"go-project/internal/model"
	"go-project/internal/repository"
	"go-project/internal/service"
	"go-project/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type authenticationHandler struct {
	service  service.AuthenticationService
	validate *validator.Validate
}

func InitAuthenticationHandler(db *gorm.DB, e *gin.Engine, validator *validator.Validate) {

	userRepo := repository.NewUserRepositoryDB(db)
	externalRepo := repository.NewExternalRepositoryDB(db)
	service := service.NewAuthenticationService(userRepo, externalRepo)
	handler := authenticationHandler{service, validator}

	v1 := e.Group("/v1")
	v1.POST("/authentications/sign-up", handler.signUp)
	v1.GET("/authentications/sign-out", middleware.AuthorizationMiddleware, handler.signOut)
}

func (h authenticationHandler) signUp(c *gin.Context) {

	request := model.DefaultPayload[model.SignUpRequest]{}
	if err := c.Bind(&request); err != nil {
		util.ErrorHandler(c, http.StatusBadRequest, err)
		return
	}

	if err := h.validate.Struct(request); err != nil {
		util.ErrorHandler(c, http.StatusUnprocessableEntity, err)
		return
	}

	response, err := h.service.SignUp(request)
	if err != nil {
		util.ErrorHandler(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h authenticationHandler) signOut(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "done"})
}
