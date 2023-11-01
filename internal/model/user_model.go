package model

import "time"

type UserModel struct {
	Id             uint       `gorm:"column:id;primaryKey;autoIncrement:true"`
	UserLogin      string     `gorm:"column:user_login;type:varchar(300);not null:true;"`
	EmployeeCode   string     `gorm:"column:employee_code;type:varchar(300);not null;unique"`
	Email          string     `gorm:"column:email;type:varchar(300);not null:true"`
	NameThai       string     `gorm:"column:name_thai;type:varchar(300);not null:true"`
	SurnameThai    string     `gorm:"column:surname_thai;type:varchar(300);not null:true"`
	NameEnglish    string     `gorm:"column:name_english;type:varchar(300)"`
	SurnameEnglish string     `gorm:"column:surname_english;type:varchar(300)"`
	IsActive       bool       `gorm:"column:is_active;not null:true;index:idx_isactive"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	CreatedUserId  uint       `gorm:"column:created_user_id;not null:true"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"`
	UpdatedUserId  *uint      `gorm:"column:updated_user_id"`
}
