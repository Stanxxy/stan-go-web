package controller

import (
	"net/http"
	"regexp"
	"fmt"

	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/Stanxxy/stan-go-web/internal/models"
	"github.com/Stanxxy/stan-go-web/internal/utils"
	"github.com/Stanxxy/stan-go-web/internal/models/user"
	echo "github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

func (ctrl User) SignUp(c echo.Context) {
	cc := c.(*context.AppContext)

	var signupRequest SignUpRequest

	err := cc.Bind(&signUpRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	resp := SignUpResponse{}
	userNew := models.User{
		Account:  signupRequest.Account,
		Password: signupRequest.Password,
		PhoneNum: signupRequest.PhoneNum,
		Email:    signupRequest.Email,
		ZipCode:  signupRequest.Zipcode,
	}

	// First check if the account existed in database.

	// Check phone
	userFilter := models.User{PhoneNum: userNew.PhoneNum}
	rowsAffected, err := cc.UserStore.RetrieveOne(&userFilter)
	if rowsAffected >= 1 || err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Check email
	userFilter = models.User{Email: userNew.Email}
	rowsAffected, err = cc.UserStore.RetrieveOne(&userFilter)
	if rowsAffected >= 1 || err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// check useraccount
	userFilter = models.User{Account: userNew.Account}
	rowsAffected, err = cc.UserStore.RetrieveOne(&userFilter)
	if rowsAffected >= 1 || err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// create user in databases
	err = cc.UserStore.Create(&userNew)

	if err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	resp.UserId = user.ID
	resp.Message = "Successfully signed up"
	return c.String(http.StatusOK, resp)
}

func (ctrl user.User) UpdateAddress(c echo.Context) error {
	cc := c.(*context.AppContext)

	var updateAddressRequest UpdateAddressRequest

	err := cc.Bind(&updateAddressRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	user := models.User{
		ID:       updateAddressRequest.UserId,
		State:    updateAddressRequest.State,
		City:     updateAddressRequest.City,
		Street:   updateAddressRequest.Street,
		Unit:     updateAddressRequest.Unit,
		Zipcode:     updateAddressRequest.Zipcode,
	}

	rowsAffected, err := cc.UserStore.UpdateOne(&user)

	if err != nil {
		b := errors.NewBoom(errors.EntityUpdateError, errors.ErrorText(errors.EntityUpdateError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	resp := UpdateAddressResponse{user.ID, "Address info updated"}
	return c.JSON(http.StatusOK, resp)
}

func (ctrl User) UpdatePaymentMethod(c echo.Context) error {
	cc := c.(*context.AppContext)

	var updatePaymentInfoRequest UpdatePaymentInfoRequest

	err := cc.Bind(&updatePaymentInfoRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// How to initialize a string array
	user := models.User{
		ID: updatePaymentInfoRequest.UserId,
		PaymentMethod: map[string]string{
			"Zelle":  updatePaymentInfoRequest.Zelle,
			"Venmo":  updatePaymentInfoRequest.Venmo,
			"WeChat": updatePaymentInfoRequest.WeChat,
			"Alipay": updateAddressRequest.Alipay,
			"BTC":    updateAddressRequest.BTC,
		},
	}

	rowsAffected, err := cc.UserStore.UpdateOne(&user)

	if err != nil {
		b := errors.NewBoom(errors.EntityUpdateError, errors.ErrorText(errors.EntityUpdateError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	resp := UpdatePaymentInfoResponse{user.ID, "payment info updated"}
	return c.JSON(http.StatusOK, resp)
}

func (ctrl User) Login(c echo.Context) {
	cc := c.(*context.AppContext)

	var loginRequest LoginRequest

	err := cc.Bind(&loginRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Check for which login method it choose
	user := models.User{}
	resp := LoginRsponse{}

	// Use phone number login we use verification code
	// use account/email login we use password
	if len(loginRequest.Account) != 0 {
		user.Account = loginRequest.Account
		user.Password = loginRequest.Password
	} else if len(loginRequest.Email) != 0 {
		user.Email = loginRequest.Email
		user.Password = loginRequest.Password
	} else if len(loginRequest.PhoneNum) != 0 {
		user.PhoneNum = loginRequest.PhoneNum
		// When we use phone number, we use sms to send verification code to log in
		// sms logic
	} else {
		resp.UserId = "-1"
		resp.Message = "Please put in account/phonenumber/email"
		return c.String(http.StatusOK, resp)
	}

	rowsAffected, err := cc.UserStore.RetrieveOne(&user)
	if err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Check found or not
	if rowsAffected == 0 {
		resp.UserId = "-1"
		resp.Message = "Wrong password or user name"
	}
	resp.UserId = user.ID
	resp.Message = "successfully login"
	return c.JSON(http.StatusOK, resp)
}

// TODO area: ------------------------------------------
// for changing password, we maintain a password update table in cache.
// for deatils please visit https://stackoverflow.com/questions/1102781/best-way-for-a-forgot-password-implementation
func (ctrl User) ChangePassword(c echo.Context) {
	cc := c.(*context.AppContext)

	var changePasswordRequest ChangePasswordRequest
	resp := ChangePasswordRsponse{}

	err := cc.Bind(&changePasswordRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// handle old password and new password
	// capture the changeRequestID in request body.
	val, err := cc.CacheStore.get(changePasswordRequest.ChangeRequestId)
	if err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	triplet := utils.PasswordChangeTriplet{}
	if err := triplet.Scan(val); err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	timeDiff := time.Now().Sub(triplet.ReqTime)
	if timeDiff.Hours() > 0 || timeDiff.Minutes() > 15{
		resp.UserId = triplet.UserId
		resp.Message = "link is expired"
		return c.JSON(http.StatusOK, resp)
	} 

	// for user interface it will be handled by frontend server
	
	// check if old password is equal to new password
	newPassword := utils.DecryptRequest(changePasswordRequest.NewPassword)
	// We will encrypt the user password in the future and just use normal account right now.
	user := models.User{ID:  changePasswordRequest.UserId}
	rowsAffected, err := cc.UserStore.RetrieveOne(&user)
	if err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	// Check found or not
	if rowsAffected == 0 {
		resp.UserId = "-1"
		resp.Message = "Wrong user id"
		return c.JSON(http.StatusOK, resp)
	}

	if user.Password == newPassword {
		resp.UserId = user.ID
		resp.Message = "The new password is the same with old one"
		return c.JSON(http.StatusOK, resp)
	}

	// security and sanity checking
	
	// rules: 12 - 16 in terms of length.
	if len(newPassword) < 12 || len(newPassword) > 16 {
		resp.UserId = user.ID
		resp.Message = "The password should be between 12 to 16 characters long"
		return c.JSON(http.StatusOK, resp)
	}

	// only include [0-9a-zA-Z'-!"#$%& ()*,./:;?@ []^_` {|}~+<=>]
	re := regexp.MustCompile("^(?=[a-zA-Z0-9~@#$^*()_+=[\]{}|\\,.?:-]*$)(?!.*[<>'\"/;`%])")
	if !re.matchString(newPassword) {
		resp.UserId = user.ID
		resp.Message = "The password should only contain digit0-9, alphabet letters a-z and A-Z, and special characters~@#$^*()_+=[\]{}|\\,.?:- "
		return c.JSON(http.StatusOK, resp)
	}


	// encrypt when saving password.
	user.Password = newPassword
	rowsAffected, err := cc.UserStore.UpdateOne(&user)
	if err != nil {
		b := errors.NewBoom(errors.EntityUpdateError, errors.ErrorText(errors.EntityUpdateError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	// Check found or not
	if rowsAffected == 0 {
		resp.UserId = "-1"
		resp.Message = "fail to update password"
		return c.JSON(http.StatusOK, resp)
	}
	
	resp.UserId = user.ID
	resp.Message = "password updated"
	return c.JSON(http.StatusOK, resp)
}

func (ctrl User) ForgetPasswordByEmail(c echo.Context) {
	// get password by providing email. Use when user forget the account
	cc := c.(*context.AppContext)

	var forgetPasswordByEmailRequest ForgetPasswordByEmailRequest
	resp := ForgetPasswordByEmailRsponse{}

	err := cc.Bind(&forgetPasswordByEmailRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// check the email in database
	user := models.User{
		Email: forgetPasswordByEmailRequest.Email
	}
	rowsAffected, err := cc.UserStore.RetrieveOne(&user)
	if err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	// Check found or not
	if rowsAffected == 0 {
		// if cannot find email, response with message
		resp.UserId = "-1"
		resp.Message = "the email has not been registered"
		return c.JSON(http.StatusOK, resp)
	}

	// if can find the email:

	// create an reqID(token), a time and a userid with reqID as the key
	receivedTime := time.Now()
	token = utils.CreateToken(user.Email, receivedTime)
	triplet := utils.PasswordChangeTriplet{
		ReqId: token,
		ReqTime: receivedTime,
		UserId: user.ID
	}

	// insert the triplet into the cache. ID has to be encrepted. Set ttl to be 24h
	requestTtl, _ := time.ParseDuration("48h")
	err := cc.CacheStore.Set(utils.EncryptToken(token), triplet, requestTtl)
	if err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	} 
	// call the smtp service and send an email with the new reqID as url parameter in a reset password page url
	smtpClient := utils.SmtpClient{}
	smtpClient.GetAuthFromGoogle("stan.x.liu.18@gmail.com", ":Qwer1997asdf")
	smtpClient.SendEmailFromGoogle("testEmail", "Hello smtp", "stan.x.liu.18@gmail.com", []string{user.Email})

	// response to front end to confirm email is sent. / cannot find the email.
	resp.UserId = user.ID
	resp.Message = "reset password request sent to your email."
	return c.JSON(http.StatusOK, resp)
}

func (ctrl User) ForgetPasswordByAccount(c echo.Context) {
	// get password by providing account. Use when user forget password but keeps account
	cc := c.(*context.AppContext)

	var forgetPasswordByAccountRequest ForgetPasswordByAccountRequest
	resp := ForgetPasswordByAccountRsponse{}

	err := cc.Bind(&forgetPasswordByAccountRequest)
	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// check the account in database
	user := models.User{
		Account: forgetPasswordByAccountRequest.Account
	}
	rowsAffected, err := cc.UserStore.RetrieveOne(&user)
	if err != nil {
		b := errors.NewBoom(errors.EntityQueryError, errors.ErrorText(errors.EntityQueryError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	
	// if cannot find account, response with error
	if rowsAffected == 0 {
		// if cannot find email, response with message
		resp.UserId = "-1"
		resp.Message = "the email has not been registered"
		return c.JSON(http.StatusOK, resp)
	}
	
	// if can find the account:

	// create an reqID(token), a time and a userid with reqID as the key
	receivedTime := time.Now()
	token = utils.CreateToken(user.Email, receivedTime)
	triplet := utils.PasswordChangeTriplet{
		ReqId: token,
		ReqTime: receivedTime,
		UserId: user.ID
	}

	// insert the triplet into the cache. ID has to be encrepted. Set ttl to be 24h
	requestTtl, _ := time.ParseDuration("48h")
	err := cc.CacheStore.Set(utils.EncryptToken(token), triplet, requestTtl)
	if err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	} 
	// call the smtp service and send an email with the new reqID as url parameter in a reset password page url
	smtpClient := utils.SmtpClient{}
	smtpClient.GetAuthFromGoogle("stan.x.liu.18@gmail.com", ":Qwer1997asdf")
	smtpClient.SendEmailFromGoogle("testEmail", fmt.Sprintf("<This is the frontend URL for password reset>?reqID=%s", token), "stan.x.liu.18@gmail.com", []string{user.Email})

	// response to front end to confirm email is sent. / cannot find the email.
	resp.UserId = user.ID
	resp.Message = fmt.Sprintf("reset password request sent to %s", utils.MuskAnEmail(&user.Email) )
	return c.JSON(http.StatusOK, resp)
}