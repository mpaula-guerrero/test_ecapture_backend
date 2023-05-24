package perfiles

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Perfil struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	Nombre    string    `json:"nombre" db:"nombre" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewPerfil(nombre string) *Perfil {
	return &Perfil{
		Nombre: nombre,
	}
}

func (m *Perfil) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
