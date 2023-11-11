package model

type UserModel struct {
	Id       uint   `gorm:"column:id;primaryKey;autoIncrement:true"`
	Username string `gorm:"column:username;type:varchar(300);not null:true"`
	Password string `gorm:"column:password;type:varchar(300);not null:true"`
}
