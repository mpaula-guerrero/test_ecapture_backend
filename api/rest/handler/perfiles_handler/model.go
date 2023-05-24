package perfiles_handler

type PerfilRequest struct {
	ID   string `json:"id"`
	Name string `json:"nombre"`
}

type PerfilResponse struct {
	ID   string `json:"id"`
	Name string `json:"nombre"`
}
