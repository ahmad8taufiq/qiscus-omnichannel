package app

import (
	"io"
	"net/http"

	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/response"

	"github.com/sirupsen/logrus"
)

func WebhookHandler(log *logrus.Logger, svc service.WebhookService) http.HandlerFunc {
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

		data, err := svc.ProcessWebhook(body)
		if err != nil {
			log.WithError(err).Error("Failed to process webhook")
			response.BadRequest(w, "Invalid JSON format")
			return
		}

		log.WithFields(logrus.Fields{
			"email":          data.Email,
			"is_new_session": data.IsNewSession,
			"is_resolved":    data.IsResolved,
			"latest_service": data.LatestService,
			"name":           data.Name,
			"room_id":        data.RoomId,
		}).Info("✅ Parsed Message")

		if data.CandidateAgent != nil {
			log.WithFields(logrus.Fields{
				"id":             data.CandidateAgent.ID,
				"email":          data.CandidateAgent.Email,
				"name":           data.CandidateAgent.Name,
				"type":           data.CandidateAgent.Type,
				"type_as_string": data.CandidateAgent.TypeAsString,
				"is_available":   data.CandidateAgent.IsAvailable,
			}).Info("🧑‍💼 Candidate Agent Details")
		} else {
			log.Info("⚠️ CandidateAgent is nil")
		}

		response.Success(w, "Success", data)
	}
}
