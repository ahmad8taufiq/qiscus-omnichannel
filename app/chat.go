package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
	"qiscus-omnichannel/tools/response"
)

func ChatWithDelayHandler(chatSvc service.ChatService, commentSvc service.CommentService, authSvc service.AuthService, roomSvc service.RoomService, redisSvc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			response.NotFound(w, "Method not allowed")
			return
		}

		delayMs := 0
		delayParam := r.URL.Query().Get("delay")
		if delayParam != "" {
			val, err := strconv.Atoi(delayParam)
			if err != nil {
				logger.Logger.WithError(err).Error("Invalid delay value")
				response.BadRequest(w, "Invalid delay query param")
				return
			}
			delayMs = val
		}

		var requests []models.InitiateChatRequest
		if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
			logger.Logger.WithError(err).Error("Invalid JSON")
			response.BadRequest(w, "Invalid JSON format")
			return
		}
		defer r.Body.Close()

		var results []map[string]interface{}

		nonce, err := GetNonce(redisSvc)
		if err != nil {
			logger.Logger.WithError(err).Error("Failed to get nonce")
		}

		for i := range requests {
			requests[i].AppID = config.AppConfig.QiscusAppID
			requests[i].ChannelID = "130821"
			requests[i].Avatar = "https://omnichannel.qiscus.com/img/ic_qiscus_client.png"
			requests[i].Nonce = nonce
			requests[i].Email = requests[i].UserID
			requests[i].Origin = "https://omnichannel.qiscus.com/iframes/v4/rvcbl-fcsngqk40iyo7ks/multichannel-widget/130821"

			initChatResp, err := chatSvc.InitiateChat(&requests[i])
			if err != nil {
				logger.Logger.WithError(err).Error("Failed to initiate chat")
				results = append(results, map[string]interface{}{
					"user_id": requests[i].UserID,
					"error":   "initiate_chat_failed",
				})
				continue
			}

			verifyTokenResp, err := authSvc.VerifyToken(&models.VerifyTokenRequest{
				IdentityToken: initChatResp.Data.IdentityToken,
			})
			if err != nil {
				logger.Logger.WithError(err).Error("Failed to verify token")
				results = append(results, map[string]interface{}{
					"user_id": requests[i].UserID,
					"error":   "verify_token_failed",
				})
				continue
			}

			roomIdResp, err := roomSvc.GetRoomById(initChatResp.Data.CustomerRoom.RoomID, verifyTokenResp.Results.User.Token, "")
			if err != nil {
				logger.Logger.WithError(err).Error("Failed to get room by id")
				results = append(results, map[string]interface{}{
					"user_id": requests[i].UserID,
					"error":   "get_room_by_id_failed",
				})
				continue
			}

			commentReq := models.PostCommentRequest{
				Comment:  "Halo, saya " + requests[i].Name + " (" + requests[i].Email + ") ingin bertanya tentang",
				TopicID: strconv.FormatInt(roomIdResp.Results.Room.LastTopicID, 10),
				Type:    "text",
				Payload: nil,
				Extras: map[string]interface{}{},
			}

			if delayMs > 0 {
				time.Sleep(time.Duration(delayMs) * time.Second)
			}

			resp, err := commentSvc.PostComment(verifyTokenResp.Results.User.Token, initChatResp.Data.CustomerRoom.UserID, &commentReq)
			if err != nil {
				logger.Logger.WithError(err).Error("Failed to post comment")
				results = append(results, map[string]interface{}{
					"comment": commentReq.Comment,
					"topic_id": commentReq.TopicID,
					"error":   "post_comment_failed",
				})
				continue
			}

			results = append(results, map[string]interface{}{
				"comment": commentReq.Comment,
				"topic_id": commentReq.TopicID,
				"status":  resp.Status,
			})
		}

		response.Success(w, "Processed", results)
	}
}
