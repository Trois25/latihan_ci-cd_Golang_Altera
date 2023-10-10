package usecase

import (
	"belajar-go-echo/app/mocks"
	"belajar-go-echo/features/user"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

// func TestGetAllData(t *testing.T) {
// 	repoData := new(mocks.DataInterface)
// 	returnData := []user.UserCore{{ID: 1, Email: "wawan@gmail.com", Password: "123wawan"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repoData.On("Insert", mock.Anything).Return(returnData, nil).Once()

// 		srv := New(repoData)

// 		data, err := srv.GetAllUsers()
// 		assert.NoError(t, err)

// 		assert.Nil(t, err)
// 		assert.NotEmpty(t,data)
// 		repoData.AssertExpectations(t)
// 	})
// }

func TestCreateUser(t *testing.T) {
	repoData := new(mocks.DataInterface)
	userUC := New(repoData)

	mockUser := user.UserCore{Email: "wawan@gmail.com", Password: "123321juan"}

	repoData.On("Insert", mockUser).Return(1, nil)
	row,err := userUC.Insert(mockUser)

    assert.Nil(t,err)
    assert.Equal(t,1,row)
}

func TestCreateUserFail(t *testing.T) {
	repoData := new(mocks.DataInterface)
	userUC := New(repoData)

	mockUser := user.UserCore{Email: "", Password: ""}

	repoData.On("Insert", mockUser).Return(1, nil)
	row,err := userUC.Insert(mockUser)

    assert.NotNil(t,err)
    assert.Equal(t,0,row)
}

func TestCheckLoginEmptyEmailAndPassword(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := New(mockDataRepo)

	userData,token, err := uc.Login("", "")
	assert.EqualError(t, err, "error login")

	assert.Equal(t, user.UserCore{}, userData)
	assert.Equal(t, "error, email or password can't be empty", token)
}

func TestCheckLoginUserNotFound(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := New(mockDataRepo)

	mockDataRepo.On("Login", "wawan@gmail.com", "123321").Return(user.UserCore{}, "", errors.New("record not found"))
	userData, token, err := uc.Login("wawan@gmail.com", "123321")

	assert.EqualError(t, err, "record not found")

	assert.Equal(t, user.UserCore{}, userData)
	assert.Equal(t, "", token)
}

func TestCheckLoginSuccess(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := New(mockDataRepo)

	checkUserData := user.UserCore{
		Email: "wawan@gmail.com",
	}
	tokenGiven := "token123"

	mockDataRepo.On("Login", "wawan@gmail.com", "123321").Return(checkUserData, tokenGiven, nil)
	userData, token, err := uc.Login("wawan@gmail.com", "123321")

	assert.Nil(t, err)
	assert.Equal(t, checkUserData, userData)
	assert.Equal(t, tokenGiven, token)
}

func TestFindAllUsersSuccess(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := New(mockDataRepo)

	expectedUsers := []user.UserCore{
		{Email: "hilal@gmail.com"},
		{Email: "hilal2@gmail.com"},
	}

	mockDataRepo.On("GetAllUsers").Return(expectedUsers, nil)
	users, err := uc.GetAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, users)
}
