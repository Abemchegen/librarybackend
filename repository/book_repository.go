package repository

import (
	"context"
	"librarybackend/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	database   mongo.Database
	collection string
}

func NewBookRepository(database mongo.Database, collection string) domain.BookRepository {
	return &BookRepository{
		database:   database,
		collection: collection}

}

func (p *BookRepository) CreateBook(Book domain.Book) (domain.Book, error) {
	objID := primitive.NewObjectID()
	Book.ID = objID
	_, err := p.database.Collection(p.collection).InsertOne(context.Background(), Book)
	if err != nil {
		return domain.Book{}, err
	}

	return Book, nil
}

func (p *BookRepository) GetAllBook() ([]domain.Book, error) {
	var Books []domain.Book
	cursor, err := p.database.Collection(p.collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &Books); err != nil {
		return nil, err
	}
	return Books, nil
}

func (p *BookRepository) GetBookByID(id string) (domain.Book, error) {
	var Book domain.Book
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Book{}, err
	}
	err = p.database.Collection(p.collection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&Book)
	if err != nil {
		return domain.Book{}, err
	}
	return Book, nil
}

func (p *BookRepository) UpdateBook(Book domain.Book) (domain.Book, error) {
	_, err := p.database.Collection(p.collection).UpdateOne(context.Background(), bson.M{"_id": Book.ID}, bson.M{"$set": Book})
	if err != nil {
		return domain.Book{}, err
	}
	return Book, nil
}

func (p *BookRepository) DeleteBook(id string) (domain.Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Book{}, err
	}
	var Book domain.Book
	err = p.database.Collection(p.collection).FindOneAndDelete(context.Background(), bson.M{"_id": objID}).Decode(&Book)
	if err != nil {
		return domain.Book{}, err
	}
	return Book, nil
}
func (p *BookRepository) LendBook(bookid string, studentid string, studentname string) (domain.Book, error) {
	objID, err := primitive.ObjectIDFromHex(bookid)
	if err != nil {
		return domain.Book{}, err
	}
	_, err = p.database.Collection(p.collection).UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$inc": bson.M{"Quantity": -1}},
	)
	if err != nil {
		return domain.Book{}, nil
	}

	book, err := p.GetBookByID(bookid)

	return book, err
}

func (p *BookRepository) ReturnBook(bookid string, studentid string) (domain.Book, error) {
	objID, err := primitive.ObjectIDFromHex(bookid)
	if err != nil {
		return domain.Book{}, err
	}
	_, err = p.database.Collection(p.collection).UpdateOne(
		context.Background(),
		bson.M{"_id": objID, "lent_to": studentid},
		bson.M{"$inc": bson.M{"Quantity": 1}},
	)
	if err != nil {
		return domain.Book{}, nil
	}

	book, err := p.GetBookByID(bookid)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}
