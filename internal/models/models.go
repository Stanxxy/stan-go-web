package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"database/sql"

	"github.com/Stanxxy/stan-go-web/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model facilitate database interactions
type Model struct {
	models map[string]reflect.Value
	isOpen bool
	DB     *gorm.DB
	Conn   *sql.DB
}

// NewModel returns a new Model without opening database connection
func NewModel() *Model {
	return &Model{
		models: make(map[string]reflect.Value),
	}
}

// IsOpen returns true if the Model has already established connection
// to the database
func (m *Model) IsOpen() bool {
	return m.isOpen
}

// OpenWithConfig opens database connection with the settings found in cfg
func (m *Model) OpenWithConfig(cfg *config.Configuration) error {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})

	if err != nil {
		return err
	}

	conn, err := db.DB()

	if err != nil {
		return err
	}

	// https://github.com/go-sql-driver/mysql/issues/461
	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxIdleConns(0)
	conn.SetMaxOpenConns(20)

	m.DB = db
	m.Conn = conn
	m.isOpen = true
	return nil
}

// Register adds the values to the models registry
func (m *Model) Register(values ...interface{}) error {

	// do not work on them.models first, this is like an insurance policy
	// whenever we encounter any error in the values nothing goes into the registry
	models := make(map[string]reflect.Value)
	if len(values) > 0 {
		for _, val := range values {
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr {
				rVal = rVal.Elem()
			}
			switch rVal.Kind() {
			case reflect.Struct:
				models[getTypeName(rVal.Type())] = reflect.New(rVal.Type())
			default:
				return errors.New("models must be structs")
			}
		}
	}
	for k, v := range models {
		m.models[k] = v
	}
	return nil
}

func (m *Model) InitAllTables() {
	for _, v := range m.models {
		m.DB.Migrator().DropTable(v.Interface())
		m.DB.Migrator().CreateTable(v.Interface())
	}
}

func (m *Model) InitConstraints() {
	// TODO
}

func (m *Model) InitIndexies() {
	// TODO
}

// AutoMigrateAll runs migrations for all the registered models
func (m *Model) AutoMigrateAll() {
	for _, v := range m.models {
		m.DB.AutoMigrate(v.Interface())
	}
}

// AutoDropAll drops all tables of all registered models
func (m *Model) AutoDropAll() {
	for _, v := range m.models {
		m.DB.Migrator().DropTable(v.Interface())
	}
}

func getTypeName(typ reflect.Type) string {
	if typ.Name() != "" {
		return typ.Name()
	}
	split := strings.Split(typ.String(), ".")
	return split[len(split)-1]
}
