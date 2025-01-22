package repository

import (
	"context"
	"errors"
	"librarybackend/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	database   mongo.Database
	collection string
}

func NewStudentRepository(database mongo.Database, collection string) domain.StudentRepository {
	return &StudentRepository{
		database:   database,
		collection: collection,
	}
}

func (ur *StudentRepository) EnterLibrary(student domain.Student) (domain.Student, error) {
	ctx := context.Background()

	filter := bson.M{"studentid": student.SchoolID, "leavetime": nil}
	var existingActivity domain.Activity
	err := ur.database.Collection(ur.collection).FindOne(ctx, filter).Decode(&existingActivity)
	if err == nil {
		return domain.Student{}, errors.New("student has already entered the library")
	}

	activity := domain.Activity{
		ID:          primitive.NewObjectID(),
		StudentName: student.Name,
		StudentID:   student.SchoolID,
		EntryTime:   time.Now(),
		LeaveTime:   nil,
	}

	_, err = ur.database.Collection(ur.collection).InsertOne(ctx, activity)
	if err != nil {
		return domain.Student{}, err
	}

	return student, nil
}

func (ur *StudentRepository) LeaveLibrary(student domain.Student) error {
	ctx := context.Background()
	now := time.Now()

	filter := bson.M{"studentid": student.SchoolID, "leavetime": nil}
	update := bson.M{"$set": bson.M{"leavetime": &now}}

	result, err := ur.database.Collection(ur.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("student did not enter the library")
	}

	return nil
}

// GetAllActivities retrieves all activities in the database.
func (ur *StudentRepository) GetStudentActivity() ([]domain.Activity, error) {
	ctx := context.Background()
	var activities []domain.Activity

	cursor, err := ur.database.Collection(ur.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var activity domain.Activity
		if err := cursor.Decode(&activity); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return activities, nil
}

func (ur *StudentRepository) GetUniqueStudentCountPerDay() (map[string]int, error) {
	ctx := context.Background()

	collection := ur.database.Collection(ur.collection)

	// MongoDB aggregation pipeline
	pipeline := mongo.Pipeline{
		// Step 1: Add a date-only field from `entrytime`
		{
			{
				Key: "$addFields", Value: bson.D{
					{Key: "date", Value: bson.M{
						"$dateToString": bson.M{
							"format": "%Y-%m-%d", // Format as YYYY-MM-DD
							"date":   "$entrytime",
						},
					}},
				},
			},
		},
		// Step 2: Group by date and student ID
		{
			{
				Key: "$group", Value: bson.D{
					{Key: "_id", Value: bson.D{
						{Key: "date", Value: "$date"},
						{Key: "studentid", Value: "$studentid"},
					}},
				},
			},
		},
		// Step 3: Group by date again to count unique students
		{
			{
				Key: "$group", Value: bson.D{
					{Key: "_id", Value: "$_id.date"},
					{Key: "uniqueStudents", Value: bson.D{
						{Key: "$sum", Value: 1},
					}},
				},
			},
		},
	}

	// Execute the aggregation
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Parse the results
	results := make(map[string]int)
	for cursor.Next(ctx) {
		var result struct {
			Date           string `bson:"_id"`
			UniqueStudents int    `bson:"uniqueStudents"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results[result.Date] = result.UniqueStudents
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (ur *StudentRepository) GetCurrentVisitors() ([]domain.Activity, error) {
	ctx := context.Background()
	collection := ur.database.Collection(ur.collection)

	// Filter: leavetime is nil
	filter := bson.M{"leavetime": nil}

	// Find all documents matching the filter
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Parse the results
	var currentVisitors []domain.Activity
	for cursor.Next(ctx) {
		var activity domain.Activity
		if err := cursor.Decode(&activity); err != nil {
			return nil, err
		}
		currentVisitors = append(currentVisitors, activity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return currentVisitors, nil
}
