package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const API_KEY = "b1c742f6ac364fb8a9ea78f888365294"
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"
	const POLLING_URL = TRANSCRIPT_URL + "/" + "e88b92e9-c110-4c3d-8546-24c731c29855"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", POLLING_URL, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	if result["status"] == "completed" {
		fmt.Println(result["text"])
	}
}
