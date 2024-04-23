package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Function to run the given request and parse the JSON response into the given type
func DoRequest[T any](request *http.Request) *T {
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		return nil
	}

	var result *T

	buf := bytes.Buffer{}
	buf.ReadFrom(response.Body)

	// For successful responses
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		err = json.Unmarshal(buf.Bytes(), &result)

		if err != nil {
			log.Println(err)
			return nil
		}

		return result
	} else {
		logString := "Unsuccessful status code '" + strconv.FormatInt(int64(response.StatusCode), 10) + "' when querying '" + request.URL.String() + "'"

		log.Println(logString)
	}

	return nil
}
