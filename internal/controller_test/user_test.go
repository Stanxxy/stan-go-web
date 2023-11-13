package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Stanxxy/stan-go-web/internal/models"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type UserFakeStore struct{}

func (s *UserFakeStore) First(m *models.User) error {
	return nil
}
func (s *UserFakeStore) Find(m *[]models.User) error {
	return nil
}
func (s *UserFakeStore) Create(m *models.User) error {
	return nil
}
func (s *UserFakeStore) Ping() error {
	return nil
}

func TestSignUp(t *testing.T) {
	// Given
	body := map[string]any{
		"account":   "dummy_account",
		"password":  "dummy_password",
		"email":     "stan.x.liu@gmail.com", // save the real email for blocking test cases
		"phoenNum":  "19939283432",
		"checkCode": "dummy_checkcode",
		"zipcode":   03002,
	}
	jsonData, _ := json.Marshal(body) // ignore the error

	// When
	req := httptest.NewRequest(echo.POST, "/users/signUp", bytes.NewReader(jsonData))
	rec := httptest.NewRecorder()
	e.server.Echo.ServeHTTP(rec, req)

	// Then
	assert.Equal(t, http.StatusOK, rec.Code)

	// Search in the database to check if the write succeeded or not
	user := models.User{
		Username: body["account"].(string),
		Password: body["password"].(string),
		Email:    body["email"].(string),
		PhoneNum: body["phoneNum"].(string),
		Zipcode:  body["zipcode"].(string),
	}
	rowsAffected, err := e.appContext.UserStore.RetrieveOne(&user)
	assert.Equal(t, nil, err)
	assert.Equal(t, rowsAffected, 1)
}

func TestUpdateAddress(t *testing.T) {
	// s := echo.New()
	// // g := s.Group("/api")

	// // req := httptest.NewRequest(echo.GET, "/api/users/"+e.testUser.ID, nil)
	// rec := httptest.NewRecorder()

	// // userCtrl := &User{}

	// cc := &context.AppContext{
	// 	Config:    e.config,
	// 	UserStore: &UserFakeStore{},
	// }

	// s.Use(middleware.AppContext(cc))

	// // g.GET("/users/:id", userCtrl.GetUserJSON)
	// // s.ServeHTTP(rec, req)

	// assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdatePaymentMethod(t *testing.T) {

}

func TestLogin(t *testing.T) {

}

func TestChangePassword(t *testing.T) {

}

func TestForgetPasswordByEmail(t *testing.T) {

}

func TestForgetPasswordByAccount(t *testing.T) {

}
