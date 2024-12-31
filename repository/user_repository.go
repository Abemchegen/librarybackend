package repository

import (
	"context"
	"fmt"
	"librarybackend/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(database mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection}

}

func (ur *userRepository) CreateAccount(user domain.User) (domain.User, error) {
	ctx := context.Background()
	user.ID = primitive.NewObjectID()

	_, err := ur.database.Collection(ur.collection).InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *userRepository) Login(user domain.User) (domain.User, error) {
	ctx := context.Background()

	filter := bson.M{"email": user.Email}
	err := ur.database.Collection(ur.collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *userRepository) GetAllUserByEmial(email string) (domain.User, error) {
	ctx := context.Background()
	var user domain.User

	filter := bson.M{"email": email}
	err := ur.database.Collection(ur.collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *userRepository) GetByID(id string) (domain.User, error) {
	ctx := context.Background()
	var user domain.User

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	err := ur.database.Collection(ur.collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *userRepository) UpdateProfile(id string, user domain.User) (domain.User, error) {
	// Get collection
	collection := repo.database.Collection(repo.collection)

	// Convert the provided id string to MongoDB ObjectID
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	// Create an empty bson.M for updates
	updateFields := bson.M{}

	// Only update fields that are not empty or zero-value
	if user.Name != "" {
		updateFields["name"] = user.Name
	}
	if user.Email != "" {
		updateFields["email"] = user.Email
	}
	if user.Password != "" {
		// Optionally, hash the password before updating
		updateFields["password"] = user.Password
	}

	// Check if there are any fields to update
	if len(updateFields) == 0 {
		return domain.User{}, fmt.Errorf("no fields to update")
	}

	// Prepare the update document
	update := bson.M{"$set": updateFields}

	// Define the updated user object
	var updatedUser domain.User

	// Find and update the document, then decode the updated user
	err = collection.FindOneAndUpdate(context.Background(), bson.M{"_id": idPrimitive}, update).Decode(&updatedUser)
	if err != nil {
		return domain.User{}, err
	}

	// get the updated_user by id

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return repo.GetUserByID(ctx, id)
}

func (ur *userRepository) GetAllUser() ([]domain.User, error) {
	ctx := context.Background()
	var users []domain.User

	// filter := bson.M{"role": "driver"}
	cursor, err := ur.database.Collection(ur.collection).Find(ctx, bson.M{})
	if err != nil {
		return []domain.User{}, err
	}

	for cursor.Next(ctx) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}

	filter := bson.M{"_id": objectID}

	err = ur.database.Collection(ur.collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	fmt.Println(user, "+++++++++++++++++++++++++++++++++++++++++++")
	return user, nil
}
