package repository

import (
	"belajar-go-echo/app/configs"
	"belajar-go-echo/features/user"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DbConn = configs.InitDBTest()

func TestCreateUser(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	DbConn.AutoMigrate(&User{})

	repo := New(DbConn)
	t.Run("Test Create User", func(t *testing.T) {
		mockUser := user.UserCore{Email: "wawan@gmail.com", Password: "123321juan"}
		row, err := repo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
}

func TestCreateUserFail(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	//DbConn.AutoMigrate(&User{})

	repo := New(DbConn)
	t.Run("Test Create User", func(t *testing.T) {
		mockUser := user.UserCore{Email: "wawan@gmail.com", Password: "123321juan"}
		row, err := repo.Insert(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestLoginSuccess(t *testing.T) {
	// DbConn.Migrator().DropTable(&User{})
	DbConn.AutoMigrate(&User{})
	repo := New(DbConn)

	testLoginUser := UserResponse{
		Email:    "wawan@gmail.com",
		Password: "123321",
	}

	DbConn.Create(&testLoginUser)
	data, token, err := repo.Login("wawan@gmail.com", "123321")
	
	assert.Nil(t, err)
	assert.Nil(t, token)
	assert.Equal(t, "wawan@gmail.com", data.Email)
	assert.Equal(t, "123321", data.Password)
}

func TestLoginFail(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	DbConn.AutoMigrate(&User{})

	repo := New(DbConn)
	testLoginUser := user.UserCore{
		Email:    "wawan@gmail.com",
		Password: "123321",
	}
	DbConn.Create(&testLoginUser)

	data, token, err := repo.Login("wawan@gmail.com", "123321")
	fmt.Println(data)

	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "", data.Email)
	assert.Equal(t, "", data.Password)
}

func TestGetAllUsers(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	DbConn.AutoMigrate(&User{})

	repo := New(DbConn)

	mockUser := user.UserCore{Email: "wawan@gmail.com", Password: "123321juan"}
	row, _ := repo.Insert(mockUser)

	data, err := repo.GetAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, 1, row)
	assert.NotEmpty(t, data)
}

func TestGetAllUsersFail(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})

	repo := New(DbConn)

	data, err := repo.GetAllUsers()

	assert.NotNil(t, err)
	assert.Empty(t, data)
}
