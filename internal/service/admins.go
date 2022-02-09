package service

import (
	"context"

	"github.com/sigit14ap/go-commerce/internal/domain"
	"github.com/sigit14ap/go-commerce/internal/domain/dto"
	"github.com/sigit14ap/go-commerce/internal/repository"
)

type AdminsService struct {
	repo repository.Admins
}

func (a *AdminsService) FindByCredentials(ctx context.Context, signInDTO dto.SignInDTO) (domain.Admin, error) {
	return a.repo.FindByCredentials(ctx, signInDTO.Email, signInDTO.Password)
}

func NewAdminsService(repo repository.Admins) *AdminsService {
	return &AdminsService{
		repo: repo,
	}
}
