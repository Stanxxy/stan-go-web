package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Stanxxy/stan-go-web/internal/core"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// GoogleOAuthConfig is the OAuth2 configuration for Google authentication.
var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET"),
	RedirectURL:  "http://localhost:8080/api/auth/google/callback",
	Scopes:       []string{"profile", "email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://oauth2.googleapis.com/token",
	},
}

// FacebookOAuthConfig is the OAuth2 configuration for Facebook authentication.
var FacebookOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("FACEBOOK_CLEINT_ID"),
	ClientSecret: os.Getenv("FACEBOOK_SECRET"),
	RedirectURL:  "http://localhost:8080/api/auth/facebook/callback",
	Scopes:       []string{"public_profile", "email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://www.facebook.com/v17.0/dialog/oauth",
		TokenURL: "https://graph.facebook.com/v17.0/oauth/access_token",
	},
}

// HandleGoogleLogin redirects the user to the Google authentication page.
func HandleGoogleLogin(c echo.Context) error {
	url := GoogleOAuthConfig.AuthCodeURL("state")
	fmt.Printf(url)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleGoogleCallback handles the callback from the Google authentication page.
func HandleGoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	fmt.Printf(code)
	token, err := GoogleOAuthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		b := errors.NewBoom(errors.InternalError, errors.ErrorText(errors.InternalError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	c.Logger().Info("get token", token)

	// TODO: Verify token and create user account or sign in existing user.

	return c.String(http.StatusOK, "Google Callback verified")
}

// HandleFacebookLogin redirects the user to the Facebook authentication page.
func HandleFacebookLogin(c echo.Context) error {
	url := FacebookOAuthConfig.AuthCodeURL("state")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleFacebookCallback handles the callback from the Facebook authentication page.
func HandleFacebookCallback(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := FacebookOAuthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		b := errors.NewBoom(errors.InternalError, errors.ErrorText(errors.InternalError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	c.Logger().Info("get token", token)

	// TODO: Verify token and create user account or sign in existing user.
	return c.String(http.StatusOK, "Facebook Callback verified")
}

// RegisterAuthRoutes registers the authentication routes with the provided router.
func RegisterAuthRoutes(server *core.Server) {

	g := server.Echo.Group("/api")
	g.GET("/auth/google/login", HandleGoogleLogin)
	g.GET("/auth/google/callback", HandleGoogleCallback)
	g.GET("/auth/facebook/login", HandleFacebookLogin)
	g.GET("/auth/facebook/callback", HandleFacebookCallback)
}
