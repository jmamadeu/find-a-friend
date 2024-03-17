package port

import (
	"context"

	"find-a-friend/find-a-friend-api/internal/core/domain"
)

type OrganizationRepository interface {
	Create(ctx context.Context, params *domain.Organization) error
}

type CreateOrganizationParams struct {
	Address  string
	Email    string
	Name     string
	Password string
}
