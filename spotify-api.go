package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Token struct {
	Token   string `json:"access_token"`
	Type    string `json:"token_type"`
	Expires int    `json:"expires_in"`
	Time    int
}

type Spotify struct {
	client_id     string
	client_secret string
	AccessToken   Token
}

type Artists struct {
	External  map[string]string `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Genres []string `json:"genres"`
	Href   string   `json:"href"`
	ID     string   `jsos:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Type       string `json:"type"`
	URL        string `json:"url"`
}

type User struct {
	DisplayName  string            `json:"display_name"`
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `jsos:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type AlbumItem struct {
	Artists          []Artists         `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMS       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       struct {
		ExternalUrls map[string]string `json:"external_urls"`
		Href         string            `json:"href"`
		ID           string            `json:"id"`
		Type         string            `json:"type"`
		URL          string            `json:"url"`
	} `json:"linked_from"`
	Restrictions struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`
	Name        string `json:"name"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	IsLocal     bool   `json:"is_local"`
}

type Album struct {
	AlbumType        string            `json:"album_type"`
	TotalTracks      int               `json:"total_tracks"`
	AvailableMarkets []string          `json:"available_markets"`
	ExternalUrls     map[string]string `json:"external_urls"`
	Href             string            `json:"href"`
	ID               string            `jsos:"id"`
	Images           []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	Restrictions         struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`
	Type    string    `json:"type"`
	URL     string    `json:"url"`
	Artists []Artists `json:"artists"`
	Tracks  struct {
		Href     string      `json:"href"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous string      `json:"previous"`
		Total    int         `json:"total"`
		Items    []AlbumItem `json:"items"`
	} `json:"tracks"`
	Copyrights []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"copyrights"`
	ExternalIDs struct {
		ISRC string `json:"isrc"`
		EAN  string `json:"ean"`
		UPC  string `json:"upc"`
	} `json:"external_ids"`
	Genres     []string `json:"genres"`
	Label      string   `json:"label"`
	Popularity int      `json:"popularity"`
}

type Track struct {
	Album struct {
		AlbumType        string            `json:"album_type"`
		TotalTracks      int               `json:"total_tracks"`
		AvailableMarkets []string          `json:"available_markets"`
		External         map[string]string `json:"external_urls"`
		Href             string            `json:"href"`
		ID               string            `json:"id"`
		Images           []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		Restrictions         struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Type    string    `json:"type"`
		URL     string    `json:"url"`
		Artists []Artists `json:"artists"`
	} `json:"album"`
	Artists []Artists `json:"artists"`
}

type PlaylistItem struct {
	AddedAt string `json:"added_at"`
	AddedBy User   `json:"added_by"`
	IsLocal bool   `json:"is_local"`
	Track   Track  `json:"track"`
}

type Playlist struct {
	Collaborative bool              `json:"collaborative"`
	Description   string            `json:"description"`
	External      map[string]string `json:"external_urls"`
	Followers     struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name       string `json:"name"`
	Owner      User   `json:"owner"`
	Public     bool   `json:"public"`
	SnapshotID string `json:"snapshot_id"`
	Tracks     struct {
		Href     string         `json:"href"`
		Limit    int            `json:"limit"`
		Next     string         `json:"next"`
		Offset   int            `json:"offset"`
		Previous string         `json:"previous"`
		Total    int            `json:"total"`
		Items    []PlaylistItem `json:"items"`
	} `json:"tracks"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type SearchTrack struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		Artists          []Artists         `json:"artists"`
		AvailableMarkets []string          `json:"available_markets"`
		DiscNumber       int               `json:"disc_number"`
		DurationMS       int               `json:"duration_ms"`
		Explicit         bool              `json:"explicit"`
		ExternalUrls     map[string]string `json:"external_urls"`
		Href             string            `json:"href"`
		ID               string            `json:"id"`
		IsPlayable       bool              `json:"is_playable"`
		LinkedFrom       struct {
			ExternalUrls map[string]string `json:"external_urls"`
			Href         string            `json:"href"`
			ID           string            `json:"id"`
			Type         string            `json:"type"`
			URL          string            `json:"url"`
		} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Name        string `json:"name"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URL         string `json:"url"`
		IsLocal     bool   `json:"is_local"`
	} `json:"items"`
}

type SearchArtist struct {
	Href     string    `json:"href"`
	Limit    int       `json:"limit"`
	Next     string    `json:"next"`
	Offset   int       `json:"offset"`
	Previous string    `json:"previous"`
	Total    int       `json:"total"`
	Items    []Artists `json:"items"`
}

type SearchAlbum struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
	Items    []struct {
		AlbumType        string            `json:"album_type"`
		TotalTracks      int               `json:"total_tracks"`
		AvailableMarkets []string          `json:"available_markets"`
		ExternalUrls     map[string]string `json:"external_urls"`
		Href             string            `json:"href"`
		ID               string            `json:"id"`
		Images           []struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		Restrictions         struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Type    string    `json:"type"`
		URL     string    `json:"url"`
		Artists []Artists `json:"artists"`
	} `json:"items"`
}

type SearchPlaylist struct {
	Href     string     `json:"href"`
	Limit    int        `json:"limit"`
	Next     string     `json:"next"`
	Offset   int        `json:"offset"`
	Previous string     `json:"previous"`
	Total    int        `json:"total"`
	Items    []Playlist `json:"items"`
}

type Search struct {
	Tracks   SearchTrack    `json:"tracks"`
	Artists  SearchArtist   `json:"artists"`
	Album    SearchAlbum    `json:"albums"`
	Playlist SearchPlaylist `json:"playlists"`
}

func (al *Album) GetItems() []AlbumItem {
	return al.Tracks.Items
}

func (pl *Playlist) GetItems() []PlaylistItem {
	return pl.Tracks.Items
}

func New(id, secret string) *Spotify {
	return &Spotify{
		client_id:     id,
		client_secret: secret,
		AccessToken:   Token{},
	}
}

func (st *Spotify) valid() error {
	if st.AccessToken.Token == "" {
		if err := st.NewToken(); err != nil {
			return err
		}
	}
	if int(time.Now().Unix())-st.AccessToken.Time >= st.AccessToken.Expires {
		if err := st.NewToken(); err != nil {
			return err
		}
	}
	return nil
}

func (st *Spotify) authGET(url string) ([]byte, error) {
	if err := st.valid(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", st.AccessToken.Type, st.AccessToken.Token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return body, nil
}

func (st *Spotify) NewToken() error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", st.client_id)
	data.Set("client_secret", st.client_secret)
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return err
	}
	token.Time = int(time.Now().Unix())
	st.AccessToken = token
	return nil
}

func (st *Spotify) GetArtist(id string) (*Artists, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/artists/%s", id))
	if err != nil {
		return nil, err
	}
	var info Artists
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (st *Spotify) GetUser(id string) (*User, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/users/%s", id))
	if err != nil {
		return nil, err
	}
	var info User
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (st *Spotify) GetAlbum(id string) (*Album, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/albums/%s", id))
	if err != nil {
		return nil, err
	}
	var info Album
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (st *Spotify) GetTrack(id string) (*Track, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", id))
	if err != nil {
		return nil, err
	}
	var info Track
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (st *Spotify) GetPlaylist(id string) (*Playlist, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", id))
	if err != nil {
		return nil, err
	}
	var info Playlist
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (st *Spotify) SearchTrack(q string, limit int) (*SearchTrack, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=%d", strings.ReplaceAll(q, " ", ""), limit))
	if err != nil {
		return nil, err
	}
	var info Search
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info.Tracks, nil
}

func (st *Spotify) SearchArtist(q string, limit int) (*SearchArtist, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist&limit=%d", strings.ReplaceAll(q, " ", ""), limit))
	if err != nil {
		return nil, err
	}
	var info Search
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info.Artists, nil
}

func (st *Spotify) SearchAlbum(q string, limit int) (*SearchAlbum, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=album&limit=%d", strings.ReplaceAll(q, " ", ""), limit))
	if err != nil {
		return nil, err
	}
	var info Search
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info.Album, nil
}

func (st *Spotify) SearchPlaylist(q string, limit int) (*SearchPlaylist, error) {
	body, err := st.authGET(fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=playlist&limit=%d", strings.ReplaceAll(q, " ", ""), limit))
	if err != nil {
		return nil, err
	}
	var info Search
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}
	return &info.Playlist, nil
}
