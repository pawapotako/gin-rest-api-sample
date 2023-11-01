package service

import (
	"rename-service-name-here/internal/model"
	"rename-service-name-here/internal/repository"
	"rename-service-name-here/internal/util"
	"time"

	"go.uber.org/zap"
)

type UserService interface {
	GetUserById(id uint) (*model.DefaultPayload[model.GetUserResponse], error)
	GetUserByPaging(page int, limit int, sort string, employeeCode string, userLogin string, userFullname string) (*model.PagingPayload, error)
	NewUser(request model.DefaultPayload[model.NewUserRequest]) (*model.DefaultPayload[model.NewUserResponse], error)
	UpdateUser(request model.DefaultPayload[model.UpdateUserRequest]) (*model.DefaultPayload[model.GetUserResponse], error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{
		userRepo: userRepo,
	}
}

func (s userService) GetUserById(id uint) (*model.DefaultPayload[model.GetUserResponse], error) {
	entity, err := s.userRepo.GetUserById(id)
	if err != nil {
		util.Logger.Error("error", zap.Error(err))
		return nil, err
	}

	responses := model.GetUserResponse{
		Id:             entity.Id,
		UserLogin:      entity.UserLogin,
		EmployeeCode:   entity.EmployeeCode,
		Email:          entity.Email,
		NameThai:       entity.NameThai,
		SurnameThai:    entity.SurnameThai,
		NameEnglish:    entity.NameEnglish,
		SurnameEnglish: entity.SurnameEnglish,
		IsActive:       entity.IsActive,
		CreatedAt:      entity.CreatedAt,
		CreatedUserId:  entity.CreatedUserId,
		UpdatedAt:      entity.UpdatedAt,
		UpdateUserId:   entity.UpdatedUserId,
	}

	return &model.DefaultPayload[model.GetUserResponse]{
		Data: responses}, nil
}

func (s userService) GetUserByPaging(page int, limit int, sort string, employeeCode string, userLogin string, userFullname string) (*model.PagingPayload, error) {
	pagination := model.PagingPayload{
		Paging: model.Paging{
			Page:  page,
			Limit: limit,
			Sort:  sort,
		},
	}

	results, data, err := s.userRepo.GetUserByPaging(pagination, employeeCode, userLogin, userFullname)
	if err != nil {
		return nil, err
	}
	results.Data = data
	return results, nil
}

func (s userService) NewUser(request model.DefaultPayload[model.NewUserRequest]) (*model.DefaultPayload[model.NewUserResponse], error) {

	location, _ := time.LoadLocation("Asia/Bangkok")
	var nowTime = time.Now().In(location)

	user := model.UserModel{
		UserLogin:      request.Data.UserLogin,
		EmployeeCode:   request.Data.EmployeeCode,
		Email:          request.Data.Email,
		NameThai:       request.Data.NameThai,
		SurnameThai:    request.Data.SurnameThai,
		NameEnglish:    request.Data.NameEnglish,
		SurnameEnglish: request.Data.SurnameEnglish,
		IsActive:       true,
		CreatedAt:      nowTime,
		CreatedUserId:  request.Data.CreatedUserId,
	}

	newUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		util.Logger.Error("error", zap.Error(err))
		return nil, err
	}

	response := model.NewUserResponse{
		Id: newUser.Id,
	}
	return &model.DefaultPayload[model.NewUserResponse]{
		Data: response}, nil
}

func (s userService) UpdateUser(request model.DefaultPayload[model.UpdateUserRequest]) (*model.DefaultPayload[model.GetUserResponse], error) {

	var data = request.Data
	location, _ := time.LoadLocation("Asia/Bangkok")
	var nowTime = time.Now().In(location)

	updateData := model.UserModel{
		Id:             data.Id,
		UserLogin:      data.UserLogin,
		EmployeeCode:   data.EmployeeCode,
		Email:          data.Email,
		NameThai:       data.NameThai,
		SurnameThai:    data.SurnameThai,
		NameEnglish:    data.NameEnglish,
		SurnameEnglish: data.SurnameEnglish,
		IsActive:       data.IsActive,
		UpdatedAt:      &nowTime,
		UpdatedUserId:  &data.UpdatedUserId,
	}
	entity, err := s.userRepo.UpdateUser(updateData)
	if err != nil {
		util.Logger.Error("error", zap.Error(err))
		return nil, err
	}

	responses := model.GetUserResponse{
		Id:             entity.Id,
		UserLogin:      entity.UserLogin,
		EmployeeCode:   entity.EmployeeCode,
		Email:          entity.Email,
		NameThai:       entity.NameThai,
		SurnameThai:    entity.SurnameThai,
		NameEnglish:    entity.NameEnglish,
		SurnameEnglish: entity.SurnameEnglish,
		IsActive:       entity.IsActive,
		CreatedAt:      entity.CreatedAt,
		CreatedUserId:  entity.CreatedUserId,
		UpdatedAt:      entity.UpdatedAt,
		UpdateUserId:   entity.UpdatedUserId,
	}

	return &model.DefaultPayload[model.GetUserResponse]{
		Data: responses}, nil
}
