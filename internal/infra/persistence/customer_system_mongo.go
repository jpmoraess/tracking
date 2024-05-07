package persistence

import (
	"context"

	"github.com/jpmoraess/tracking/configs"
	"github.com/jpmoraess/tracking/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerSystemMongo struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewCustomerSystemMongo(client *mongo.Client) *CustomerSystemMongo {
	return &CustomerSystemMongo{
		client: client,
		coll:   client.Database(configs.DBNAME).Collection("customer_system"),
	}
}

func (c *CustomerSystemMongo) GetByID(ctx context.Context, id string) (*entity.CustomerSystem, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entity entity.CustomerSystem
	if err := c.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c *CustomerSystemMongo) Insert(ctx context.Context, entity *entity.CustomerSystem) (*entity.CustomerSystem, error) {
	res, err := c.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
