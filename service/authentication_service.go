package service

import (
	"go-project/model"
	"go-project/repository"
	"go-project/util"
)

type AuthenticationService interface {
	VerifyGoogleIdToken(idToken string) (*model.DefaultPayload[model.AccessTokenResponse], error)
	SignIn(request model.DefaultPayload[model.SignInRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error)
}

type authenticationService struct {
	externalServiceRepo repository.ExternalServiceRepository
}

func NewAuthenticationService(externalServiceRepo repository.ExternalServiceRepository) AuthenticationService {

	return authenticationService{
		externalServiceRepo: externalServiceRepo,
	}
}

func (s authenticationService) VerifyGoogleIdToken(idToken string) (*model.DefaultPayload[model.AccessTokenResponse], error) {

	_, err := s.externalServiceRepo.VerifyGoogleIdToken(idToken)
	if err != nil {
		return nil, err
	}

	jwtToken, err := util.GenerateToken()
	if err != nil {
		return nil, err
	}

	AccessTokenResponse := model.AccessTokenResponse{
		AccessToken: *jwtToken,
	}

	return &model.DefaultPayload[model.AccessTokenResponse]{
		Data: AccessTokenResponse}, nil
}

func (s authenticationService) SignIn(request model.DefaultPayload[model.SignInRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error) {

	// implement login logic here

	jwtToken, err := util.GenerateToken()
	if err != nil {
		return nil, err
	}

	AccessTokenResponse := model.AccessTokenResponse{
		AccessToken: *jwtToken,
	}

	return &model.DefaultPayload[model.AccessTokenResponse]{
		Data: AccessTokenResponse}, nil
}
