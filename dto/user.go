package dto

type LoginRequest struct {
	Username string `json:"username" valid:"required~username tidak boleh kosong"`
	Password string `json:"password" valid:"required~password tidak boleh kosong"`
}

type TokenResponse struct {
	Token string `json:"username"`
}

type LoginResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statuscode"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}
