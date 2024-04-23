package api

import (
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
func (c *JellyfinClient) GetSessions() *[]JellyfinSession {
	request := c.NewRequest("GET", "/Sessions", nil)

	result := DoRequest[[]JellyfinSession](request)
	return result
}

// Get a list of VirtualFolders from the Jellyfin API `/Library/VirtualFolders` endpoint
func (c *JellyfinClient) GetVirtualFolders() *[]JellyfinVirtualFolder {
	request := c.NewRequest("GET", "/Library/VirtualFolders", nil)

	folders := DoRequest[[]JellyfinVirtualFolder](request)

	return folders
}

// Get items belonging to the VirtualFolder with the given Id
func (c *JellyfinClient) GetItems(parentId string) *JellyfinItemsResponse {
	queryValues := url.Values{}
	queryValues.Set("parentId", parentId)
	queryValues.Set("recursive", "true")
	url := "/Items?" + queryValues.Encode()

	request := c.NewRequest("GET", url, nil)

	response := DoRequest[JellyfinItemsResponse](request)

	return response
}

// Get a list of all users from the Jellyfin API
func (c *JellyfinClient) GetUsers() *[]JellyfinUser {
	url := "/Users"
	request := c.NewRequest("GET", url, nil)

	users := DoRequest[[]JellyfinUser](request)

	return users
}
