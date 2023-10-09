package usecase

import (
	"belajar-go-echo/features/user"
	"errors"
)

type userUsecase struct {
	userRepository user.DataInterface
}

// Insert implements user.UseCaseInterface.
func (uc *userUsecase) Insert(data user.UserCore) (row int, err error) {
	if data.Email == "" || data.Password == "" {
		return 0, errors.New("error, email or password can't be empty")
	}
	erruser, _ := uc.userRepository.Insert(data)
	return erruser, nil
}

// Login implements user.UseCaseInterface.
func (uc *userUsecase) Login(email string, password string) (user.UserCore, string, error) {
	if email == "" || password == "" {
		return user.UserCore{}, "error, email or password can't be empty", errors.New("error login")
	}

	// Call the repository function to perform the login
	logindata, token, err := uc.userRepository.Login(email, password)
	if err != nil {
		// Handle the error from the repository
		return user.UserCore{}, "", err
	}

	// Return the login data and a success message
	return logindata, token, nil

}

// GetAllUsers implements user.UseCaseInterface.
func (uc *userUsecase) GetAllUsers() ([]user.UserCore, error) {
	users, err := uc.userRepository.GetAllUsers()
	if err != nil {
		return nil, errors.New("error get data")
	}
	return users, nil
}

func New(Useruc user.DataInterface) user.UseCaseInterface {
	return &userUsecase{
		userRepository: Useruc,
	}
}
