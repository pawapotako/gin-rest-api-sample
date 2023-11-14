package mocks

import (
	"go-project/internal/model"

	"github.com/stretchr/testify/mock"
)

type mockConstructorTestingTNewAuthenticationService interface {
	mock.TestingT
	Cleanup(func())
}

type AuthenticationService struct {
	mock.Mock
}

func NewAuthenticationService(t mockConstructorTestingTNewAuthenticationService) *AuthenticationService {
	mock := &AuthenticationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (m *AuthenticationService) SignUp(request model.DefaultPayload[model.SignUpRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error) {
	ret := m.Called(request)

	var r0 *model.DefaultPayload[model.AccessTokenResponse]
	if rf, ok := ret.Get(0).(func(model.DefaultPayload[model.SignUpRequest]) *model.DefaultPayload[model.AccessTokenResponse]); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DefaultPayload[model.AccessTokenResponse])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.DefaultPayload[model.SignUpRequest]) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
