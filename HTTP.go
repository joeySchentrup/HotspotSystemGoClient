package HotspotSystemGoClient

import (
	"os"
	"log"
	"net/http"
)

func _HTTP_Get(url string) (http.Response) {
		
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("_HTTP_Get NewRequest: ", err)
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	//Set the header with the API Key: http://www.hotspotsystem.com/apidocs/api/reference/#security-apikeyinheader
	req.Header.Set("sn-apikey", os.Getenv("HOTSPOT_SYSTEM_API_KEY"))

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("_HTTP_Get Do: ", err)
	}

	return *resp
}