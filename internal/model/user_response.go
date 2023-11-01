package model

import "time"

// GetUserResponse godoc
// @Summary Get User response structure
// @Description Detailed description for get User response
// @Model GetUserResponse
type GetUserResponse struct {
	Id             uint       `json:"id"`
	UserLogin      string     `json:"userLogin"`
	EmployeeCode   string     `json:"employeeCode"`
	Email          string     `json:"email"`
	NameThai       string     `json:"nameThai" `
	SurnameThai    string     `json:"surnameThai"`
	NameEnglish    string     `json:"nameEnglish" `
	SurnameEnglish string     `json:"surnameEnglish" `
	IsActive       bool       `json:"isActive"`
	CreatedAt      time.Time  `json:"createdDate"`
	CreatedUserId  uint       `json:"createdUserId"`
	UpdatedAt      *time.Time `json:"updatedDate"`
	UpdateUserId   *uint      `json:"updatedUserId"`
}

// NewUserResponse godoc
// @Summary New User response structure
// @Description Detailed description for new User response
// @Model NewUserResponse
type NewUserResponse struct {
	Id uint `json:"id"`
}
