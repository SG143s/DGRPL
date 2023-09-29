package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAu(c *gin.Context) {
	url := "http://localhost:8001/getauinformation"

	res, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var responseBody interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		c.JSON(res.StatusCode, gin.H{"error": string(body)})
		return
	}

	// Send the parsed JSON response
	c.JSON(res.StatusCode, responseBody)
}
