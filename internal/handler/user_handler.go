package handler

import (
	"net/http"
	"rename-service-name-here/internal/model"
	"rename-service-name-here/internal/repository"
	"rename-service-name-here/internal/service"
	"rename-service-name-here/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type userHandler struct {
	service  service.UserService
	validate *validator.Validate
}

func InitUserHandler(db *gorm.DB, e *echo.Echo, validator *validator.Validate) {
	repo := repository.NewUserRepositoryDB(db)
	service := service.NewUserService(repo)
	handler := userHandler{service, validator}

	v1 := e.Group("/v1")
	v1.GET("/users", handler.getUsers)
	v1.GET("/users/:id", handler.getUserById)
	v1.POST("/users", handler.newUser)
	v1.PATCH("/users", handler.updateUser)
}

// Get User by Paging godoc
// @Summary Get user by paging with a Get request
// @Description Sends a GET request to get user by paging
// @Tags User
// @Accept  json
// @Produce  json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Param sort query string false "sort"
// @Param employeeCode query string false "employeeCode"
// @Param userFullname query string false "userFullname"
// @Param userLogin query string false "userLogin"
// @Success 200 {object} model.DefaultPayload[model.GetUserResponse]
// @Failure 400  {object}  util.AppErrors
// @Failure 401  {object}  util.AppErrors
// @Failure 500  {object}  util.AppErrors
// @Router /v1/users [get]
func (h userHandler) getUsers(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	sort := ctx.QueryParam("sort")
	employeeCode := ctx.QueryParam("employeeCode")
	userFullname := ctx.QueryParam("userFullname")
	userLogin := ctx.QueryParam("userLogin")
	response, err := h.service.GetUserByPaging(page, limit, sort, employeeCode, userLogin, userFullname)
	if err != nil {
		errs := util.AppErrors{}
		errs.Add(util.NewUnexpectedError())

		return errorHandler(ctx, errs)
	}

	return ctx.JSON(http.StatusOK, response)
}

// Get User godoc
// @Summary Get user with a Get request
// @Description Sends a GET request to get user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.DefaultPayload[model.GetUserResponse]
// @Failure 400  {object}  util.AppErrors
// @Failure 401  {object}  util.AppErrors
// @Failure 500  {object}  util.AppErrors
// @Router /v1/users/{id} [get]
func (h userHandler) getUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := h.service.GetUserById(uint(id))
	if err != nil {
		errs := util.AppErrors{}
		errs.Add(util.NewNotFoundError("Record not found"))
		return errorHandler(c, errs)
	}

	return c.JSON(http.StatusOK, response)
}

// Create User godoc
// @Summary Create User with a POST request
// @Description Sends a POST request to create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param userRequest body model.DefaultPayload[model.NewUserRequest] true "User Request"
// @Success 200 {object} model.DefaultPayload[model.NewUserResponse]
// @Failure 400  {object}  util.AppErrors
// @Failure 422  {object}  util.AppErrors
// @Failure 500  {object}  util.AppErrors
// @Router /v1/users [post]
func (h userHandler) newUser(ctx echo.Context) error {
	errs := util.AppErrors{}
	request := model.DefaultPayload[model.NewUserRequest]{}
	if err := ctx.Bind(&request); err != nil {
		errs.Add(util.NewBadRequestError())
		return errorHandler(ctx, errs)
	}

	if err := h.validate.Struct(request); err != nil {
		errs.Add(util.NewValidationError(err.Error()))
		return errorHandler(ctx, errs)
	}

	response, err := h.service.NewUser(request)
	if err != nil {
		errs.Add(util.NewCustomError(err.Error()))
		return errorHandler(ctx, errs)
	}
	return ctx.JSON(http.StatusOK, response)
}

// Update User godoc
// @Summary Update User with a PATCH request
// @Description Sends a PATCH request to Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param userRequest body model.DefaultPayload[model.NewUserRequest] true "User Request"
// @Success 200 {object} model.DefaultPayload[model.GetUserResponse]
// @Failure 400  {object}  util.AppErrors
// @Failure 422  {object}  util.AppErrors
// @Failure 500  {object}  util.AppErrors
// @Router /v1/users [patch]
func (h userHandler) updateUser(ctx echo.Context) error {
	errs := util.AppErrors{}
	request := model.DefaultPayload[model.UpdateUserRequest]{}
	if err := ctx.Bind(&request); err != nil {
		errs.Add(util.NewBadRequestError())
		return errorHandler(ctx, errs)
	}

	if err := h.validate.Struct(request); err != nil {
		errs.Add(util.NewValidationError(err.Error()))
		return errorHandler(ctx, errs)
	}

	response, err := h.service.UpdateUser(request)
	if err != nil {
		errs.Add(util.NewCustomError(err.Error()))
		return errorHandler(ctx, errs)
	}
	return ctx.JSON(http.StatusOK, response)
}
