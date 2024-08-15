package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetProductByIDHandler(c *gin.Context) {
	// Extract product ID from route parameters
	productID := c.Param("productId")

	// Validate the product ID format (assuming it's in the form "COMPANY-ID")
	parts := strings.Split(productID, "-")
	if len(parts) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	company := parts[0]
	id := parts[1]

	// Validate the company parameter
	if !contains(COMPANIES, company) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company parameter"})
		return
	}

	// Define the API URL
	apiURL := fmt.Sprintf("http://20.244.56.144/test/companies/%s/products/%s", company, id)

	accessToken := os.Getenv("ACCESS_TOKEN")

	if accessToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Access token not set"})
		return
	}

	// Creating a new HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Add the access token to the request header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Making the API call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make API request"})
		return
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		c.JSON(resp.StatusCode, gin.H{"error": string(body)})
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read API response"})
		return
	}

	// Parse the JSON response
	var product map[string]interface{}
	if err := json.Unmarshal(body, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	// Return the product details as a JSON response
	c.JSON(http.StatusOK, gin.H{"product": product})
}
