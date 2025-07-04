package data

import (
	"context"
	"errors"
	"time"

	"example.com/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// services is to build up the logic related to Database

func CreateTask(task models.Task) (primitive.ObjectID, error) {
	task.CreatedAt = time.Now().UTC()
	task.UpdatedAt = time.Now().UTC()
	// comment Create a new task in the database With a timeout context of 5 seconds this means that if the operation takes longer than 5 seconds, it will be cancelled.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)	
	defer cancel()

	result, err := TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err	
	}
	defer cursor.Close(ctx)

	var tasks []models.Task

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil

}

func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func UpdateTask(id primitive.ObjectID, update models.UpdateTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		return err
	}

	updateDoc := bson.M{
		"$set": bson.M{
			"title":       update.Title,
			"description": update.Description,
			"status":      update.Status,
		},
	}

	result, err := TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, updateDoc)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no task found with the given ID")
	}

	return nil	
}

func DeleteTask(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		return err
	}

	result, err := TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}



















