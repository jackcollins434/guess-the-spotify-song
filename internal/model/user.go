package model

import "time"

type User struct {
	ID                       uint
	Name                     *string
	Email                    *string
	SpotifyId                string
	SpotifyAccessToken       string
	SpotifyRefreshToken      string
	SpotifyAccessTokenExpiry time.Time
	IsFullyRegistered        bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
