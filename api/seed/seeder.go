package seed

import (
	"log"

	"github.com/aliabdullahcp/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "admin",
		Email:    "admin@gmail.com",
		Password: "password",
		UserRole: "Admin",
	},
	models.User{
		Nickname: "Ali Abdullah",
		Email:    "aliabdullah@gmail.com",
		Password: "password",
		UserRole: "User",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
