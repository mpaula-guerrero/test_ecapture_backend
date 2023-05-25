package usuarios

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/internal/models"
)

type psqldb struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewUserPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psqldb {
	return &psqldb{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psqldb) Create(m *User) error {
	const sqlInsert = `INSERT INTO public.usuarios(id_perfil, usuario, password, nombres, apellidos, created_at, updated_at) 
						VALUES (:id_perfil, :usuario, :password, :nombres, :apellidos, Now(), Now())`
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert User: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psqldb) Update(m *User) error {
	const sqlUpdate = `UPDATE public.usuarios SET id_perfil = :id_perfil, usuario = :usuario, password = :password, nombres = :nombres, apellidos = :apellidos , updated_at = Now() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update user: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psqldb) Delete(id int) error {
	const sqlDelete = `DELETE FROM public.usuarios WHERE id = :id `
	m := User{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete User: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID obtiene un registro de la BD
func (s *psqldb) GetByID(id int) (*User, error) {
	const sqlGetByID = `SELECT id, id_perfil, usuario, password, nombres, apellidos, created_at, updated_at FROM public.usuarios WHERE id = $1 `
	mdl := User{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get User: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetByUser obtiene un registro de la BD
func (s *psqldb) GetByUser(usuario string) (*User, error) {
	const sqlGetByUser = `SELECT id, id_perfil, usuario, password, nombres, apellidos, created_at, updated_at FROM public.usuarios WHERE usuario = $1 `
	mdl := User{}
	err := s.DB.Get(&mdl, sqlGetByUser, usuario)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get User: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll obtiene todos los registros de la BD
func (s *psqldb) GetAll() ([]*User, error) {
	const sqlGetAll = `SELECT id, id_perfil, usuario, password, nombres, apellidos, created_at, updated_at FROM public.usuarios `
	ms := []*User{}
	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get all User: %v", err)
		return ms, err
	}
	return ms, nil
}
