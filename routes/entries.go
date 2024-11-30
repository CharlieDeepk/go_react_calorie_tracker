package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/CharlieDeepk/go_react_calorie_tracker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var entryCollection *mongo.Collection = openCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entry models.Entry

	if bindErr := c.BindJSON(&entry); bindErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Bind error": bindErr})
		return
	}
	if validateErr := validator.New().Struct(entry); validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Validator error": validateErr})
		return
	}
	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"Insert error": insertErr})
	}
	c.JSON(http.StatusOK, result)

}
func GetEntries(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entries []bson.M
	cursor, findErr := entryCollection.Find(ctx, bson.M{})
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Find error": findErr.Error()})
		return
	}
	if curErr := cursor.All(ctx, &entries); curErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"Cursor Error": curErr.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)

}
func GetEntryById(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, hexErr := primitive.ObjectIDFromHex(entryId)
	if hexErr != nil {
		log.Printf("Hex Error: %v", hexErr)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entry bson.M
	if findErr := entryCollection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry); findErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"Find error": findErr.Error()})
		return
	}
	c.JSON(http.StatusOK, entry)

}
func GetEntriesByIngredient(c *gin.Context) {
	ingredient := c.Params.ByName("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var entries []bson.M
	cursor, findErr := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Find error": findErr.Error()})
		return
	}
	if curErr := cursor.All(ctx, &entries); curErr != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"Cursor Error": curErr.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)

}
func UpdateEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, hexErr := primitive.ObjectIDFromHex(entryId)
	if hexErr != nil {
		log.Printf("Hex Error: %v", hexErr)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entry models.Entry
	if bindErr := c.BindJSON(&entry); bindErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Bind error": bindErr})
		return
	}
	if validateErr := validator.New().Struct(entry); validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Validator error": validateErr})
		return
	}
	result, updateErr := entryCollection.ReplaceOne(ctx, bson.M{"_id": docId}, bson.M{"dish": entry.Dish, "fats": entry.Fats, "ingredients": entry.Ingredients, "calories": entry.Calories})
	if updateErr != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"Insert error": updateErr})
	}
	c.JSON(http.StatusOK, result.ModifiedCount)

}
func UpdateIngredient(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, hexErr := primitive.ObjectIDFromHex(entryId)
	if hexErr != nil {
		log.Printf("Hex Error: %v", hexErr)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient
	if bindErr := c.BindJSON(&ingredient); bindErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Bind error": bindErr})
		return
	}

	result, updateErr := entryCollection.UpdateOne(ctx, bson.M{"_id": docId}, bson.D{{Key: "$set", Value: bson.D{{Key: "ingredients", Value: ingredient.Ingredients}}}})
	if updateErr != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"Insert error": updateErr})
	}
	c.JSON(http.StatusOK, result.ModifiedCount)

}
func DeleteEntry(c *gin.Context) {
	entryId := c.Params.ByName("id")
	docId, hexErr := primitive.ObjectIDFromHex(entryId)
	if hexErr != nil {
		log.Printf("Hex Error: %v", hexErr)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	result, delErr := entryCollection.DeleteOne(ctx, bson.M{"_id": docId})
	if delErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Delete error": delErr.Error()})
		return
	}
	c.JSON(http.StatusOK, result.DeletedCount)
}
