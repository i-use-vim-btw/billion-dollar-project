package main

import (
	"billion-dollar-project/router/server/control"
	"billion-dollar-project/router/server/data"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Collections
	router.POST("/gandi/collections/create", control.CreateCollection)

	// Vectors
	router.POST("/gandi/vectors/get", data.GetWithID)
	router.POST("/gandi/vectors/insert", data.InsertVectors)
	router.POST("/gandi/vectors/delete", data.DeleteVectors)
	router.POST("/gandi/vectors/upsert", data.UpsertVectors)

	// Indexes
	router.POST("/gandi/indexes/create", control.CreateIndex)
	router.POST("/gandi/indexes/drop", control.DropIndex)
	router.POST("/gandi/indexes/describe", control.DescribeIndex)
	router.POST("/gandi/indexes/list", control.ListIndexes)

	router.Run("localhost:8080")
}
