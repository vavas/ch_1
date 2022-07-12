package service

import "github.com/vavas/challenge/pkg/repository"

type EmailService struct {
	repo repository.Email
}

func NewEmailService(repo repository.Email) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) Create(email string) error {
	return s.repo.Create(email)
}

func (s *EmailService) GetAll() ([]string, error) {
	return s.repo.GetAll()
}

func (s *EmailService) Delete(email string) error {
	return s.repo.Delete(email)
}
