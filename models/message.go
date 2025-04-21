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
