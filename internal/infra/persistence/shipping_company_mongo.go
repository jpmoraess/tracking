package persistence

import (
	"context"

	"github.com/jpmoraess/tracking/configs"
	"github.com/jpmoraess/tracking/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShippingCompanyMongo struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewShippingCompanyMongo(client *mongo.Client) *ShippingCompanyMongo {
	return &ShippingCompanyMongo{
		client: client,
		coll:   client.Database(configs.DBNAME).Collection("shipping_company"),
	}
}

func (s *ShippingCompanyMongo) GetByID(ctx context.Context, id string) (*entity.ShippingCompany, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entity entity.ShippingCompany
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (s *ShippingCompanyMongo) Insert(ctx context.Context, entity *entity.ShippingCompany) (*entity.ShippingCompany, error) {
	res, err := s.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
