package main

import (
	"github.com/bryandg23/owners_api/controllers/owners"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/owner", owners.CreateOwner)
	router.GET("/owner/:ownerId", owners.ReadOneOwner)
	router.GET("/appartment/:appartmentNumber/owner", owners.ReadAppartmentOwner)

	router.Run()
}
