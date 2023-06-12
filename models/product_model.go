package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/database"
	"github.com/aabdullahgungor/go-restapi-rabbitmq/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrProductIDIsNotValid   = errors.New("id is not valid")
	ErrProductNameIsNotEmpty = errors.New("product name cannot be empty")
	ErrProductNotFound       = errors.New("the product cannot be found")
)

type ProductModel struct {
}

func (productModel ProductModel) GetAll() ([]entities.Product, error) {
	productCollection := database.GetDB().Collection("product")

	var products []entities.Product
	productCursor, err := productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	if err = productCursor.All(context.TODO(), &products); err != nil {
		panic(err)
	}

	return products, err
}

func (productModel ProductModel) GetById(id string) (entities.Product, error) {

	productCollection := database.GetDB().Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)

	var product entities.Product
	err := productCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&product)
	if err != nil {
		return entities.Product{}, ErrProductNotFound
	}

	return product, nil

}

func (productModel ProductModel) Create(product *entities.Product) error {
	productCollection := database.GetDB().Collection("product")

	result, err := productCollection.InsertOne(context.TODO(), product)

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", result.InsertedID)

	return err
}

func (productModel ProductModel) Edit(product *entities.Product) error {
	productCollection := database.GetDB().Collection("product")

	bsonBytes, err := bson.Marshal(&product)

	if err != nil {
		panic(fmt.Errorf("can't marshal:%s", err))
	}

	var upt bson.M
	bson.Unmarshal(bsonBytes, &upt)

	update := bson.M{"$set": upt}

	result, err := productCollection.UpdateOne(context.TODO(), bson.M{"_id": product.Id}, update)

	if err != nil {
		panic(err)
	}

	log.Println("Number of documents updated:" + strconv.Itoa(int(result.ModifiedCount)))

	return err
}

func (productModel ProductModel) Delete(id string) error {
	productCollection := database.GetDB().Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})

	// check for errors in the deleting
	if err != nil {
		panic(err)
	}

	// display the number of documents deleted
	log.Println("deleting the first result from the search filter\n" + "Number of documents deleted:" + strconv.Itoa(int(result.DeletedCount)))

	return err
}
