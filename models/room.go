package models

type MarkAsResolvedRequest struct {
	RoomID        string `json:"room_id"`
	Notes         string `json:"notes,omitempty"`
	LastCommentID string `json:"last_comment_id,omitempty"`
}

type MarkAsResolvedResponse struct {
	Data struct {
		RoomInfo struct {
			Room struct {
				RoomAvatarURL   string `json:"room_avatar_url"`
				RoomChannelID   string `json:"room_channel_id"`
				RoomID          string `json:"room_id"`
				RoomName        string `json:"room_name"`
				RoomOptions     string `json:"room_options"`
				RoomType        string `json:"room_type"`
			} `json:"room"`
		} `json:"room_info"`
		Service struct {
			CreatedAt             string `json:"created_at"`
			FirstCommentID        string `json:"first_comment_id"`
			FirstCommentTimestamp string `json:"first_comment_timestamp"`
			IsResolved            bool   `json:"is_resolved"`
			LastCommentID         string `json:"last_comment_id"`
			Notes                 string `json:"notes"`
			ResolvedAt            string `json:"resolved_at"`
			RetrievedAt           string `json:"retrieved_at"`
			RoomID                string `json:"room_id"`
			UpdatedAt             string `json:"updated_at"`
			UserID                int    `json:"user_id"`
		} `json:"service"`
	} `json:"data"`
}

type SdkRoomResponse struct {
	Results struct {
		Room struct {
			LastCommentID int64 `json:"last_comment_id"`
		} `json:"room"`
	} `json:"results"`
}