package store

type Database interface {
	RetrieveOne(m *any) (int64, error)

	Create(m *any) (int64, error)

	RetrieveMany(m *[]any) (int64, error)

	RetrieveMany(conditions *map[string]interface{}, m *[]models.User) (int64, error)

	UpdateMany(conditions *map[string]interface{}, m *[]models.User) (int64, error)

	UpdateOne(m *any) (int64, error)
	
	Delete(m *any) (int64, error)

	Ping() error
}
