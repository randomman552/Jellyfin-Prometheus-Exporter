package api

// Type representing a session from Jellyfin's API
type JellyfinSession struct {
	Id                    string   `json:"Id"`
	UserId                string   `json:"UserId"`
	UserName              string   `json:"UserName"`
	Client                string   `json:"Client"`
	LastActivityDate      string   `json:"LastActivityDate"`
	LastPlaybackCheckIn   string   `json:"LastPlaybackCheckIn"`
	DeviceName            string   `json:"DeviceName"`
	ApplicationVersion    string   `json:"ApplicationVersion"`
	IsActive              bool     `json:"IsActive"`
	SupportsMediaControl  bool     `json:"SupportsMediaControl"`
	SupportsRemoteControl bool     `json:"SupportsRemoteControl"`
	HasCustomDeviceName   bool     `json:"HasCustomDeviceName"`
	PlayableMediaTpes     []string `json:"PlayableMediaTpes"`

	PlayState      *JellyfinSessionPlayState      `json:"PlayState"`
	NowPlayingItem *JellyfinSessionNowPlayingItem `json:"NowPlayingItem"`
}

// Type representing a sessions play state from Jellyfin's API
type JellyfinSessionPlayState struct {
	IsPaused   bool   `json:"IsPaused"`
	PlayMethod string `json:"PlayMethod"`
}

// Type representing a Now Playing item for a Jellyfin session
type JellyfinSessionNowPlayingItem struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Container string `json:"Container"`
	Type      string `json:"Type"`
	MediaType string `json:"MediaType"`
}

// Type representing a Virtual Folder returned from the Jellyfin API
type JellyfinVirtualFolder struct {
	Name           string `json:"Name"`
	CollectionType string `json:"CollectionType"`
	ItemId         string `json:"ItemId"`
}

// Type representing a response from the Jellyfin items API endpoint
type JellyfinItemsResponse struct {
	TotalRecordCount uint64         `json:"TotalRecordCount"`
	StartIndex       uint64         `json:"StartIndex"`
	Items            []JellyfinItem `json:"Items"`
}

// Type representing an item from the Jellyfin items API endpoint
type JellyfinItem struct {
	Name      string `json:"Name"`
	Container string `json:"Container"`
	Type      string `json:"Type"`
}

// A user returned from Jellyfins API
type JellyfinUser struct {
	Name   string `json:"Name"`
	Policy struct {
		IsAdministrator        bool   `json:"IsAdministrator"`
		IsDisabled             bool   `json:"IsDisabled"`
		AuthenticationProvider string `json:"AuthenticationProviderId"`
	} `json:"Policy"`
}
