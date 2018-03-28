package HotspotSystemGoClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

API_KEY string = "This is an API Key"

func _HTTP_Get(url string) (http.Response) {
		// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	//Set the header with the API Key: http://www.hotspotsystem.com/apidocs/api/reference/#security-apikeyinheader
	req.Header.Set("sn-apikey", API_KEY)

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	return resp
}