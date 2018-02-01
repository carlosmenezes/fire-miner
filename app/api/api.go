package api

import (
	"encoding/json"
	"fire-miner/app/models"
	"log"
	"net/http"
)

func EwbfStatus(uri string) (ewbfResult models.EwbfResult, err error) {

	resp, err := makeRequest(uri)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&ewbfResult); err != nil {
		log.Println(err)
	}

	return ewbfResult, nil
}

func BMinerStatus(uri string) (bminerResult models.BMinerResult, err error) {
	resp, err := makeRequest(uri)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&bminerResult); err != nil {
		log.Println(err)
	}

	return bminerResult, nil
}

func makeRequest(uri string) (response *http.Response, err error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	response, err = client.Do(req)
	if err != nil {
		log.Println("Do: ", err)
		return
	}

	return response, nil
}
