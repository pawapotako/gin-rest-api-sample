package service

import (
	"go-project/internal/model"
	"go-project/internal/repository"
	"go-project/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	SignUp(request model.DefaultPayload[model.SignUpRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error)
}

type authenticationService struct {
	userRepo     repository.UserRepository
	externalRepo repository.ExternalRepository
}

func NewAuthenticationService(userRepo repository.UserRepository, externalRepo repository.ExternalRepository) AuthenticationService {

	return authenticationService{
		userRepo:     userRepo,
		externalRepo: externalRepo,
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

	response, err := s.userRepo.Create(entity)
	if err != nil {
		return nil, err
	}

	userStreaming := model.UserStreaming{
		UserId:   response.Id,
		Username: response.Username,
	}

	s.externalRepo.EventStreaming(userStreaming)

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
