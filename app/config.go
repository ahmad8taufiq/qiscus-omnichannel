package app

import (
	"encoding/json"
	"net/http"
	"qiscus-omnichannel/config"
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	log "qiscus-omnichannel/tools/logger"
	"qiscus-omnichannel/tools/response"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
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

func GetCredentials(redisSvc service.RedisService) (adminToken, sdkEmail, sdkToken string, err error) {
	cachedAdminToken, err := redisSvc.GetCache(config.AppConfig.AdminToken)
	if err == nil && cachedAdminToken != "" {
		log.Logger.Infof("✅ Cached Admin Token: %s", cachedAdminToken)

		cachedSdkEmail, _ := redisSvc.GetCache(config.AppConfig.SdkEmail)
		cachedSdkToken, _ := redisSvc.GetCache(config.AppConfig.SdkToken)

		return cachedAdminToken, cachedSdkEmail, cachedSdkToken, nil
	}

	authRepo := repository.NewAuthRepository()
	authSvc := service.NewAuthService(authRepo)

	authResp, err := authSvc.Login(config.AppConfig.QiscusEmail, config.AppConfig.QiscusPassword)
	if err != nil {
		log.Logger.WithError(err).Error("❌ Login failed")
		return "", "", "", err
	}

	log.Logger.Infof("authResp: %v", authResp)

	adminToken = authResp.Data.User.AuthenticationToken
	sdkEmail = authResp.Data.User.SdkEmail
	sdkToken = authResp.Data.Details.SdkUser.Token

	ttl := time.Hour * 24 * 30
	cacheItems := map[string]string{
		config.AppConfig.AdminToken: adminToken,
		config.AppConfig.SdkEmail:   sdkEmail,
		config.AppConfig.SdkToken:   sdkToken,
	}

	for key, value := range cacheItems {
		if err := redisSvc.SetCache(key, value, ttl); err != nil {
			log.Logger.WithFields(logrus.Fields{
				"key":   key,
				"value": value,
			}).WithError(err).Error("❌ Failed to set credential cache")
		} else {
			log.Logger.Infof("✅ Cached %s: %s", key, value)
		}
	}

	log.Logger.Infof("✅ Admin Token (from login): %s", adminToken)
	return adminToken, sdkEmail, sdkToken, nil
}

func GetNonce(redisSvc service.RedisService) (nonce string, err error) {
	cachedNonce, err := redisSvc.GetCache(config.AppConfig.Nonce)
	if err == nil && cachedNonce != "" {
		log.Logger.Infof("✅ Cached Nonce: %s", cachedNonce)

		return cachedNonce, nil
	}

	authRepo := repository.NewAuthRepository()
	authSvc := service.NewAuthService(authRepo)

	nonceResp, err := authSvc.Nonce()
	if err != nil {
		log.Logger.WithError(err).Error("❌ Get Nonce failed")
		return "", err
	}

	log.Logger.Infof("nonceResp: %v", nonceResp)

	ttl := time.Minute * 10
	if err := redisSvc.SetCache(config.AppConfig.Nonce, nonceResp.Results.Nonce, ttl); err != nil {
		log.Logger.WithFields(logrus.Fields{
			"key":   config.AppConfig.Nonce,
			"value": nonceResp.Results.Nonce,
		}).WithError(err).Error("❌ Failed to set nonce cache")
	}

	log.Logger.Infof("✅ Nonce (from qiscus): %s", nonceResp.Results.Nonce)
	return nonceResp.Results.Nonce, nil
}
