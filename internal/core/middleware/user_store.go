package middleware

import (
	"database/sql"

	"github.com/Stanxxy/stan-go-web/internal/models"
	"gorm.io/gorm"
)

// UserStore implements the UserStore interface
type UserStore struct {
	DB   *gorm.DB
	Conn *sql.DB
}

// Update the interface to satisfy CRUD ops
func (s *UserStore) RetrieveOne(m *models.User) (int64, error) {
	result := s.DB.First(m) // assuree we could do type conversion like this
	return result.RowsAffected, result.Error
}

func (s *UserStore) Create(m *models.User) (int64, error) {
	result := s.DB.Create(*m)
	return result.RowsAffected, result.Error
}

func (s *UserStore) RetrieveManyNoCondition(m *[]models.User) (int64, error) {
	result := s.DB.Find(m)
	return result.RowsAffected, result.Error
}

func (s *UserStore) RetrieveManyWithCondition(conditions *map[string]any, m *[]models.User) (int64, error) {
	result := s.DB.Where(conditions).Find(m)
	return result.RowsAffected, result.Error
}

func (s *UserStore) Delete(m *models.User) (int64, error) {
	result := s.DB.Delete(m)
	return result.RowsAffected, result.Error
}

func (s *UserStore) UpdateOne(m *models.User) (int64, error) {
	result := s.DB.Updates(*m)
	return result.RowsAffected, result.Error
}

func (s *UserStore) UpdateMany(conditions *map[string]any, m *[]models.User) (int64, error) {
	result := s.DB.Where(conditions).Updates(*m)
	return result.RowsAffected, result.Error
}

func (s *UserStore) Ping() error {
	return s.Conn.Ping()
}
