package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/create", id), "application/json", bytes.NewReader(sendData))

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

func DescribeCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/describe", id), "application/json", bytes.NewReader(sendData))

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

func DropCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/drop", id), "application/json", bytes.NewReader(sendData))

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

func GetLoadState(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/get_load_state", id), "application/json", bytes.NewReader(sendData))

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

func GetStats(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/get_stats", id), "application/json", bytes.NewReader(sendData))

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

func HasCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/has", id), "application/json", bytes.NewReader(sendData))

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

func ListCollections(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/list", id), "application/json", bytes.NewReader(sendData))

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

func LoadCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/load", id), "application/json", bytes.NewReader(sendData))

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

func ReleaseCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/release", id), "application/json", bytes.NewReader(sendData))

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

func RenameCollection(c *gin.Context) {
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

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code": http.StatusNoContent,
		})
		return
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/rename", id), "application/json", bytes.NewReader(sendData))

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
