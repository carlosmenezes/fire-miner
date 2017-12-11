package api

import (
	"encoding/json"
	"fire-miner/app/models"
	"log"
	"net/http"
)

func Status(uri string) (ewbfResult models.EwbfResult, err error) {

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&ewbfResult); err != nil {
		log.Fatal(err)
	}

	return
}
