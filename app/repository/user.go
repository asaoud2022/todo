package repository

import (
	"fmt"

	"github.com/asaoud2022/todo/app/models"
	"github.com/asaoud2022/todo/config"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Save(user models.User) (uint32, error)
	Update(models.User) error
	Delete(models.User) error
	FindAll() []*models.User
	FindByID(userID uint32) (*models.User, error)
	DeleteByID(userID uint32) error
	FindByName(name string) (*models.User, error)
	FindByField(fieldName, fieldValue string) (*models.User, error)
	UpdateSingleField(user models.User, fieldName, fieldValue string) error
}
type userDatabase struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	if db == nil {
		c, err := config.LoadConfig()

		if err != nil {
			log.Fatalln("Failed at config", err)
		}

		_, err = Connect(&c)
		if err != nil {
			log.Error(err)
		}
	}
	return &userDatabase{
		connection: db,
	}
}

func (db userDatabase) DeleteByID(userID uint32) error {
	user := models.User{}
	user.ID = userID
	result := db.connection.Delete(&user)
	return result.Error
}

func (db userDatabase) Save(user models.User) (uint32, error) {
	result := db.connection.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (db userDatabase) Update(user models.User) error {
	result := db.connection.Save(&user)
	return result.Error
}

func (db userDatabase) Delete(user models.User) error {
	result := db.connection.Delete(&user)
	return result.Error
}

func (db userDatabase) FindAll() []*models.User {
	var users []*models.User
	db.connection.Preload(clause.Associations).Find(&users)
	return users
}

func (db userDatabase) FindByID(userID uint32) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, "id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) FindByName(name string) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, "name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) FindByField(fieldName, fieldValue string) (*models.User, error) {
	var user models.User
	result := db.connection.Preload(clause.Associations).Find(&user, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}

func (db userDatabase) UpdateSingleField(user models.User, fieldName, fieldValue string) error {
	result := db.connection.Model(&user).Update(fieldName, fieldValue)
	return result.Error
}
