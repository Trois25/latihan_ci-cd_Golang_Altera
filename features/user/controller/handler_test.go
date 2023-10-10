package controller

import (
	"bytes"
	// "belajar-go-echo/features/user/entity"
	// "belajar-go-echo/ap/middlewares"
	"belajar-go-echo/features/user"
	"belajar-go-echo/features/user/repository"
	"encoding/json"
	"errors"

	mocks "belajar-go-echo/app/mocks"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsersCreateUser(t *testing.T) {
	// Membuat mock UseCaseInterface
	useCase := new(mocks.UseCaseInterface)

	// Membuat instance UserController dengan mock UseCaseInterface
	userController := New(useCase)

	// Data yang akan digunakan sebagai input dalam request
	mockData := repository.User{
		Email:    "wawansutiardji@gmail.com",
		Password: "123321wawan",
	}

	// Marshal data mock menjadi JSON
	requestJSON, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err, "error")
	}

	// Membuat instance Echo dan request HTTP
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Ekspektasi bahwa metode Insert akan dipanggil dengan argumen yang sesuai
	useCase.On("Insert", mock.Anything).Return(1, nil).Once()

	// Menjalankan CreateUser
	err = userController.CreateUser(c)

	// Memeriksa bahwa tidak ada kesalahan yang terjadi
	assert.NoError(t, err)

	// Memeriksa bahwa status kode HTTP adalah 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Memeriksa bahwa respons JSON sesuai dengan ekspektasi
	expectedResponse := `{"message":"success insert data","data":1}`
	assert.JSONEq(t, expectedResponse, rec.Body.String())

	// Verifikasi bahwa metode Insert dipanggil dengan argumen yang benar
	useCase.AssertExpectations(t)
}
func TestUsersCreateUserFail(t *testing.T) {
	// Membuat mock UseCaseInterface
	useCase := new(mocks.UseCaseInterface)

	// Membuat instance UserController dengan mock UseCaseInterface
	userController := New(useCase)

	// Data yang akan digunakan sebagai input dalam request
	mockData := repository.User{
		Email:    "wawansutiardji@gmail.com",
		Password: "123321wawan",
	}

	// Marshal data mock menjadi JSON
	requestJSON, err := json.Marshal(mockData)
	if err != nil {
		t.Error(err, "error")
	}

	// Membuat instance Echo dan request HTTP
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Ekspektasi bahwa metode Insert akan dipanggil dengan argumen yang sesuai
	// dan mengembalikan kesalahan
	expectedError := errors.New("error insert data")
	useCase.On("Insert", mock.Anything).Return(0, expectedError).Once()

	// Menjalankan CreateUser
	userController.CreateUser(c)

	// Memeriksa bahwa status kode HTTP adalah 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Tidak perlu memeriksa isi respons JSON karena dalam situasi ini kita
	// hanya ingin memeriksa bahwa kesalahan terjadi saat mencoba
	// melakukan inser data.

	// Verifikasi bahwa metode Insert dipanggil dengan argumen yang benar
	useCase.AssertExpectations(t)
}

func TestUserLogin(t *testing.T) {
	// Membuat mock UseCaseInterface
	useCase := new(mocks.UseCaseInterface)

	// Membuat instance UserController dengan mock UseCaseInterface
	userController := New(useCase)

	// Data yang akan digunakan sebagai input dalam request
	mockRequest := UserRequest{
		Email:    "wawansutiardji@gmail.com",
		Password: "123321wawan",
	}

	// Marshal data mock menjadi JSON
	requestJSON, err := json.Marshal(mockRequest)
	if err != nil {
		t.Error(err, "error")
	}

	// Membuat instance Echo dan request HTTP
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Data yang akan digunakan sebagai respons dari UseCaseInterface
	mockResponse := user.UserCore{
		Email: "wawansutiardji@gmail.com",
		Token: "example_token",
	}

	// Ekspektasi bahwa metode Login akan dipanggil dengan argumen yang sesuai
	useCase.On("Login", mockRequest.Email, mockRequest.Password).Return(mockResponse, "example_token", nil).Once()

	// Menjalankan Login
	err = userController.Login(c)

	// Memeriksa bahwa tidak ada kesalahan yang terjadi
	assert.NoError(t, err)

	// Memeriksa bahwa status kode HTTP adalah 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Memeriksa bahwa respons JSON sesuai dengan ekspektasi
	expectedResponse := `{"message":"login success","email":"wawansutiardji@gmail.com","token":"example_token"}`
	assert.JSONEq(t, expectedResponse, rec.Body.String())

	// Verifikasi bahwa metode Login dipanggil dengan argumen yang benar
	useCase.AssertExpectations(t)
}

func TestUserLoginFail(t *testing.T) {
	// Membuat mock UseCaseInterface
	useCase := new(mocks.UseCaseInterface)

	// Membuat instance UserController dengan mock UseCaseInterface
	userController := New(useCase)

	// Data yang akan digunakan sebagai input dalam request
	mockRequest := UserRequest{
		Email:    "wawansutiardji@gmail.com",
		Password: "123321wawan",
	}

	// Marshal data mock menjadi JSON
	requestJSON, err := json.Marshal(mockRequest)
	if err != nil {
		t.Error(err, "error")
	}

	// Membuat instance Echo dan request HTTP
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Ekspektasi bahwa metode Login akan dipanggil dengan argumen yang sesuai
	// dan mengembalikan kesalahan
	expectedError := errors.New("error login")
	useCase.On("Login", mockRequest.Email, mockRequest.Password).Return(user.UserCore{}, "", expectedError).Once()

	// Menjalankan Login
	userController.Login(c)

	// Memeriksa bahwa status kode HTTP adalah 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Tidak perlu memeriksa isi respons JSON karena dalam situasi ini kita
	// hanya ingin memeriksa bahwa kesalahan terjadi saat mencoba
	// melakukan login.

	// Verifikasi bahwa metode Login dipanggil dengan argumen yang benar
	useCase.AssertExpectations(t)
}

func TestGetAllUsersHandler(t *testing.T) {
	// Membuat mock UseCaseInterface
	useCase := new(mocks.UseCaseInterface)

	// Membuat instance UserController dengan mock UseCaseInterface
	userController := New(useCase)

	// Data yang akan digunakan sebagai respons dari UseCaseInterface
	mockResponse := []user.UserCore{
		{ID: 1, Email: "user1@example.com"},
		{ID: 2, Email: "user2@example.com"},
	}

	// Ekspektasi bahwa metode GetAllUsers akan dipanggil dan mengembalikan hasil yang sesuai
	useCase.On("GetAllUsers").Return(mockResponse, nil).Once()

	// Membuat instance Echo dan request HTTP
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Menjalankan GetAllUsers
	err := userController.GetAllUsers(c)

	// Memeriksa bahwa tidak ada kesalahan yang terjadi
	assert.NoError(t, err)

	// Memeriksa bahwa status kode HTTP adalah 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Mengubah respons sesuai dengan struktur UserCore
	expectedResponse := `{"message":"get all data","data":[
        {"id":1,"email":"user1@example.com","password":"","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","token":""},
        {"id":2,"email":"user2@example.com","password":"","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","token":""}
    ]}`

	assert.JSONEq(t, expectedResponse, rec.Body.String())

	// Verifikasi bahwa metode GetAllUsers dipanggil
	useCase.AssertExpectations(t)
}

func TestGetAllUsersHandlerFail(t *testing.T) {
	// Membuat mock UseCaseInterface
	useCase := new(mocks.UseCaseInterface)

	// Membuat instance UserController dengan mock UseCaseInterface
	userController := New(useCase)

	// Ekspektasi bahwa metode GetAllUsers akan dipanggil dan mengembalikan kesalahan
	expectedError := errors.New("error get all data")
	useCase.On("GetAllUsers").Return(nil, expectedError).Once()

	// Membuat instance Echo dan request HTTP
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Menjalankan GetAllUsers
	userController.GetAllUsers(c)

	// Memeriksa bahwa status kode HTTP adalah 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Tidak perlu memeriksa isi respons JSON karena dalam situasi ini kita
	// hanya ingin memeriksa bahwa kesalahan terjadi saat mencoba
	// melakukan pengambilan data.

	// Verifikasi bahwa metode GetAllUsers dipanggil
	useCase.AssertExpectations(t)
}
