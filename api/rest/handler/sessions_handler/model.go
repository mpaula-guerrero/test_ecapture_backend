package sessions_handler

type LoginRequest struct {
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
}

type LoginResponse struct {
}

