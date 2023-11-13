package context

import (
	"github.com/Stanxxy/stan-go-web/config"
	"github.com/Stanxxy/stan-go-web/internal/i18n"
	"github.com/Stanxxy/stan-go-web/internal/store"
	echo "github.com/labstack/echo/v4"
)

// AppContext is the new context in the request / response cycle
// We can use the db store, cache and central configuration
type AppContext struct {
	echo.Context
	UserStore store.UserDatabase
	// FoodStore store.Database
	// FoodStore store.FoodStore - we need another store class to visit the food, business, order, etc.
	CacheStore store.Cache
	Config     *config.Configuration
	Loc        i18n.I18ner
}
