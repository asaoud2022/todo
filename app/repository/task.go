package repository

import (
	"fmt"

	"github.com/asaoud2022/todo/app/models"
	"github.com/asaoud2022/todo/config"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository interface {
	Save(task models.Task) (uint64, error)
	Update(models.Task) error
	Delete(models.Task) error
	FindAll() []*models.Task
	FindByID(taskID uint64) (*models.Task, error)
	DeleteByID(taskID uint64) error
	FindByName(name string) (*models.Task, error)
	FindByField(fieldName, fieldValue string) (*models.Task, error)
	UpdateSingleField(task models.Task, fieldName, fieldValue string) error
}
type taskDatabase struct {
	connection *gorm.DB
}

func NewTaskRepository() TaskRepository {
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
	return &taskDatabase{
		connection: db,
	}
}

func (db taskDatabase) DeleteByID(taskID uint64) error {
	task := models.Task{}
	task.ID = taskID
	result := db.connection.Delete(&task)
	return result.Error
}

func (db taskDatabase) Save(task models.Task) (uint64, error) {
	result := db.connection.Create(&task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil
}

func (db taskDatabase) Update(task models.Task) error {
	result := db.connection.Save(&task)
	return result.Error
}

func (db taskDatabase) Delete(task models.Task) error {
	result := db.connection.Delete(&task)
	return result.Error
}

func (db taskDatabase) FindAll() []*models.Task {
	var tasks []*models.Task
	db.connection.Preload(clause.Associations).Find(&tasks)
	return tasks
}

func (db taskDatabase) FindByID(taskID uint64) (*models.Task, error) {
	var task models.Task
	result := db.connection.Preload(clause.Associations).Find(&task, "id = ?", taskID)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &task, nil
	}
	return nil, nil
}

func (db taskDatabase) FindByName(name string) (*models.Task, error) {
	var task models.Task
	result := db.connection.Preload(clause.Associations).Find(&task, "name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &task, nil
	}
	return nil, nil
}

func (db taskDatabase) FindByField(fieldName, fieldValue string) (*models.Task, error) {
	var task models.Task
	result := db.connection.Preload(clause.Associations).Find(&task, fmt.Sprintf("%s = ?", fieldName), fieldValue)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &task, nil
	}
	return nil, nil
}

func (db taskDatabase) UpdateSingleField(task models.Task, fieldName, fieldValue string) error {
	result := db.connection.Model(&task).Update(fieldName, fieldValue)
	return result.Error
}
