package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-project/internal/model"
	mocksRepo "go-project/mocks/repository"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type testSignUp struct {
	Input  model.DefaultPayload[model.SignUpRequest]
	Expect any
}

func NewGinContext[T any](method, endpoint string, body T) *gin.Context {

	bodyBytes, _ := json.Marshal(body)

	request, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(bodyBytes))
	request.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	return c
}

func TestSignUp(t *testing.T) {

	// Case 1 -> Success
	t.Run("When success, Should return model and error is nil", func(t *testing.T) {

		test := testSignUp{
			Input: model.DefaultPayload[model.SignUpRequest]{
				Data: model.SignUpRequest{
					Username: "pawapotako.p@gmail.com",
					Password: "12345678",
				},
			},
			Expect: "",
		}

		mockUserRepo := &mocksRepo.UserRepository{}

		rand.Seed(time.Now().UnixNano())
		randomID := rand.Intn(1000000)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(test.Input.Data.Password), bcrypt.DefaultCost)

		mockUserRepo.On("Create", mock.AnythingOfType("model.UserModel")).Return(&model.UserModel{
			Id:       uint(randomID),
			Username: "pawapotako.p@gmail.com",
			Password: string(hashedPassword),
		}, nil)

		service := NewAuthenticationService(mockUserRepo)
		actual, err := service.SignUp(test.Input)

		mockUserRepo.AssertExpectations(t)
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.Data)
		assert.NotEmpty(t, actual.Data)
		assert.NotEmpty(t, actual.Data.AccessToken)
		assert.Nil(t, err)

	})

	// Case 2 -> Error
	t.Run("When error, Should return nil model and error is not nil", func(t *testing.T) {

		test := testSignUp{
			Input: model.DefaultPayload[model.SignUpRequest]{
				Data: model.SignUpRequest{
					Username: "pawapotako.p@gmail.com",
					Password: "12345678",
				},
			},
			Expect: "mock error",
		}

		mockUserRepo := &mocksRepo.UserRepository{}

		mockUserRepo.On("Create", mock.AnythingOfType("model.UserModel")).Return(nil, errors.New("mock error"))

		service := NewAuthenticationService(mockUserRepo)
		actual, err := service.SignUp(test.Input)

		mockUserRepo.AssertExpectations(t)
		assert.Nil(t, actual)
		assert.NotNil(t, err)
	})
}
