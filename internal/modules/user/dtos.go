package user

import "time"

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}
type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserPreviewPerson struct {
	Username   string    `json:"username"`
	Bio        string    `json:"bio"`
	Created_At time.Time `json:"created_at"`
}
