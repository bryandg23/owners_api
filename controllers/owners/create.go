package owners

import (
	"context"
	"fmt"
	"net/http"
	"time"

	collection "github.com/bryandg23/owners_api/collection"
	database "github.com/bryandg23/owners_api/databases"
	model "github.com/bryandg23/owners_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func CreateOwner(c *gin.Context) {
	var databaseClient = database.ConnectDB()
	var ownersCollection = collection.GetCollection(databaseClient, "Owners")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	owner := new(model.Owner)
	defer cancel()

	if err := c.BindJSON(owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println(err)
		return
	}

	ownerPayload := model.Owner{
		Id:               primitive.NewObjectID(),
		FirstName:        owner.FirstName,
		LastName:         owner.LastName,
		Email:            owner.Email,
		Birthday:         owner.Birthday,
		AppartmentNumber: owner.AppartmentNumber,
		Contract:         owner.Contract,
		PhoneNumber:      owner.PhoneNumber,
	}

	_, err_exists := GetOwnerByFilter(bson.D{{Key: "appartmentnumber", Value: owner.AppartmentNumber}})

	if err_exists == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "An owner was already assigned for the indicated appartment"})
		return
	}

	result, err := ownersCollection.InsertOne(ctx, ownerPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Added successfully",
		"Data":    map[string]interface{}{"data": result},
	})
}
