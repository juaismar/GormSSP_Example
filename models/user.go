package models

type User struct {
	ID    int    `primary_key" `
	Name  string `form:"name"`
	Email string `form:"email"`
	Age   int    `form:"age"`
}
