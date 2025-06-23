package app

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"qiscus-omnichannel/models"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/logger"
	tools "qiscus-omnichannel/tools/parser"
	"qiscus-omnichannel/tools/response"

	"github.com/sirupsen/logrus"
)

func WebhookResolveHandler(redisSvc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			response.NotFound(w, "Method not allowed")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Logger.WithFields(logrus.Fields{
				"method":   r.Method,
				"endpoint": r.URL.Path,
				"error":    err.Error(),
			}).Error("Failed to read request body")
			response.BadRequest(w, "Failed to read request body")
			return
		}
		defer r.Body.Close()

		logger.Logger.WithFields(logrus.Fields{
			"method":   r.Method,
			"endpoint": r.URL.Path,
			"raw":      string(body),
		}).Info("📩 Incoming request")

		data, err := tools.Parser[models.CustomerResponse](body)
		if err != nil {
			logger.Logger.WithError(err).Error("Failed to process webhook")
			response.BadRequest(w, "Invalid JSON format")
			return
		}

		var agents []models.Agents
		cachedAgents, err := redisSvc.GetCache("agents")
		
		if err == nil && cachedAgents != "" {
			err = json.Unmarshal([]byte(cachedAgents), &agents)
			if err != nil {
				logger.Logger.Info("⚠️ Failed to unmarshal cached agents, starting fresh")
				agents = []models.Agents{}
			}
		} else {
			logger.Logger.Info("📦 No cached agents found or Redis error:", err)
		}
		
		for i, agent := range agents {
			if agent.ID == data.ResolvedBy.ID {
				agents[i].CurrentCustomerCount -= 1
				break
			}
		}
		
		updatedJSON, err := json.Marshal(agents)
		if err != nil {
			logger.Logger.Error("❌ Failed to marshal updated agents:", err)
			return
		}

		err = redisSvc.SetCache("agents", string(updatedJSON), 10*time.Minute)
		if err != nil {
			logger.Logger.Error("❌ Failed to save updated agents to Redis:", err)
			return
		}

		redisSvc.SetCache("agents", string(updatedJSON), 10*time.Minute)

		logger.Logger.Infof("👤 Agents: %v", agents)

		response.Success(w, "Success", data)
	}
}
