package usuarios

import (
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

type ServicesUser interface {
	Create(Id_perfil int, Usuario string, Password string, Nombres string, Apellidos string) error
	Update(id int, Id_perfil int, Usuario string, Password string, Nombres string, Apellidos string) error
	Delete(id int) error
	GetByID(id int) (*User, error)
	GetByUser(usuario string) (*User, error)
	GetAll() ([]*User, error)
}

type service struct {
	repository ServicesUserRepository
	user       *models.User
	txID       string
}

func (s service) Create(Id_perfil int, Usuario string, Password string, Nombres string, Apellidos string) error {

	m := NewUser(Id_perfil, Usuario, Password, Nombres, Apellidos)

	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create User :", err)
		return err
	}
	return nil
}

func (s service) Update(id int, Id_perfil int, Usuario string, Password string, Nombres string, Apellidos string) error {
	m := NewUser(Id_perfil, Usuario, Password, Nombres, Apellidos)
	m.ID = id
	valid, err := m.Validate()
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return err
	}

	if err := s.repository.Update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update User :", err)
		return err
	}
	return nil
}

func (s service) Delete(id int) error {
	u, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get User :", err)
		return err
	}
	if u == nil {
		logger.Error.Println(s.txID, " - couldn't get User %d to delete:", id)
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		logger.Error.Println(s.txID, " - couldn't delete User :", err)
		return err
	}
	return nil
}

func (s service) GetByID(id int) (*User, error) {
	u, err := s.repository.GetByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get User :", err)
		return nil, err
	}
	return u, nil
}

func (s service) GetAll() ([]*User, error) {
	u, err := s.repository.GetAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get Users :", err)
		return nil, err
	}
	return u, nil
}
func (s service) GetByUser(usuario string) (*User, error) {
	u, err := s.repository.GetByUser(usuario)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get User :", err)
		return nil, err
	}
	return u, nil
}

func NewUserService(repository ServicesUserRepository, user *models.User, txID string) ServicesUser {
	return &service{
		repository: repository,
		user:       user,
		txID:       txID,
	}
}
