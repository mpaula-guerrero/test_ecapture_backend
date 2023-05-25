package usuarios_handler

type UserRequest struct {
	ID        int    `json:"id"`
	Id_perfil int    `json:"id_perfil"`
	Usuario   string `json:"usuario"`
	Password  string `json:"password"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos" `
}

type UserResponse struct {
	ID        int    `json:"id"`
	Id_perfil int    `json:"id_perfil"`
	Usuario   string `json:"usuario"`
	Password  string `json:"password"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos" `
}
