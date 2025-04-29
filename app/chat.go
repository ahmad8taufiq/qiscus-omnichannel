package app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"qiscus-omnichannel/models"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
	"qiscus-omnichannel/tools/response"
)

func ChatWithDelayHandler(chatSvc service.ChatService, commentSvc service.CommentService) http.HandlerFunc {
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

		for _, req := range requests {
			chatResp, err := chatSvc.InitiateChat(&req)
			if err != nil {
				logger.Logger.WithError(err).Error("Failed to initiate chat")
				results = append(results, map[string]interface{}{
					"user_id": req.UserID,
					"error":   "initiate_chat_failed",
				})
				continue
			}

			commentReq := models.PostCommentRequest{
				UserID:  req.UserID,
				RoomID:  chatResp.Data.CustomerRoom.RoomID,
				Message: "Halo, saya (" + req.UserID + ") ingin bertanya tentang",
			}

			if delayMs > 0 {
				time.Sleep(time.Duration(delayMs) * time.Second)
			}

			_, err = commentSvc.PostComment(chatResp.Data.IdentityToken, &commentReq)
			if err != nil {
				logger.Logger.WithError(err).Error("Failed to post comment")
				results = append(results, map[string]interface{}{
					"user_id": req.UserID,
					"room_id": chatResp.Data.CustomerRoom.RoomID,
					"error":   "post_comment_failed",
				})
				continue
			}

			results = append(results, map[string]interface{}{
				"user_id": req.UserID,
				"room_id": chatResp.Data.CustomerRoom.RoomID,
				"status":  "success",
			})
		}

		response.Success(w, "Processed", results)
	}
}
