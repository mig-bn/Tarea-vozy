package user_repositor

import (
	"Tarea-vozy/database"
	m "Tarea-vozy/models"
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user m.User) error {

	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil

}

func Read() (m.Users, error) {
	/*var user m.User
	var users m.Users

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		defer cur.Close(ctx)
		return nil, err
	}

	for cur.Next(ctx) {
		err = cur.Decode(&user)

		if err != nil {
			return users, err
		}

		users = append(users, &user)
		fmt.Println(users)
	}

	return users, nil*/
	var users m.Users
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user m.User
		err = cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		users = append(users, &user)
	}
	return users, nil
}

func Update(user m.User, userID string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_At": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil
	}

	return nil

}

func Delete(userID string) error {

	var err error

	oid, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil

}
