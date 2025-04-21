package models

type AssignAgentRequest struct {
	RoomID  string `json:"room_id"`
	AgentID int    `json:"agent_id"`
}

type AssignAgentResponse struct {
	Data struct {
		AddedAgent Agent  `json:"added_agent"`
		Service    Service `json:"service"`
	} `json:"data"`
}

type Service struct {
	RoomID         string `json:"room_id"`
	UserID         int    `json:"user_id"`
	IsResolved     bool   `json:"is_resolved"`
	RetrievedAt    string `json:"retrieved_at"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
