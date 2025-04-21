package models

type Agent struct {
	ID            int     `json:"id"`
	Email         string  `json:"email"`
	Name          string  `json:"name"`
	Type          int     `json:"type"`
	TypeAsString  string  `json:"type_as_string"`
	IsAvailable   bool    `json:"is_available"`
}

type MaxCustomerPerAgentRequest struct {
	MaxCustomerPerAgent int `json:"max_customer_per_agent"`
}