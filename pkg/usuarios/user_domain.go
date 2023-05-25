package usuarios

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id" valid:"-"`
	Id_perfil int       `json:"id_perfil" db:"id_perfil" valid:"-"`
	Usuario   string    `json:"usuario" db:"usuario" valid:"required"`
	Password  string    `json:"password" db:"password" valid:"required"`
	Nombres   string    `json:"nombres" db:"nombres" valid:"required"`
	Apellidos string    `json:"apellidos" db:"apellidos" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUser(Id_perfil int, Usuario string, Password string, Nombres string, Apellidos string) *User {
	return &User{
		Id_perfil: Id_perfil,
		Usuario:   Usuario,
		Password:  Password,
		Nombres:   Nombres,
		Apellidos: Apellidos,
	}
}
func (m *User) Validate() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
