package api

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

type JellyfinSessionPlayState struct {
	IsPaused   bool   `json:"IsPaused"`
	PlayMethod string `json:"PlayMethod"`
}

type JellyfinSessionNowPlayingItem struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Container string `json:"Container"`
	Type      string `json:"Type"`
	MediaType string `json:"MediaType"`
}
