package owners

import (
	"context"
	"net/http"
	"strconv"
	"time"

	getcollection "github.com/bryandg23/owners_api/collection"
	database "github.com/bryandg23/owners_api/databases"
	model "github.com/bryandg23/owners_api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadOneOwner(c *gin.Context) {
	ownerId := c.Param("ownerId")

	objId, _ := primitive.ObjectIDFromHex(ownerId)

	res, err := GetOwnerByFilter(bson.D{{Key: "id", Value: objId}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "success", "data": res})
}

func ReadAppartmentOwner(c *gin.Context) {
	appartmentNumber, _ := strconv.Atoi(c.Param("appartmentNumber"))

	res, err := GetOwnerByFilter(bson.D{{Key: "appartmentnumber", Value: appartmentNumber}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "success", "data": res})
}

func GetOwnerByFilter(filter bson.D) (model.Owner, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	databaseClient := database.ConnectDB()
	ownerCollection := getcollection.GetCollection(databaseClient, "Owners")
	var result model.Owner
	defer cancel()

	err := ownerCollection.FindOne(ctx, filter).Decode(&result)

	return result, err
}
