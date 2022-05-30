package seed

import (
	"log"

	"github.com/asaoud2022/todo/app/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Ali Saoud",
		Email:    "asaoud@yahoo.com",
		Password: "password123",
	},
	models.User{
		Nickname: "Diana Saoud",
		Email:    "asaoud2@gmail.com",
		Password: "password123",
	},
}

var tasks = []models.Task{
	models.Task{
		Name:      "Task 1",
		Completed: false,
	},
	models.Task{
		Name:      "Task 2",
		Completed: false,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Task{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Task{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	
	err = db.Debug().Model(&models.Task{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		tasks[i].UserID = users[i].ID

		err = db.Debug().Model(&models.Task{}).Create(&tasks[i]).Error
		if err != nil {
			log.Fatalf("cannot seed task table: %v", err)
		}
	}
}
