package service

import "github.com/vavas/challenge/pkg/repository"

type Email interface {
	Create(email string) error
	GetAll() ([]string, error)
	Delete(email string) error
}

type Service struct {
	Email
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Email: NewEmailService(repos.Email),
	}
}
