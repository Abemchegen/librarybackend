package repository

import (
	"context"
	"librarybackend/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type recordRepository struct {
	database   mongo.Database
	collection string
}

func NewRecordRepository(db mongo.Database, collection string) domain.RecordRepository {
	return &recordRepository{
		database:   db,
		collection: collection}
}

func (r *recordRepository) CreateRecord(student domain.Student, book domain.Book, lentdate string, duedate string, lenttype string, returnstatus string, returncondition string) (domain.Record, error) {
	record := domain.Record{
		Student:         student,
		Book:            book,
		LentDate:        lentdate,
		ReturnDue:       duedate,
		LentType:        lenttype,
		ReturnStatus:    returnstatus,
		ReturnCondition: returncondition,
	}

	_, err := r.database.Collection(r.collection).InsertOne(context.TODO(), record)
	if err != nil {
		return domain.Record{}, err
	}
	return record, nil
}

func (r *recordRepository) GetAllRecord() ([]domain.Record, error) {
	var records []domain.Record
	cursor, err := r.database.Collection(r.collection).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var record domain.Record
		err := cursor.Decode(&record)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *recordRepository) UpdateRecord(studentid string, bookid string, returnstatus string, returncondition string) (domain.Record, error) {
	record := domain.Record{
		ReturnStatus:    returnstatus,
		ReturnCondition: returncondition,
	}

	filter := bson.M{"student._SchoolID": studentid, "book._Bookid": bookid}
	update := bson.M{
		"$set": record,
	}
	_, err := r.database.Collection(r.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.Record{}, err
	}
	return record, nil
}
