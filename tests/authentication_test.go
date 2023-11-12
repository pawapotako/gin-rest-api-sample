package tests

import (
	"database/sql"
	"fmt"
	"go-project/model"
	"go-project/repository"
	"go-project/service"
	"go-project/util"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type fakeAuthenticationController struct {
	service service.AuthenticationService
}

func InitTestAuthenticationHandler() *fakeAuthenticationController {

	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("cannot set timezone", err)
	}

	time.Local = ict

	config := util.Config{}

	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.test")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Unable to decode into struct", err)
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Asia%%2FBangkok",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("cannot connect to db ", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot open db ", err)
	}

	externalServiceRepo := repository.NewExternalServiceRepositoryDB(db)
	userRepo := repository.NewUserRepositoryDB(db)
	service := service.NewAuthenticationService(externalServiceRepo, userRepo)

	return &fakeAuthenticationController{
		service: service,
	}
}

func (f *fakeAuthenticationController) FakeSignUp(request model.DefaultPayload[model.SignUpRequest]) (*model.DefaultPayload[model.AccessTokenResponse], error) {
	response, err := f.service.SignUp(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type testRegister struct {
	Label  string
	Input  model.DefaultPayload[model.SignUpRequest]
	Expect any
	Type   string
}

func TestRegister(t *testing.T) {
	controller := InitTestAuthenticationHandler()
	tests := []testRegister{
		// Case 1 -> No error
		{
			Label: "no error",
			Input: model.DefaultPayload[model.SignUpRequest]{
				Data: model.SignUpRequest{
					Username: "pawapotako.p@gmail.com",
					Password: "12345678",
				},
			},
			Expect: "no error",
			Type:   "",
		},
		// Case 2 -> Username must be unique
		{
			Label: "unique username",
			Input: model.DefaultPayload[model.SignUpRequest]{
				Data: model.SignUpRequest{
					Username: "pawapotako.p@gmail.com",
					Password: "12345678",
				},
			},
			Expect: "user with username pawapotako.p@gmail.com already exists",
			Type:   "error",
		},
		// Case 3 -> Response must not be nil
		{
			Label: "response must not be nil",
			Input: model.DefaultPayload[model.SignUpRequest]{
				Data: model.SignUpRequest{
					Username: "pawapotako.p@gmail.com",
					Password: "12345678",
				},
			},
			Expect: "",
			Type:   "result",
		},
	}

	for i := range tests {
		switch tests[i].Type {
		case "no error":
			if _, err := controller.FakeSignUp(tests[i].Input); err != nil {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, err.Error())
			}
		case "error":
			if _, err := controller.FakeSignUp(tests[i].Input); err.Error() != tests[i].Expect.(string) {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, err.Error())
			}
		case "result":
			result, err := controller.FakeSignUp(tests[i].Input)
			if err != nil {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, err.Error())
			} else if result == nil {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, "<nil>")
			}
		}
	}
}
