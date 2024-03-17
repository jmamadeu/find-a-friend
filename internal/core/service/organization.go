package service

import (
	"context"

	"fiend-a-friend/fiend-a-friend-api/internal/core/domain"
	"fiend-a-friend/fiend-a-friend-api/internal/port"
	"fiend-a-friend/fiend-a-friend-api/internal/utils"

	"github.com/google/uuid"
)

type Organization struct {
	organizationRepository port.OrganizationRepository
}

func NewOrganization(orgRepo port.OrganizationRepository) *Organization {
	return &Organization{
		orgRepo,
	}
}

func (service *Organization) CreateOrganization(ctx context.Context, params port.CreateOrganizationParams) (organization *domain.Organization, err error) {
	passwordHash, _ := utils.HashPassword(params.Password)

	organization = &domain.Organization{
		Id:           uuid.New(),
		Address:      params.Address,
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: passwordHash,
	}

	err = service.organizationRepository.Create(ctx, organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}
