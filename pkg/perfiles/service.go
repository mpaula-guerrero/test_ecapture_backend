package perfiles

import (
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

type ServicesPerfil interface {
	Create(nombre string) error
	Update(id int, nombre string) error
	Delete(id int) error
	GetByID(id int) (*Perfil, error)
	GetAll() ([]*Perfil, error)
}

type service struct {
	repository ServicesPerfilRepository
	user       *models.User
	txID       string
}

func (s service) Create(nombre string) error {

	m := NewPerfil(nombre)

	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create Perfil :", err)
		return err
	}
	return nil
}

func (s service) Update(id int, nombre string) error {
	m := NewPerfil(nombre)
	m.ID = id
	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Perfil :", err)
		return err
	}
	return nil
}

func (s service) Delete(id int) error {
	p, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Perfil :", err)
		return err
	}
	if p == nil {
		logger.Error.Println(s.txID, " - couldn't get Perfil %d to delete:", id)
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		logger.Error.Println(s.txID, " - couldn't delete Perfil :", err)
		return err
	}
	return nil
}

func (s service) GetByID(id int) (*Perfil, error) {
	p, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Perfil :", err)
		return nil, err
	}
	return p, nil
}

func (s service) GetAll() ([]*Perfil, error) {
	p, err := s.repository.GetAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Perfiles :", err)
		return nil, err
	}
	return p, nil
}

func NewPerfilService(repository ServicesPerfilRepository, user *models.User, txID string) ServicesPerfil {
	return &service{
		repository: repository,
		user:       user,
		txID:       txID,
	}
}
