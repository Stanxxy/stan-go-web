package controller

// Request Types
type (
	SignUpRequest struct {
		Account    string `json:"account" validate:"required"`
		Password   string `json:"password" validate:"required"`
		Email      string `json:"email" validate:"required"`
		PhoneNum   string `json:"phoneNum" validate:"required"`
		CheckCode  string `json:"checkCode" validate:"required"`
		Zipcode   string `json:"zipcode"  validate:"required"`
	}
	UpdateAddressRequest struct {
		UserId string `json:"userId" validate:"required"`
		State  string `json:"state" validate:"required"`
		City   string `json:"city" validate:"required"`
		Street string `json:"street" validate:"required"`
		Unit   string `json:"unit"`
		Zipcode   string `json:"zipcode" validate:"required"`
	}
	UpdatePaymentInfoRequest struct {
		UserId string `json:"account" validate:"required"`
		Zelle  string `json:"zelle"`
		Venmo  string `json:"venmo"`
		WeChat string `json:"WeChat"`
		Alipay string `json:"Alipay"`
		BTC    string `json:"BTC"`
	}
	LoginRequest struct {
		Account  string `json:"account" default:"*NULL*"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" default:"*NULL*"`
		PhoneNum string `json:"phoneNum" default:"*NULL*"`
	}
	ChangePasswordRequest struct {
		NewPassword string `json:"newPassword"`
		ChangeRequestId string `json:"changeRequestId"`
		UserId      string `json:"userId"`
	}
	ForgetPasswordByEmailRequest struct {
		Email     string `json:"email"`
		CheckCode string `json:"checkCode"`
	}
	ForgetPasswordByAccountRequest struct {
		Account   string `json:"account"`
		CheckCode string `json:"checkCode"`
	}
)

// Response Types
type (
	SignUpResponse struct {
		UserId  string `json:"userId"`
		Message string `json:"message"`
	}
	UpdateAddressResponse struct {
		UserId  string `json:"userId"`
		Message string `json:"message"`
	}
	UpdatePaymentInfoResponse struct {
		UserId  string `json:"userId"`
		Message string `json:"message"`
	}
	LoginRsponse struct {
		UserId  string `json:"userId"`
		Message string `json:"message"`
	}
	ChangePasswordRsponse struct {
		UserId  string `json:"userId"`
		Message string `json:"message"`
	}
	ForgetPasswordByEmailRsponse struct {
		UserId  string `json:"userId"`
		Message string `json:"message"`
	}
	ForgetPasswordByAccountRsponse struct {
		UserId    string `json:"userId"`
		Message   string `json:"message"`
	}
)
