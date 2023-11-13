package store

import "github.com/Stanxxy/stan-go-web/internal/models"

type UserDatabase interface {
	RetrieveOne(m *models.User) (int64, error)

	Create(m *models.User) (int64, error)

	RetrieveManyNoCondition(m *[]models.User) (int64, error)

	RetrieveManyWithCondition(conditions *map[string]any, m *[]models.User) (int64, error)

	UpdateMany(conditions *map[string]any, m *[]models.User) (int64, error)

	UpdateOne(m *models.User) (int64, error)

	Delete(m *models.User) (int64, error)

	Ping() error
}
