package core

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"github.com/Stanxxy/stan-go-web/config"
	"github.com/Stanxxy/stan-go-web/internal/cache"
	"github.com/Stanxxy/stan-go-web/internal/i18n"
	"github.com/Stanxxy/stan-go-web/internal/models"
	echo "github.com/labstack/echo/v4"
)

type Server struct {
	Echo          *echo.Echo            // HTTP middleware
	config        *config.Configuration // Configuration
	db            *gorm.DB              // Database connection
	cache         *redis.Client         // Redis cache connection
	modelRegistry *models.Model         // Model registry for migration
}

// NewServer will create a new instance of the application
func NewServer(config *config.Configuration) *Server {
	server := &Server{}
	server.config = config
	i18n.Configure(config.LocaleDir, config.Lang, config.LangDomain)
	server.modelRegistry = models.NewModel()
	err := server.modelRegistry.OpenWithConfig(config)

	if err != nil {
		log.Fatalf("gorm: could not connect to db %q", err)
	}

	server.cache = cache.NewCache(config)
	server.db = server.modelRegistry.DB
	server.Echo = NewRouter(server)

	return server
}

func (s *Server) InitDB() {
	// migration for dev
	// Do not init data. We will manually init data from cockroachdb interface

	s.modelRegistry.AutoDropAll()

	user := models.User{}
	err := s.modelRegistry.Register(user)

	userqa := models.UserQA{}
	err = s.modelRegistry.Register(userqa)
	if err != nil {
		s.Echo.Logger.Fatal(err)
	}

	order := models.Order{}
	err = s.modelRegistry.Register(order)
	if err != nil {
		s.Echo.Logger.Fatal(err)
	}

	orderDetails := models.OrderDetails{}
	err = s.modelRegistry.Register(orderDetails)

	food := models.Food{}
	err = s.modelRegistry.Register(food)

	cartItem := models.CartItem{}
	err = s.modelRegistry.Register(cartItem)

	s.modelRegistry.AutoMigrateAll()
}

// GetDB returns gorm (ORM)
func (s *Server) GetDB() *gorm.DB {
	return s.db
}

// GetCache returns the current redis client
func (s *Server) GetCache() *redis.Client {
	return s.cache
}

// GetConfig return the current app configuration
func (s *Server) GetConfig() *config.Configuration {
	return s.config
}

// GetModelRegistry returns the model registry
func (s *Server) GetModelRegistry() *models.Model {
	return s.modelRegistry
}

// Start the http server
func (s *Server) Start(addr string) error {
	return s.Echo.Start(addr)
}

// ServeStaticFiles serve static files for development purpose
func (s *Server) ServeStaticFiles() {
	s.Echo.Static("/assets", s.config.AssetsBuildDir)
}

// GracefulShutdown Wait for interrupt signal
// to gracefully shutdown the server with a timeout of 5 seconds.
func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// close cache
	if s.cache != nil {
		cErr := s.cache.Close()
		if cErr != nil {
			s.Echo.Logger.Fatal(cErr)
		}
	}

	// close database connection
	if s.modelRegistry.Conn != nil {
		dErr := s.modelRegistry.Conn.Close()
		if dErr != nil {
			s.Echo.Logger.Fatal(dErr)
		}
	}

	// shutdown http server
	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
	}
}
