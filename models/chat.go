package models

type InitiateChatRequest struct {
	AppID     string      `json:"app_id"`
	ChannelID string      `json:"channel_id"`
	UserID    string      `json:"user_id"`
	Name      string      `json:"name"`
	Avatar    string      `json:"avatar"`
	Nonce     string      `json:"nonce"`
	Extras    interface{} `json:"extras"`
	Email     string      `json:"email"`
	Origin    string      `json:"origin"`
}

type InitiateChatResponse struct {
	Data struct {
		CustomerRoom struct {
			ChannelID            int         `json:"channel_id"`
			Extras               interface{} `json:"extras"`
			ID                   int         `json:"id"`
			IsHandledByBot       bool        `json:"is_handled_by_bot"`
			IsResolved           bool        `json:"is_resolved"`
			IsWaiting            interface{} `json:"is_waiting"`
			LastCommentSender    interface{} `json:"last_comment_sender"`
			LastCommentSenderType interface{} `json:"last_comment_sender_type"`
			LastCommentText      interface{} `json:"last_comment_text"`
			LastCommentTimestamp interface{} `json:"last_comment_timestamp"`
			LastCustomerTimestamp interface{} `json:"last_customer_timestamp"`
			Name                 string      `json:"name"`
			RoomBadge            interface{} `json:"room_badge"`
			RoomID               string      `json:"room_id"`
			SessionID            interface{} `json:"session_id"`
			Source               string      `json:"source"`
			UserAvatarURL        string      `json:"user_avatar_url"`
			UserID               string      `json:"user_id"`
		} `json:"customer_room"`
		IdentityToken string      `json:"identity_token"`
		IsSecure      bool        `json:"is_secure"`
		IsSessional   bool        `json:"is_sessional"`
		SdkUser       interface{} `json:"sdk_user"`
	} `json:"data"`
	Status int `json:"status"`
}
