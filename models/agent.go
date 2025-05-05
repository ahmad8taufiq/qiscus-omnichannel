package models

type Agent struct {
	ID            int     `json:"id"`
	Email         string  `json:"email"`
	Name          string  `json:"name"`
	Type          int     `json:"type"`
	TypeAsString  string  `json:"type_as_string"`
	IsAvailable   bool    `json:"is_available"`
}

type AgentRequest struct {
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

type MaxCustomerPerAgentRequest struct {
	MaxCustomerPerAgent int `json:"max_customer_per_agent"`
}

type AvailableAgentsResponse struct {
	Data struct {
		Agents []Agents `json:"agents"`
	} `json:"data"`
}

type Agents struct {
	ID                   int    `json:"id"`
	CurrentCustomerCount int    `json:"current_customer_count"`
}