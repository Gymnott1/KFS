package main

import (
	"log"

	"clientcompany/config"
	"clientcompany/db"
	"clientcompany/models"

	"golang.org/x/crypto/bcrypt"
)

type seedUser struct {
	Name     string
	Email    string
	Password string
	Role     string
}

var users = []seedUser{
	{Name: "Admin User", Email: "admin@clientcompany.com", Password: "Admin@1234", Role: "admin"},
	{Name: "Jane Smith", Email: "jane@clientcompany.com", Password: "User@1234", Role: "user"},
	{Name: "John Doe", Email: "john@clientcompany.com", Password: "User@1234", Role: "user"},
}

func main() {
	config.Load()
	db.Connect()

	for _, u := range users {
		var existing models.User
		if err := db.DB.Where("email = ?", u.Email).First(&existing).Error; err == nil {
			log.Printf("skipping %s — already exists", u.Email)
			continue
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash password for %s: %v", u.Email, err)
		}

		user := models.User{
			Name:     u.Name,
			Email:    u.Email,
			Password: string(hash),
			Role:     u.Role,
		}
		if err := db.DB.Create(&user).Error; err != nil {
			log.Fatalf("failed to create user %s: %v", u.Email, err)
		}
		log.Printf("created [%s] %s <%s>", u.Role, u.Name, u.Email)
	}

	log.Println("seed complete")
}
