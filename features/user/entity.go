package user

import "time"

type UserCore struct {
    ID        uint   `json:"id"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    Token     string `json:"token"`
}

type DataInterface interface {
	Insert(data UserCore) (row int, err error)
	GetAllUsers()([]UserCore,error) 
	Login(email, password string) (UserCore, string, error)
}

type UseCaseInterface interface {
	GetAllUsers()([]UserCore,error) 
	Insert(data UserCore) (row int, err error)
	Login(email, password string) (UserCore, string, error)
}
