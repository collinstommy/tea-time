package data

import "github.com/pocketbase/dbx"

type Playlist struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	SpotifyId string `db:"spotifyId"`
	User      string `db:"user"`
}

func GetAllPlaylists(db dbx.Builder) ([]Playlist, error) {
	playlists := []Playlist{}

	err := db.
		NewQuery("SELECT id, name, spotifyId, user FROM playlists LIMIT 100").
		All(&playlists)

	return playlists, err
}

// func CreatePlaylist(db dbx.Builder, playlist Playlist) error {
// 	_, err := db.
// 		NewQuery("INSERT INTO playlists (id, name, spotifyId, user) VALUES (:id, :name, :spotifyId, :user)").
// 		Bind(playlist).
// 		Execute()

// 	return err
// }
