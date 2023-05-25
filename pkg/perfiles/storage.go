package perfiles

import (
	"github.com/jmoiron/sqlx"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesPerfilRepository interface {
	Create(m *Perfil) error
	Update(m *Perfil) error
	Delete(id int) error
	GetByID(id int) (*Perfil, error)
	GetAll() ([]*Perfil, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesPerfilRepository {
	var s ServicesPerfilRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return NewPerfilPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
