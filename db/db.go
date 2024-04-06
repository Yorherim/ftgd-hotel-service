package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const DatabaseName = "hotel-reservation"

func ToObjectID(id string) (*primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return &oid, nil
}
