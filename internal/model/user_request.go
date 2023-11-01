package model

// NewUserRequest godoc
// @Summary New User request structure
// @Description Detailed description for new User request
// @Model NewUserRequest
type NewUserRequest struct {
	UserLogin      string `json:"userLogin"  validate:"required"`
	EmployeeCode   string `json:"employeeCode" validate:"required"`
	Email          string `json:"email" validate:"required"`
	NameThai       string `json:"nameThai" validate:"required"`
	SurnameThai    string `json:"surnameThai"  validate:"required"`
	NameEnglish    string `json:"nameEnglish" `
	SurnameEnglish string `json:"surnameEnglish" `
	IsActive       bool   `json:"isActive"`
	CreatedUserId  uint   `json:"createdUserId"`
}

// UpdateUserRequest godoc
// @Summary Update User request structure
// @Description Detailed description for update User request
// @Model UpdateUserRequest
type UpdateUserRequest struct {
	Id             uint   `json:"id" validate:"required"`
	UserLogin      string `json:"userLogin"  validate:"required"`
	EmployeeCode   string `json:"employeeCode" validate:"required"`
	Email          string `json:"email" validate:"required"`
	NameThai       string `json:"nameThai" validate:"required"`
	SurnameThai    string `json:"surnameThai"  validate:"required"`
	NameEnglish    string `json:"nameEnglish" `
	SurnameEnglish string `json:"surnameEnglish" `
	IsActive       bool   `json:"isActive"`
	UpdatedUserId  uint   `json:"updatedUserId"`
}
