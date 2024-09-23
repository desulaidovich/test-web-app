package models

type Token struct {
	ID           int    `db:"id"`
	GUID         string `db:"guid"`
	RegreshToken []byte `db:"refresh_token"`
}
