package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	log "qiscus-omnichannel/tools/logger"
	"qiscus-omnichannel/tools/response"
	"strconv"
	"time"
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

func GetAdminToken(redisSvc service.RedisService) (string, error) {
	cachedAdminToken, err := redisSvc.GetCache(config.AppConfig.AdminTokenKey)
	if err == nil && cachedAdminToken != "" {
		log.Logger.Infof("✅ Cached Admin Token: %s", cachedAdminToken)
		return cachedAdminToken, nil
	}

	authRepo := repository.NewAuthRepository(fmt.Sprintf("%s/api/v1/auth", config.AppConfig.QiscusBaseURL))
	authSvc := service.NewAuthService(authRepo)

	authResp, err := authSvc.Login(config.AppConfig.QiscusEmail, config.AppConfig.QiscusPassword)
	if err != nil {
		log.Logger.WithError(err).Error("❌ Login failed")
		return "", err
	}

	adminToken := authResp.Data.User.AuthenticationToken

	ttl := time.Hour * 24 * 30
	if err := redisSvc.SetCache(config.AppConfig.AdminTokenKey, adminToken, ttl); err != nil {
		log.Logger.WithError(err).Error("❌ Failed to set admin token to cache")
	}

	log.Logger.Infof("✅ Admin Token (from login): %s", adminToken)
	return adminToken, nil
}
