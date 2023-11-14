package mocks

import (
	"go-project/internal/model"

	"github.com/stretchr/testify/mock"
)

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

type UserRepository struct {
	mock.Mock
}

func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (m *UserRepository) Create(a0 model.UserModel) (*model.UserModel, error) {
	ret := m.Called(a0)

	var r0 *model.UserModel
	if rf, ok := ret.Get(0).(func(model.UserModel) *model.UserModel); ok {
		r0 = rf(a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.UserModel) error); ok {
		r1 = rf(a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
