package service

import (
	"EnGin/internal/db"
	"EnGin/internal/ent"
	"context"
)

func CreateUser(ctx context.Context, username string) (*ent.User, error) {
	// 💡 直接跨包调用 db.Client
	return db.Client.User.Create().
		// SetUsername(username).
		Save(ctx)
}
