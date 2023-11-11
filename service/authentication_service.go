package service

import (
	"go-project/model"
	"go-project/repository"
	"go-project/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	SignUp(request model.DefaultPayload[model.SignUpRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error)
}

type authenticationService struct {
	externalServiceRepo repository.ExternalServiceRepository
	userRepo            repository.UserRepository
}

func NewAuthenticationService(externalServiceRepo repository.ExternalServiceRepository, userRepo repository.UserRepository) AuthenticationService {

	return authenticationService{
		externalServiceRepo: externalServiceRepo,
		userRepo:            userRepo,
	}
}

func (s authenticationService) SignUp(request model.DefaultPayload[model.SignUpRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error) {

	data := request.Data

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	entity := model.UserModel{
		Username: data.Username,
		Password: string(hashedPassword),
	}

	response, err := s.userRepo.Insert(entity)
	if err != nil {
		return nil, err
	}

	jwtToken, err := util.GenerateToken(response.Id)
	if err != nil {
		return nil, err
	}

	AccessTokenResponse := model.AccessTokenResponse{
		AccessToken: *jwtToken,
	}

	return &model.DefaultPayload[model.AccessTokenResponse]{
		Data: AccessTokenResponse}, nil
}
