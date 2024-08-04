package main

import (
	"billion-dollar-project/router/server/control"
	"billion-dollar-project/router/server/data"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Collections
	router.POST("/collections/create", control.CreateCollection)
	router.POST("/collections/drop", control.DropCollection)
	router.POST("/collections/describe", control.DescribeCollection)
	router.POST("/collections/list", control.ListCollections)
	router.POST("/collections/get_load_state", control.GetLoadState)
	router.POST("/collections/get_stats", control.GetStats)
	router.POST("/collections/load", control.LoadCollection)
	router.POST("/collections/rename", control.RenameCollection)
	router.POST("/collections/has", control.HasCollection)
	router.POST("/collections/release", control.ReleaseCollection)

	// Vectors
	router.POST("/vectors/get", data.GetWithID)
	router.POST("/vectors/insert", data.InsertVectors)
	router.POST("/vectors/delete", data.DeleteVectors)
	router.POST("/vectors/upsert", data.UpsertVectors)
	router.POST("/vectors/search", data.Search)
	router.POST("/vectors/query", data.Query)
	router.POST("/vectors/hybrid_search", data.HybridSearch)

	// Indexes
	router.POST("/indexes/create", control.CreateIndex)
	router.POST("/indexes/drop", control.DropIndex)
	router.POST("/indexes/describe", control.DescribeIndex)
	router.POST("/indexes/list", control.ListIndexes)

	// For testing purposes. Remove in production
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
