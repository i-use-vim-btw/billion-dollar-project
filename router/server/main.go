package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/gandi/collections/create", createCollection)
	router.POST("/gandi/entities/get", getWithID)
	router.POST("/gandi/entities/insert", insertVector)
	router.POST("/gandi/entities/delete", deleteVectors)
	router.POST("/gandi/entities/upsert", upsertVector)

	router.Run("localhost:8080")
}

func createCollection(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and id
	// api_key := newData["api_key"].(string)
	id := newData["id"]

	delete(newData, "api_key")
	delete(newData, "id")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/create", id), "application/json", bytes.NewBuffer(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"code": resp.StatusCode,
		"data": resp.Body,
	})
}

func getWithID(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and id
	// api_key := newData["api_key"].(string)
	id := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	fmt.Println(string(sendData))

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/get", id), "application/json", bytes.NewBuffer(sendData))

	if err != nil {
		c.JSON(resp.StatusCode, gin.H{
			"code": resp.StatusCode,
		})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"code": resp.StatusCode,
		"data": resp.Body,
	})
}

func insertVector(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and id
	// api_key := newData["api_key"].(string)
	id := newData["id"]

	delete(newData, "api_key")
	delete(newData, "id")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/insert", id), "application/json", bytes.NewBuffer(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"code": resp.StatusCode,
		"data": resp.Body,
	})
}

type DeleteData struct {
	DatabaseName   string
	CollectionName string
	Filter         string
	PartitionName  string
}

func deleteVectors(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and id
	// api_key := newData["api_key"].(string)
	id := newData["id"]

	delete(newData, "api_key")
	delete(newData, "id")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/delete", id), "application/json", bytes.NewBuffer(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"code": resp.StatusCode,
		"data": resp.Body,
	})

}

func upsertVector(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and id
	// api_key := newData["api_key"].(string)
	id := newData["id"]

	delete(newData, "api_key")
	delete(newData, "id")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/upsert", id), "application/json", bytes.NewBuffer(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"code": resp.StatusCode,
		"data": resp.Body,
	})
}
