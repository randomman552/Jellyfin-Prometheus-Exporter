package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Jellyfin API client
type JellyfinClient struct {
	Url   string
	Token string
}

// Create a new Jellyfin API client
func NewJellyfinClient(url string, token string) *JellyfinClient {
	return &JellyfinClient{
		Url:   url,
		Token: token,
	}
}

// Create a new request for the Jellyfin API
func (c *JellyfinClient) NewRequest(method string, url string, body io.Reader) *http.Request {
	url = c.Url + url
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Println(err)
	}

	token := "MediaBrowser Token=" + c.Token
	request.Header.Set("Authorization", token)

	return request
}

// Get sessions from the Jellyfin API `/Sessions` endpoint
func (c *JellyfinClient) GetSessions() *[]JellyfinSession {
	client := &http.Client{}
	request := c.NewRequest("GET", "/Sessions", nil)

	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
	}

	// Decode response
	var sessions *[]JellyfinSession
	buf := bytes.Buffer{}
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), &sessions)

	if err != nil {
		panic(err)
	}

	return sessions
}