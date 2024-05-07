package persistence

import (
	"context"

	"github.com/jpmoraess/tracking/configs"
	"github.com/jpmoraess/tracking/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerIntegrationMongo struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewCustomerIntegrationMongo(client *mongo.Client) *CustomerIntegrationMongo {
	return &CustomerIntegrationMongo{
		client: client,
		coll:   client.Database(configs.DBNAME).Collection("customer_integration"),
	}
}

func (c *CustomerIntegrationMongo) Insert(ctx context.Context, entity *entity.CustomerIntegration) (*entity.CustomerIntegration, error) {
	res, err := c.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
