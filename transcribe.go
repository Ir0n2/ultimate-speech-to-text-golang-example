package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	//audio url here from upload script
	const AUDIO_URL = "https://cdn.assemblyai.com/upload/69531591-8fc9-4733-adf6-28f6e7ed96b7"
	const API_KEY = "b1c742f6ac364fb8a9ea78f888365294"
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

	// prepare json data
	values := map[string]string{"audio_url": AUDIO_URL}
	jsonData, _ := json.Marshal(values)

	// setup HTTP client and set header
	client := &http.Client{}
	req, _ := http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)
	res, _ := client.Do(req)

	defer res.Body.Close()

	// decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// print the id of the transcribed audio
	fmt.Println(result["id"])
}
