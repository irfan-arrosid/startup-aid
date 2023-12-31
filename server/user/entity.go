package user

import "time"

type User struct {
	Id             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFilename string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
