package models

type PostCommentRequest struct {
	UserID  string `json:"user_id"`
	RoomID  string `json:"room_id"`
	Message string `json:"message"`
}

type PostCommentResponse struct {
	Results struct {
		Comment struct {
			Extras    map[string]interface{} `json:"extras"`
			ID        int64                  `json:"id"`
			Message   string                 `json:"message"`
			Payload   map[string]interface{} `json:"payload"`
			Timestamp string                 `json:"timestamp"`
			Type      string                 `json:"type"`
			User      struct {
				Active    bool                   `json:"active"`
				AvatarURL string                 `json:"avatar_url"`
				Extras    map[string]interface{} `json:"extras"`
				UserID    string                 `json:"user_id"`
				Username  string                 `json:"username"`
			} `json:"user"`
		} `json:"comment"`
	} `json:"results"`
	Status int `json:"status"`
}
