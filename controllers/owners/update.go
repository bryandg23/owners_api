package owners

import (
	"context"
	"net/http"
	"time"

	getcollection "github.com/bryandg23/owners_api/collection"
	"github.com/bryandg23/owners_api/databases"
	model "github.com/bryandg23/owners_api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateOwnerById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	databaseClient := databases.ConnectDB()
	ownerCollection := getcollection.GetCollection(databaseClient, "Owners")

	ownerId := c.Param("ownerId")
	var owner model.Owner

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ownerId)

	if err := c.BindJSON(&owner); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	edited := bson.M{
		"firstname":        owner.FirstName,
		"lastname":         owner.LastName,
		"email":            owner.Email,
		"birthday":         owner.Birthday,
		"appartmentnumber": owner.AppartmentNumber,
		"contract":         owner.Contract,
		"phonenumber":      owner.PhoneNumber,
	}

	result, err := ownerCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data updated successfully", "data": result})
}
