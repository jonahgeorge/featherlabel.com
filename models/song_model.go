package models

type SongModel struct {
	Id        *int
	Title     *string
	User      UserModel
	SignedUrl *string
}

// Delete the given song
func (s SongModel) Delete() error {
	return nil
}
