package repository

import (
	"context"

	"github.com/jpmoraess/tracking/internal/domain/entity"
)

type CustomerIntegrationRepository interface {
	Insert(context.Context, *entity.CustomerIntegration) (*entity.CustomerIntegration, error)
}
