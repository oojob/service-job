package db

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/oojob/service-job/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateJob method
func (db *Database) CreateJob(in *model.Job) (string, error) {
	var inerstionID string

	jobCollection := db.mongo.Collection("job")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(context.Background())

	in.Metadata = model.MetadataModel{
		CreatedAt:     *ptypes.TimestampNow(),
		UpdatedAt:     *ptypes.TimestampNow(),
		PublishedDate: *ptypes.TimestampNow(),
		EndDate:       *ptypes.TimestampNow(),
		LastActive:    *ptypes.TimestampNow(),
	}

	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result, err := jobCollection.InsertOne(sessionContext, in)
		if err != nil {
			return "", err
		}

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			inerstionID = oid.Hex()
		}

		return "", nil
	})

	return inerstionID, err
}

// ReadJob method
func (db *Database) ReadJob(in *model.Job) (string, error) {
	return "nil", nil
}

// UpdateJob method
func (db *Database) UpdateJob(in *model.Job) (string, error) {
	return "nil", nil
}

// DeleteJob method
func (db *Database) DeleteJob(in *model.Job) (string, error) {
	return "nil", nil
}

// ReadAllJobsByCompany method
func (db *Database) ReadAllJobsByCompany(in *model.Job) (string, error) {
	return "nil", nil
}

// ReadAllJobs method
func (db *Database) ReadAllJobs(in *model.Job) (string, error) {
	return "nil", nil
}
