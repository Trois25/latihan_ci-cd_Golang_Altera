package repositories

import (
	"praktikum/config"
	"praktikum/middlewares"
	"praktikum/models"
)

func CheckLogin(email string, password string) (models.User, string, error) {
	var data models.User
	// select * from users where email = `email` and password = `password`
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return models.User{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(data.ID),data.Name)
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

func SelectUsers() ([]models.User, error) {
	//menggunakan db
	var datausers []models.User
	// select * from users;
	tx := config.DB.Order("created_at desc").Find(&datausers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datausers, nil
}
