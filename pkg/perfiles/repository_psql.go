package perfiles

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

func NewPerfilPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psqldb {
	return &psqldb{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psqldb) Create(m *Perfil) error {
	const sqlInsert = `INSERT INTO public.perfiles (nombre,created_at, updated_at) VALUES (:nombre, Now(), Now()) `
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Perfil: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psqldb) Update(m *Perfil) error {
	const sqlUpdate = `UPDATE public.perfiles SET nombre = :nombre , updated_at = Now() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Perfil: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psqldb) Delete(id int) error {
	const sqlDelete = `DELETE FROM public.perfiles WHERE id = :id `
	m := Perfil{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Perfil: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID obtiene un registro de la BD
func (s *psqldb) GetByID(id int) (*Perfil, error) {
	const sqlGetByID = `SELECT id, nombre, created_at, updated_at FROM public.perfiles WHERE id = $1 `
	mdl := Perfil{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get Perfil: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll obtiene todos los registros de la BD
func (s *psqldb) GetAll() ([]*Perfil, error) {
	const sqlGetAll = `SELECT id, nombre, created_at, updated_at FROM public.perfiles `
	ms := []*Perfil{}
	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't get all Perfil: %v", err)
		return ms, err
	}
	return ms, nil
}
