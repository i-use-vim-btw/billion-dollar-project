package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateIndex(c *gin.Context) {
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

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/indexes/create", id), "application/json", bytes.NewReader(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	var result map[string]any
	err = json.Unmarshal(buf, &result)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, result)
	resp.Body.Close()
}

func DescribeIndex(c *gin.Context) {
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

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/indexes/describe", id), "application/json", bytes.NewReader(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	var result map[string]any
	err = json.Unmarshal(buf, &result)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, result)
	resp.Body.Close()
}

func DropIndex(c *gin.Context) {
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

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/indexes/drop", id), "application/json", bytes.NewReader(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, result)
	resp.Body.Close()
}

func ListIndexes(c *gin.Context) {
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

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/indexes/list", id), "application/json", bytes.NewReader(sendData))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(buf, &result)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	c.JSON(resp.StatusCode, result)
	resp.Body.Close()
}
