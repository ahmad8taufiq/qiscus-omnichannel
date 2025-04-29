package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
)

type AuthService interface {
	Nonce() (*models.NonceResponse, error)
	Login(email, password string) (*models.AuthResponse, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(email, password string) (*models.AuthResponse, error) {
	return s.repo.Authenticate(email, password)
}

func (s *authService) Nonce() (*models.NonceResponse, error) {
	return s.repo.GetNonce()
}