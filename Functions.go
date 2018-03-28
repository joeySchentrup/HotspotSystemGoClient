package HotspotSystemGoClient

import (
	"os"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url" 
)

API_KEY string = os.Getenv("HOTSPOT_SYSTEM_API_KEY")

//Info on paramters: http://www.hotspotsystem.com/apidocs/api/reference/#operation-getcustomers
func Customers(feilds string, sort string, limit string, offset string) ArrayOfCustomers {
	
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safeFeilds := url.QueryEscape(feilds)
	safeSort := url.QueryEscape(sort)
	safeLimit := url.QueryEscape(limit)
	safeOffset := url.QueryEscape(offset)

	url := fmt.Sprintf("https://api.hotspotsystem.com/v2.0/customers?fields=%s&sort=%s&limit=%s&offset=%s", safeFeilds, safeSort, safeLimit, safeOffset)

	resp := _HTTP_Get(url)

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record ArrayOfCustomers

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}