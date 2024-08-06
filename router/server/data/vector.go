package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWithID(c *gin.Context) {
	var newData map[string]any

	if err := c.BindJSON(&newData); err != nil {
		log.Println("Could not bind data, " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "JSON data could not be binded",
		})
		return
	}
	log.Println("Data binding successfull, marshalling data")
	// Do something with api key and host
	// api_key := newData["api_key"].(string)
	host := newData["host"]

	delete(newData, "api_key")
	delete(newData, "host")

	sendData, err := json.Marshal(newData)

	if err != nil {
		log.Println("Could not marshal data, " + err.Error())
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": "Content does not exist",
		})
		return
	}
	log.Println("Data marshalling successfull, sending request")
	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/entities/get", host), "application/json", bytes.NewReader(sendData))

	if err != nil {
		log.Println("Request failed, " + err.Error())
		c.JSON(resp.StatusCode, gin.H{
			"code":    resp.StatusCode,
			"message": err.Error(),
		})
		return
	}
	log.Println("Request successfull, reading response")
	buf, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Could not read response, " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Response could not be read",
		})
		return
	}
	log.Println("Response read successfull, unmarshalling response")
	var result map[string]any
	err = json.Unmarshal(buf, &result)
	if err != nil {
		log.Println("Could not unmarshal response, " + err.Error())
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": "Response could not be binded to JSON",
		})
		return
	}
	log.Println("Response unmarshalling successfull, sending response")
	code := int(result["code"].(float64))
	delete(result, "code")
	if code < 200 {
		code += 1000
	}
	c.JSON(code, result)
	resp.Body.Close()
}

func InsertVectors(c *gin.Context) {
	var newData map[string]any

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

func DeleteVectors(c *gin.Context) {
	var newData map[string]any

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

func UpsertVectors(c *gin.Context) {
	var newData map[string]any

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

func Query(c *gin.Context) {
	var newData map[string]any

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

func Search(c *gin.Context) {
	var newData map[string]any

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

func HybridSearch(c *gin.Context) {
	var newData map[string]any

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
