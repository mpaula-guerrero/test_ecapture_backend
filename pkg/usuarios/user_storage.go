package usuarios

import (
	"github.com/jmoiron/sqlx"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesUserRepository interface {
	Create(m *User) error
	Update(m *User) error
	Delete(id int) error
	GetByID(id int) (*User, error)
	GetByUser(usuario string) (*User, error)
	GetAll() ([]*User, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserRepository {
	var s ServicesUserRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return NewUserPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
