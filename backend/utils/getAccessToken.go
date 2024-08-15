package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken() string {
	tokenURL := "http://20.244.56.144/test/auth"
	var client1 = os.Getenv("CLIENT_ID")
	fmt.Println(client1, "this sis client")
	// Define the request payload
	payload := map[string]string{
		"companyName":  "Affordmed",
		"clientID":     os.Getenv("CLIENT_ID"),
		"clientSecret": os.Getenv("CLIENT_SECRET"),
		"ownerName":    "Katari Lokeswara Rao",
		"ownerEmail":   "21341A0571@gmrit.edu.in",
		"rollNo":       "21341A0571",
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	fmt.Println("Response:", string(body))
	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}
	fmt.Println("Access Token:", tokenResponse.AccessToken, "Expires in")
	// Return the access token
	return tokenResponse.AccessToken
}
