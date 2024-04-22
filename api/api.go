package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
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
func (c *JellyfinClient) GetSessions() []JellyfinSession {
	client := &http.Client{}
	request := c.NewRequest("GET", "/Sessions", nil)

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	// Decode response
	var sessions *[]JellyfinSession
	buf := bytes.Buffer{}
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), &sessions)

	if err != nil {
		panic(err)
	}

	return *sessions
}

// Get a list of VirtualFolders from the Jellyfin API `/Library/VirtualFolders` endpoint
func (c *JellyfinClient) GetVirtualFolders() []JellyfinVirtualFolder {
	client := &http.Client{}
	request := c.NewRequest("GET", "/Library/VirtualFolders", nil)

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	var folders *[]JellyfinVirtualFolder
	buf := bytes.Buffer{}
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), &folders)

	if err != nil {
		panic(err)
	}

	return *folders
}

// Get items belonging to the VirtualFolder with the given Id
func (c *JellyfinClient) GetItems(parentId string) JellyfinItemsResponse {
	queryValues := url.Values{}
	queryValues.Set("parentId", parentId)
	queryValues.Set("recursive", "true")
	url := "/Items?" + queryValues.Encode()

	client := &http.Client{}
	request := c.NewRequest("GET", url, nil)

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	var itemsResponse *JellyfinItemsResponse
	buf := bytes.Buffer{}
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), &itemsResponse)

	if err != nil {
		panic(err)
	}

	return *itemsResponse
}

// Get a list of all users from the Jellyfin API
func (c *JellyfinClient) GetUsers() []JellyfinUser {
	url := "/Users"
	client := &http.Client{}
	request := c.NewRequest("GET", url, nil)

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	var users *[]JellyfinUser
	buf := bytes.Buffer{}
	buf.ReadFrom(response.Body)

	err = json.Unmarshal(buf.Bytes(), &users)

	if err != nil {
		panic(err)
	}

	return *users
}
