package utils

import "time"

var (
	ShortTokenExpiry   = 1 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
	Suspended          = "suspended"
)
