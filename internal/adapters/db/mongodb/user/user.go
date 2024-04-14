package user

import (
	"context"
	"github.com/Yorherim/ftgd-hotel-service/internal/controller/http/v1/user"
	"github.com/Yorherim/ftgd-hotel-service/internal/domain/entity"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const DatabaseName = "hotel-reservation"

const UserColl = "users"

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	logger *zap.SugaredLogger
}

func NewMongoUserStore(client *mongo.Client, logger *zap.SugaredLogger) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DatabaseName).Collection(UserColl),
		logger: logger,
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	s.logger.Info("db GetUserByID start")

	var user entity.User

	objId, err := ToObjectID(id)
	if err != nil {
		s.logger.Errorf("db GetUserByID error ToObjectID: %s", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := s.coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&user); err != nil {
		s.logger.Errorf("db GetUserByID error FindOne: %s", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*entity.User, error) {
	s.logger.Info("db GetUsers start")

	var users []*entity.User

	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		s.logger.Errorf("db GetUsers error Find: %s", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "server error")
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user *entity.User
		if err := cur.Decode(&user); err != nil {
			s.logger.Errorf("db GetUsers error Decode: %s", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "server error")
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, dto user.CreateUserDTO) (*entity.User, error) {
	s.logger.Info("db CreateUser start")

	res, err := s.coll.InsertOne(ctx, dto)
	if err != nil {
		s.logger.Errorf("db CreateUser error InsertOne: %s", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "error create user")
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if ok {
		return &entity.User{
			Password:  dto.Password,
			LastName:  dto.LastName,
			Email:     dto.Email,
			FirstName: dto.FirstName,
			ID:        id.Hex(),
		}, nil
	}

	s.logger.Error("db CreateUser error InsertedID")
	return nil, fiber.NewError(fiber.StatusInternalServerError, "error server")
}

func ToObjectID(id string) (*primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return &oid, nil
}
