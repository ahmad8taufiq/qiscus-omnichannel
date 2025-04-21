package app

import (
	"encoding/json"
	"net/http"
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/service"
	log "qiscus-omnichannel/tools/logger"
	"qiscus-omnichannel/tools/response"
	"strconv"
)

func SetMaxCustomerPerAgentHandler(svc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			response.NotFound(w, "Method not allowed")
			return
		}

		var payload *models.MaxCustomerPerAgentRequest
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Logger.WithError(err).Error("Invalid JSON")
			response.BadRequest(w, "Invalid JSON format")
			return
		}
		defer r.Body.Close()

		err := svc.SetCache("MAX_CUSTOMER_PER_AGENT", strconv.Itoa(payload.MaxCustomerPerAgent), 0)
		if err != nil {
			log.Logger.WithError(err).Error("Failed to set MAX_CUSTOMER_PER_AGENT")
			response.InternalServerError(w, "Failed to store value")
			return
		}

		response.Success(w, "Success", nil)
	}
}

func GetMaxCustomerPerAgentHandler(svc service.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			response.NotFound(w, "Method not allowed")
			return
		}

		maxCustomerPerAgent, err := svc.GetCache("MAX_CUSTOMER_PER_AGENT")
		if err != nil {
			log.Logger.Info("MAX_CUSTOMER_PER_AGENT not yet set")
			response.Success(w, "max_customer_per_agent not yet set", nil)
			return
		}

		value, err := strconv.Atoi(maxCustomerPerAgent)
		if err != nil {
			log.Logger.WithError(err).Error("Invalid value for MAX_CUSTOMER_PER_AGENT")
			response.InternalServerError(w, "Invalid value format")
			return
		}

		response.Success(w, "Success", map[string]interface{}{
			"maxCustomerPerAgent": value,
		})
	}
}
