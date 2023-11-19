package controller_test

import (
	"os"
	"testing"

	"github.com/Stanxxy/stan-go-web/config"
	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/controller"
	"github.com/Stanxxy/stan-go-web/internal/core"
	mid "github.com/Stanxxy/stan-go-web/internal/core/middleware"
	"github.com/Stanxxy/stan-go-web/internal/i18n"
	"github.com/Stanxxy/stan-go-web/internal/models"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var e struct {
	config               *config.Configuration
	logger               *log.Logger
	server               *core.Server
	userPlaceHolder      *models.User   // the space that is used to get retrieved object from database
	userSlicePlaceHolder *[]models.User // the space that is used to get retrieved object from database
	appContext           *context.AppContext
}

func TestMain(m *testing.M) {
	e.config = &config.Configuration{
		// ConnectionString: "host=localhost user=goweb dbname=goweb port=26257 sslmode=disable TimeZone=US/Eastern",
		DNS:         "host=localhost user=goweb_test dbname=goweb_test port=26258 sslmode=disable TimeZone=US/Eastern",
		Address:     "127.0.0.1:8090",
		TemplateDir: "../templates/*.html",         // we dont have template dir right now
		LayoutDir:   "../templates/layouts/*.html", // we dont have layouts dir right now
		Dialect:     "postgres",
		RedisAddr:   ":6379",
	}

	e.server = core.NewServer(e.config)

	setup()
	code := m.Run()
	tearDown()

	os.Exit(code)
}

func setup() {
	healthCtrl := &controller.Healthcheck{}

	s := e.server
	controller.RegisterAuthRoutes(s)
	controller.RegisterUserRoutes(s)
	controller.RegisterBusinessRoutes(s)

	// health check and metrics
	e.server.Echo.GET("/.well-known/health-check", healthCtrl.GetHealthcheck)
	e.server.Echo.GET("/.well-known/metrics", echo.WrapHandler(promhttp.Handler())) // what does this do?

	e.server.InitDB()
	// e.server.GetModelRegistry().InitAllTables()

	e.appContext = &context.AppContext{
		CacheStore: &mid.CacheStore{Cache: e.server.GetCache()},
		Config:     e.config,
		UserStore:  &mid.UserStore{DB: e.server.GetDB(), Conn: e.server.GetModelRegistry().Conn},
		Loc:        i18n.New(),
	}

	// test data
	// Here we initialize some test data that should be in database prior to any test running
	user := models.User{}
	e.userPlaceHolder = &user
	userList := []models.User{user}
	e.userPlaceHolder = &userList[0]
}

func tearDown() {
	e.server.GetModelRegistry().AutoDropAll()
}

// code landfill
// mr := e.server.GetModelRegistry()

// if err != nil {
// 	e.server.Echo.Logger.Fatal(err)
// }

// mr.AutoMigrateAll()
// mr.Save(&user)
