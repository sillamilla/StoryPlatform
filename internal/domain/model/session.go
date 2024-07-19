package model

import "time"

type SessionInfo struct {
	SessionID string
	UserID    string
	CreatedAt time.Time
}
