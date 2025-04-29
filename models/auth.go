package models

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Data            Data            `json:"data"`
}

type Data struct {
	User User `json:"user"`
	Details         Details         `json:"details"`
}

type User struct {
	ID                     int     `json:"id"`
	Name                   string  `json:"name"`
	Email                  string  `json:"email"`
	AuthenticationToken    string  `json:"authentication_token"`
	AppID                  int     `json:"app_id"`
	SdkEmail               string  `json:"sdk_email"`
}

type Details struct {
	SdkUser      SdkUser    `json:"sdk_user"`
}

type SdkUser struct {
	Token      string         `json:"token"`
}

type NonceResponse struct {
	Results struct {
		ExpiredAt int64  `json:"expired_at"`
		Nonce     string `json:"nonce"`
	} `json:"results"`
}