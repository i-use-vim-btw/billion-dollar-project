package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Collection(c *gin.Context, method string) {
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

	if host != nil {
		delete(newData, "host")
	} else {
		host = "localhost:8080"
	}

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

	// host cannot be empty, router crashes.
	resp, err := http.Post(fmt.Sprintf("http://%s/v2/vectordb/collections/%s", host, method), "application/json", bytes.NewReader(sendData))

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
		str := string(buf[:])
		if str == "404 page not found" {
			c.JSON(404, str)
			return
		}
		log.Println("Could not unmarshal response, " + err.Error() + string(buf[:]))
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    http.StatusNotAcceptable,
			"message": "Response could not be binded to JSON",
		})
		return
	}
	log.Println("Response unmarshalling successfull, sending response")

	c.JSON(200, result)
	resp.Body.Close()
}

func CreateCollection(c *gin.Context) {
	Collection(c, "create")
}

func DescribeCollection(c *gin.Context) {
	Collection(c, "describe")
}

func DropCollection(c *gin.Context) {
	Collection(c, "drop")
}

func GetLoadState(c *gin.Context) {
	Collection(c, "get_load_state")
}

func GetStats(c *gin.Context) {
	Collection(c, "get_stats")
}

func HasCollection(c *gin.Context) {
	Collection(c, "has")
}

func ListCollections(c *gin.Context) {
	Collection(c, "list")
}

func LoadCollection(c *gin.Context) {
	Collection(c, "load")
}

func ReleaseCollection(c *gin.Context) {
	Collection(c, "release")
}

func RenameCollection(c *gin.Context) {
	Collection(c, "rename")
}
