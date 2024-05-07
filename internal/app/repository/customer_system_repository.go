package repository

import (
	"context"

	"github.com/jpmoraess/tracking/internal/domain/entity"
)

type CustomerSystemRepository interface {
	GetByID(context.Context, string) (*entity.CustomerSystem, error)
	Insert(context.Context, *entity.CustomerSystem) (*entity.CustomerSystem, error)
}
