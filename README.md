<div align="center">

 [![spotify](/docs/spotify.png)](https://developer.spotify.com/documentation/web-api)

  This project is a module that makes it easier to use the **Spotify API**.  
  You can get the API key from [here](https://developer.spotify.com/).
</div>

---
### ⚙️ **API Setup**
> When the **CLIENT_ID** and **SECRET_KEY** are correct, the **token will be automatically issued**.
```go
package main

import "github.com/fluffy-melli/spotify-api"

func main() {
    var CLIENT_ID = "" // It is recommended to use a ".env" file.
    var SECRET_KEY = "" // It is recommended to use a ".env" file.

    api := spotify.New(CLIENT_ID, SECRET_KEY)
}
```
---
### 🎵 Track (Music)

```go
Trackinfo, err := api.GetTrack("Track ID")
if err != nil {
    return
}
fmt.Println(Trackinfo.Album)
```

```go
SearchTrackinfo, err := api.SearchTrack("Search for track content", 5)
if err != nil {
    return
}
fmt.Println(SearchTrackinfo.Items)
```

---
### 🎶 Album

```go
Albuminfo, err := api.GetAlbum("Album ID")
if err != nil {
    return
}
fmt.Println(Albuminfo.GetItems())
```

```go
SearchAlbuminfo, err := api.SearchAlbum("Search for Album Name", 5)
if err != nil {
	return
}
fmt.Println(SearchAlbuminfo)
```

---
### 👤 User

```go
Userinfo, err := api.GetUser("User ID")
if err != nil {
    return
}
fmt.Println(Userinfo)
```

---
### 🎤 Artist

```go
Artistinfo, err := api.GetArtist("Artist ID")
if err != nil {
	return
}
fmt.Println(Artistinfo)
```

```go
SearchArtist, err := api.SearchArtist("Search for Artist Name", 5)
if err != nil {
	fmt.Println(err)
}
fmt.Println(SearchArtist.Items)
```

---
### 🎧 PlayList

```go
Playlistinfo, err := api.GetPlaylist("PlayList ID")
if err != nil {
	return
}
fmt.Println(Playlist.GetItems())
```

```go
SearchPlaylist, err := api.SearchPlaylist("Search for Playlist Name", 5)
if err != nil {
	fmt.Println(err)
}
fmt.Println(SearchPlaylist.Items)
```

---

### 📜 **License**

The `spotify-api` follows the **MIT License** by default, but please also be aware of **Spotify's License** when using it.


Copyright © All rights reserved.
