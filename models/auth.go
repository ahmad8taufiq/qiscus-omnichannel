package models

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Data struct {
		User struct {
			ID                 int    `json:"id"`
			Name               string `json:"name"`
			Email              string `json:"email"`
			AuthenticationToken string `json:"authentication_token"`
			AppID              int    `json:"app_id"`
			SDKEmail           string `json:"sdk_email"`
			SDKKey             string `json:"sdk_key"`
		} `json:"user"`
	} `json:"data"`
}