package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
	"time"
)

func main () {

	uploadUrl := upload("/home/zerocool/speech-to-text-golang-example/testfile2.wav")
	fmt.Println("upload url: ", uploadUrl)
	id := transcribe(uploadUrl)
	fmt.Println("transcript id: ", id)
	result := poll(id)
	
	fmt.Println(result)

}
//you can tell I stole this part
func poll(id string) string {
	const API_KEY = "b1c742f6ac364fb8a9ea78f888365294"
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

	client := &http.Client{}

	for {
		req, _ := http.NewRequest("GET", TRANSCRIPT_URL+"/"+id, nil)
		req.Header.Set("content-type", "application/json")
		req.Header.Set("authorization", API_KEY)

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)

		fmt.Printf("Polling status: %+v\n", result["status"])

		status := result["status"].(string)

		if status == "completed" {
			fmt.Println("Transcription completed.")
			return result["text"].(string)
		} else if status == "error" {
			fmt.Println("Transcription error:", result["error"])
			return "error"
		}

		time.Sleep(3 * time.Second) // Wait before polling again
	}
}

func transcribe(uploadurl string) string {
        //audio url here from upload script
	AUDIO_URL := uploadurl
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
        //fmt.Println(result["id"])
	return result["id"].(string)
}

func upload(filename string) string {
        const API_KEY = "b1c742f6ac364fb8a9ea78f888365294"
        const UPLOAD_URL = "https://api.assemblyai.com/v2/upload"
	//uplaod url is there for the mfing api
        // Load file
        data, err := ioutil.ReadFile(filename)
        if err != nil {
                log.Fatalln(err)
        }
        //fmt.Println("here")
        // Setup HTTP client and set header
        client := &http.Client{}
        req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))
        req.Header.Set("authorization", API_KEY)
        res, err := client.Do(req)

        if err != nil {
                log.Fatalln(err)
        }

        // decode json and store it in a map
        var result map[string]interface{}
        json.NewDecoder(res.Body).Decode(&result)

        // return the upload_url
        //fmt.Println(result["upload_url"])
	return result["upload_url"].(string)
}
