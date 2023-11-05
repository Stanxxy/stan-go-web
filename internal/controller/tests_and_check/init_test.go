package controller_test

import (
	"os"
	"testing"

	"github.com/Stanxxy/stan-go-web/config"
	"github.com/Stanxxy/stan-go-web/internal/core"
	"github.com/Stanxxy/stan-go-web/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var e struct {
	config   *config.Configuration
	logger   *log.Logger
	server   *core.Server
	testUser *models.User
}

func TestMain(m *testing.M) {
	e.config = &config.Configuration{
		ConnectionString: "host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword",
		TemplateDir:      "../templates/*.html", // we dont have template dir right now
		LayoutDir:        "../templates/layouts/*.html", // we dont have layouts dir right now
		Dialect:          "postgres",
		RedisAddr:        ":6379",
	}

	e.server = core.NewServer(e.config)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	userCtrl := &User{}
	healthCtrl := &Healthcheck{}

	// g := e.server.Echo.Group("/api")
	// g.GET("/users/:id", userCtrl.GetUserJSON)

	u := e.server.Echo.Group("/users")
	u.GET("/:id", userCtrl.GetUser)

	e.server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	e.server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler()))

	// test data
	user := models.User{Name: "Peter"}
	mr := e.server.GetModelRegistry()
	err := mr.Register(user)

	if err != nil {
		e.server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()
	mr.Save(&user)

	e.testUser = &user
}

func tearDown() {
	e.server.GetModelRegistry().AutoDropAll()
}
