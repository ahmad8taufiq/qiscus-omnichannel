package app

import (
	"io"
	"net/http"

	"qiscus-omnichannel/models"
	"qiscus-omnichannel/service"
	tools "qiscus-omnichannel/tools/parser"
	"qiscus-omnichannel/tools/response"

	"github.com/sirupsen/logrus"
)

func WebhookHandler(log *logrus.Logger, redisSvc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			response.NotFound(w, "Method not allowed")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithFields(logrus.Fields{
				"method":   r.Method,
				"endpoint": r.URL.Path,
				"error":    err.Error(),
			}).Error("Failed to read request body")
			response.BadRequest(w, "Failed to read request body")
			return
		}
		defer r.Body.Close()

		log.WithFields(logrus.Fields{
			"method":   r.Method,
			"endpoint": r.URL.Path,
			"raw":      string(body),
		}).Info("📩 Incoming request")

		data, err := tools.Parser[models.Message](body)
		if err != nil {
			log.WithError(err).Error("Failed to process webhook")
			response.BadRequest(w, "Invalid JSON format")
			return
		}

		log.WithFields(logrus.Fields{
			"room_id":        data.RoomId,
			"name":           data.Name,
			"email":          data.Email,
			"is_new_session": data.IsNewSession,
			"is_resolved":    data.IsResolved,
		}).Info("✅ Parsed Message")

		if data.IsNewSession {
			err := redisSvc.Enqueue("new_session_queue", data)
			if err != nil {
				log.WithError(err).Error("❌ Failed to enqueue to Redis")
			} else {
				log.Info("📤 Data enqueued to Redis successfully")
			}
		}

		response.Success(w, "Success", data)
	}
}
