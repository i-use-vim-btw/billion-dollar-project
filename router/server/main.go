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
	router.POST("/gandi/collections/drop", control.DropCollection)
	router.POST("/gandi/collections/describe", control.DescribeCollection)
	router.POST("/gandi/collections/list", control.ListCollections)
	router.POST("/gandi/collections/get_load_state", control.GetLoadState)
	router.POST("/gandi/collections/get_stats", control.GetStats)
	router.POST("/gandi/collections/load", control.LoadCollection)
	router.POST("/gandi/collections/rename", control.RenameCollection)
	router.POST("/gandi/collections/has", control.HasCollection)
	router.POST("/gandi/collections/release", control.ReleaseCollection)

	// Vectors
	router.POST("/gandi/vectors/get", data.GetWithID)
	router.POST("/gandi/vectors/insert", data.InsertVectors)
	router.POST("/gandi/vectors/delete", data.DeleteVectors)
	router.POST("/gandi/vectors/upsert", data.UpsertVectors)
	router.POST("/gandi/vectors/search", data.Search)
	router.POST("/gandi/vectors/query", data.Query)

	// Indexes
	router.POST("/gandi/indexes/create", control.CreateIndex)
	router.POST("/gandi/indexes/drop", control.DropIndex)
	router.POST("/gandi/indexes/describe", control.DescribeIndex)
	router.POST("/gandi/indexes/list", control.ListIndexes)

	router.Run("localhost:8080")
}
