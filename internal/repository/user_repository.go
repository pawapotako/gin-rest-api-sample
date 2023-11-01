package repository

import (
	"rename-service-name-here/internal/model"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserAll() ([]model.UserModel, error)
	GetUserById(id uint) (model.UserModel, error)
	GetUserByPaging(pageload model.PagingPayload, employeeCode string, userLogin string, userFullname string) (*model.PagingPayload, []*model.UserModel, error)
	CreateUser(entity model.UserModel) (*model.UserModel, error)
	UpdateUser(entity model.UserModel) (*model.UserModel, error)
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db}
}

func (r userRepositoryDB) GetUserAll() ([]model.UserModel, error) {
	entity := []model.UserModel{}
	if tx := r.db.Find(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return entity, nil
}

func (r userRepositoryDB) GetUserById(id uint) (model.UserModel, error) {
	entity := model.UserModel{}
	if tx := r.db.First(&entity, id); tx.Error != nil {
		return entity, tx.Error
	}
	return entity, nil
}

func (r userRepositoryDB) GetUserByPaging(pagination model.PagingPayload, employeeCode string, userLogin string, userFullname string) (*model.PagingPayload, []*model.UserModel, error) {
	entities := []*model.UserModel{}
	dbCondition := r.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return containEmployeeCode(db, employeeCode)
	}, func(db *gorm.DB) *gorm.DB {
		return containUserLogin(db, userLogin)
	}, func(db *gorm.DB) *gorm.DB {
		return containUserFullname(db, userFullname)
	}, func(db *gorm.DB) *gorm.DB {
		return containUserIsActive(db)
	},
	)
	dbCondition.Scopes(paginate(entities, &pagination, dbCondition)).Find(&entities)
	pagination.Data = entities
	return &pagination, entities, nil
}

func (r userRepositoryDB) CreateUser(entity model.UserModel) (*model.UserModel, error) {
	if tx := r.db.Create(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}

func (r userRepositoryDB) UpdateUser(entity model.UserModel) (*model.UserModel, error) {

	originalEntity := model.UserModel{}
	if tx := r.db.First(&originalEntity, entity.Id); tx.Error != nil {
		return nil, tx.Error
	}
	// Update multiple attributes with `struct`, will only update those changed & non blank fields
	// For below Update, nothing will be updated as "", 0, false are blank values of their types
	if tx := r.db.Model(&entity).Updates(entity); tx.Error != nil {
		return nil, tx.Error
	}
	newData := model.UserModel{}
	if tx := r.db.First(&newData, originalEntity.Id); tx.Error != nil {
		return nil, tx.Error
	}
	return &newData, nil
}

func containEmployeeCode(db *gorm.DB, code string) *gorm.DB {
	if code == "" {
		return db
	} else {
		return db.Where("employee_code LIKE ?", "%"+code+"%")
	}
}

func containUserFullname(db *gorm.DB, userFullname string) *gorm.DB {
	if userFullname == "" {
		return db
	} else {
		return db.Where("(name_thai LIKE ? OR surname_thai LIKE ?)", "%"+userFullname+"%", "%"+userFullname+"%")
	}
}

func containUserLogin(db *gorm.DB, userLogin string) *gorm.DB {
	if userLogin == "" {
		return db
	} else {
		return db.Where("user_login = ?", userLogin)
	}
}

func containUserIsActive(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", true)
}
