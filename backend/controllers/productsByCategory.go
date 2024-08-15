package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

var COMPANIES = []string{"AMZ", "FLP", "SNP", "MYN", "AZO"}
var CATEGORIES = []string{
	"Phone", "Computer", "TV", "Earphone", "Tablet", "Charger", "Mouse",
	"Keypad", "Bluetooth", "Pendrive", "Remote", "Speaker", "Headset",
	"Laptop", "PC",
}

func GetProductsHandler(c *gin.Context) {
	// Extract query parameters
	company := c.Query("company")
	category := c.Query("category")

	// Validate the company parameter
	if !contains(COMPANIES, company) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company parameter"})
		return
	}

	// Validate the category parameter
	if !contains(CATEGORIES, category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category parameter"})
		return
	}

	top, err := strconv.Atoi(c.DefaultQuery("top", "10"))
	if err != nil || top <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'top' parameter"})
		return
	}

	minPrice, err := strconv.Atoi(c.DefaultQuery("minPrice", "0"))
	if err != nil || minPrice < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'minPrice' parameter"})
		return
	}

	maxPrice, err := strconv.Atoi(c.DefaultQuery("maxPrice", "0"))
	if err != nil || maxPrice < minPrice {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'maxPrice' parameter"})
		return
	}
	apiURL := fmt.Sprintf("http://20.244.56.144/test/companies/%s/categories/%s/products?top=%d&minPrice=%d&maxPrice=%d",
		company, category, top, minPrice, maxPrice)

	// Creating a new HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	accessToken := os.Getenv("ACCESS_TOKEN")
	fmt.Println(accessToken)

	// Adding the access token to the request header
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
	var products []map[string]interface{}
	if err := json.Unmarshal(body, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	// Return the products as a JSON response
	c.JSON(http.StatusOK, gin.H{"products": products})
}
