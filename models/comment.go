package models

type PostCommentRequest struct {
	Comment string `json:"comment"`
	TopicID string `json:"topic_id"`
	UniqueTempID string `json:"unique_temp_id"`
	Type string `json:"type"`
	Payload interface{} `json:"payload"`
	Extras interface{} `json:"extras"`
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
