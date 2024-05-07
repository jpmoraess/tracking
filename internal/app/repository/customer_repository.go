package repository

import (
	"context"

	"github.com/jpmoraess/tracking/internal/domain/entity"
)

type CustomerRepository interface {
	GetByID(context.Context, string) (*entity.Customer, error)
	Insert(context.Context, *entity.Customer) (*entity.Customer, error)
}
