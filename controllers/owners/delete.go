package owners

import (
	"context"
	"net/http"
	"time"

	getcollection "github.com/bryandg23/owners_api/collection"
	"github.com/bryandg23/owners_api/databases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOwnerById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	databaseClient := databases.ConnectDB()

	ownerId := c.Param("ownerId")

	ownerCollection := getcollection.GetCollection(databaseClient, "Owners")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ownerId)

	result, err := ownerCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Owner deleted successfully", "data": result})
}
