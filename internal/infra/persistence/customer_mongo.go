package persistence

import (
	"context"

	"github.com/jpmoraess/tracking/configs"
	"github.com/jpmoraess/tracking/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerMongo struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewCustomerMongo(client *mongo.Client) *CustomerMongo {
	return &CustomerMongo{
		client: client,
		coll:   client.Database(configs.DBNAME).Collection("customer"),
	}
}

func (c *CustomerMongo) GetByID(ctx context.Context, id string) (*entity.Customer, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entity entity.Customer
	if err := c.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&entity); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c *CustomerMongo) Insert(ctx context.Context, entity *entity.Customer) (*entity.Customer, error) {
	res, err := c.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return entity, nil
}
