package repository

import (
	"context"

	"github.com/jpmoraess/tracking/internal/domain/entity"
)

type ShippingCompanyRepository interface {
	GetByID(context.Context, string) (*entity.ShippingCompany, error)
	Insert(context.Context, *entity.ShippingCompany) (*entity.ShippingCompany, error)
}
