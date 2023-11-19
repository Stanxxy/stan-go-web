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

func TestSignUp(t *testing.T) {

	// Given
	body := map[string]string{
		"account":   "dummy_account",
		"password":  "dummy_password",
		"email":     "stan.x.liu@gmail.com", // save the real email for blocking test cases
		"phoneNum":  "9939283432",
		"checkCode": "dummy_checkcode",
		"zipcode":   "03002",
	}
	jsonData, _ := json.Marshal(body) // ignore the error

	// When
	req := httptest.NewRequest(echo.POST, "/api/user/signUp", bytes.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.server.Echo.ServeHTTP(rec, req)

	// Then

	assert.Equal(t, http.StatusOK, rec.Code)

	// Search in the database to check if the write succeeded or not
	user := models.User{
		Account:  body["account"],
		Password: body["password"],
		Email:    body["email"],
		PhoneNum: body["phoneNum"],
		Zipcode:  body["zipcode"],
	}
	rowsAffected, err := e.appContext.UserStore.RetrieveOne(&user)
	assert.Equal(t, nil, err)
	assert.Equal(t, rowsAffected, int64(1))
}

func initData() {
	// If we need to have pre test data set up inin database just initialize them here
}

func TestUpdateAddress(t *testing.T) {
	initData()
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
