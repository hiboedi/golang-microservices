package seeds

import (
	"auth-service/web/models"
	"fmt"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}

func CreateUser(db *gorm.DB, seed models.CreateUser) error {
	err := db.Create(&seed).Error

	return err
}

func All() []Seed {
	return []Seed{
		{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				_ = CreateUser(db, models.CreateUser{
					Email:     "hi.boedi8@gmail.com",
					FirstName: "Budiman",
					LastName:  "none",
					Password:  "asdasd",
					Role:      "admin",
				})
				return nil
			},
		},
	}
}

func Run(db *gorm.DB, seeds []Seed) error {
	for _, seed := range seeds {
		err := seed.Run(db)
		if err != nil {
			return fmt.Errorf("error while running seed %s: %s", seed.Name, err)
		}
	}
	return nil
}
