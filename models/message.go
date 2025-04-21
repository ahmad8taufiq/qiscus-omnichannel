package models

type Message struct {
	CandidateAgent	*Agent	`json:"candidate_agent"`
	Email 			string	`json:"email"`
	IsNewSession    bool	`json:"is_new_session"`
	IsResolved      bool	`json:"is_resolved"`
	LatestService	string	`json:"latest_service"`
	Name 			string	`json:"name"`
	RoomId 			string	`json:"room_id"`
}

type Agent struct {
	Id				int		`json:"id"`
	Email 			string	`json:"email"`
	Name			string	`json:"name"`
	Type			int		`json:"type"`
	TypeAsString	string	`json:"type_as_string"`
	IsAvailable		bool	`json:"is_available"`
}
