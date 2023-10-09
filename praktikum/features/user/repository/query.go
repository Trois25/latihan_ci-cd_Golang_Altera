package repository

import (
	"belajar-go-echo/app/middlewares"
	"belajar-go-echo/features/user"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Insert implements user.DataInterface.
func (userRep *userRepository) Insert(data user.UserCore) (row int, err error) {
	var input = User{
		Email:    data.Email,
		Password: data.Password,
	}

	erruser := userRep.db.Save(&input)
	if erruser.Error != nil {
		return 0, erruser.Error
	}
	return 1, nil
}

// CheckByEmail implements user.DataInterface.
func (userRep *userRepository) Login(email string, password string) (user.UserCore, string, error) {
	var data User

	tx := userRep.db.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return user.UserCore{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(data.ID), data.Email)
		if errToken != nil {
			return user.UserCore{}, "", errToken
		}
	}

	var check = user.UserCore{
		ID:        data.ID,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return check, token, nil
}

// GetAllUsers implements user.DataInterface.
func (userRep *userRepository) GetAllUsers() ([]user.UserCore, error) {
	var users []User

	fmt.Println("sebelum find")
	err := userRep.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("sebelum mapping ")
	data := make([]user.UserCore, len(users))
	for i, value := range users {
		data[i] = user.UserCore{
			ID:        value.ID,
			Email:     value.Email,
			Password:  value.Password,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
	}
	fmt.Println("data :", data)

	return data, nil
}

func New(db *gorm.DB) user.DataInterface {
	return &userRepository{
		db: db,
	}
}
