package models

type CustomerResponse struct {
	Customer   CustomerInfo   `json:"customer"`
	ResolvedBy ResolvedByInfo `json:"resolved_by"`
	Service    ServiceInfo    `json:"service"`
}

type CustomerInfo struct {
	AdditionalInfo []interface{} `json:"additional_info"`
	Avatar         string        `json:"avatar"`
	Name           string        `json:"name"`
	UserID         string        `json:"user_id"`
}

type ResolvedByInfo struct {
	Email       string `json:"email"`
	ID          int `json:"id"`
	IsAvailable bool   `json:"is_available"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type ServiceInfo struct {
	FirstCommentID string      `json:"first_comment_id"`
	ID             int         `json:"id"`
	IsResolved     bool        `json:"is_resolved"`
	LastCommentID  string      `json:"last_comment_id"`
	Notes          interface{} `json:"notes"` // can be `*string` if you prefer null vs non-null
	RoomID         string      `json:"room_id"`
	Source         string      `json:"source"`
}
