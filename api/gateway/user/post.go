package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var reginf Reginf
	if err := c.ShouldBindJSON(&reginf); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}
	jsoninf, err := json.Marshal(reginf)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	url := "http://localhost:8001/register"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsoninf))
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

func Login(c *gin.Context) {
	var loginf Loginf
	if err := c.ShouldBindJSON(&loginf); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}
	jsoninf, err := json.Marshal(loginf)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	url := "http://localhost:8001/login"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsoninf))
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

func Logout(c *gin.Context) {
	url := "http://localhost:8001/logout"
	res, err := http.Post(url, "application/json", nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "logout succesful"})
}
