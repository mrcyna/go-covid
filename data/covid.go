package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const dateLayout = "2006-01-02"

type Response []struct {
	Country     string `json:"Country"`
	CountryCode string `json:"CountryCode"`
	Province    string `json:"Province"`
	City        string `json:"City"`
	CityCode    string `json:"CityCode"`
	Lat         string `json:"Lat"`
	Lon         string `json:"Lon"`
	Cases       int    `json:"Cases"`
	Status      string `json:"Status"`
	Date        string `json:"Date"`
}

func CovidData(country string, days int) (map[string]int, error) {

	const fName = "CovidData"

	// Times
	dtTo := time.Now().Format(dateLayout)
	dtFrom := time.Now().AddDate(0, 0, days*-1).Format(dateLayout)

	// Base API URL
	url := fmt.Sprintf("https://api.covid19api.com/country/%s/status/confirmed?from=%s&to=%s", country, dtFrom, dtTo)

	// Make HTTP Request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return map[string]int{}, fmt.Errorf("%v: Unable to connect to api.covid19api.com", fName)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]int{}, fmt.Errorf("%v: Unable to connect to api.covid19api.com", fName)
	}

	// Parse The Response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return map[string]int{}, fmt.Errorf("%v: Unable to parse the api.covid19api.com response", fName)
	}

	// Make Result Data
	result := make(map[string]int)
	for _, r := range response {
		result[r.Date] = r.Cases
	}

	return result, nil
}
