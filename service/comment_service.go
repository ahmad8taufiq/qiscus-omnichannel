package service

import (
	"qiscus-omnichannel/models"
	"qiscus-omnichannel/repository"
)

type CommentService interface {
	PostComment(secret string, req *models.PostCommentRequest) (*models.PostCommentResponse, error)
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) PostComment(secret string, req *models.PostCommentRequest) (*models.PostCommentResponse, error) {
	return s.repo.PostComment(secret, req)
}
