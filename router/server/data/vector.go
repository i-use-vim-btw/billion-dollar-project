package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWithID(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/get", host), "application/json", bytes.NewReader(sendData))

	if err != nil {
		c.JSON(resp.StatusCode, gin.H{
			"code": resp.StatusCode,
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
			"code":    http.StatusNotAcceptable,
			"message": err.Error(),
		})
		return
	}

	c.JSON(resp.StatusCode, result)
	resp.Body.Close()
}

func InsertVectors(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/insert", host), "application/json", bytes.NewReader(sendData))
	fmt.Println(resp.StatusCode)
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

func DeleteVectors(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/delete", host), "application/json", bytes.NewReader(sendData))

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

func UpsertVectors(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/upsert", host), "application/json", bytes.NewReader(sendData))

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

func Query(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/query", host), "application/json", bytes.NewReader(sendData))

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

func Search(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/search", host), "application/json", bytes.NewReader(sendData))

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

func HybridSearch(c *gin.Context) {
	var newData map[string]interface{}

	if err := c.BindJSON(&newData); err != nil {
		fmt.Println("Could not bind data")
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": http.StatusNotAcceptable,
		})
		return
	}

	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/search", host), "application/json", bytes.NewReader(sendData))

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
