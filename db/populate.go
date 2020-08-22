package db

import (
	"GormSSP_Example/models"
	_ "GormSSP_Example/routers"
)

func PopulateUsers() {
	users := defaultUsers()
	insert := "INSERT INTO users (id, name, email, age) VALUES (?, ?, ?, ?)"
	for _, user := range users {
		var userInDB models.User
		notFound := models.Conn.Where("email = ?", user.Email).Find(&userInDB).RecordNotFound()
		if notFound {
			models.Conn.Exec(insert, user.ID, user.Name, user.Email, user.Age)
		}
	}
}
func defaultUsers() []models.User {
	return []models.User{
		{
			ID:    1,
			Name:  "Juan",
			Email: "juan@example.com",
			Age:   20,
		},
		{
			ID:    2,
			Name:  "Jose",
			Email: "jose@example.com",
			Age:   18,
		},
		{
			ID:    3,
			Name:  "Javi",
			Email: "javi@example.com",
			Age:   22,
		},
		{
			ID:    4,
			Name:  "Sergio",
			Email: "sergio@example.com",
			Age:   22,
		},
		{
			ID:    5,
			Name:  "Maria",
			Email: "maria@example.com",
			Age:   19,
		},
		{
			ID:    6,
			Name:  "Laura",
			Email: "laura@example.com",
			Age:   18,
		},
		{
			ID:    7,
			Name:  "Marta",
			Email: "marta@example.com",
			Age:   23,
		},
		{
			ID:    8,
			Name:  "Clara",
			Email: "clara@example.com",
			Age:   21,
		},
	}
}
