package service

import (
	"context"
	"ent-scaffold/internal/db"
	"ent-scaffold/internal/ent"
)

func CreateUser(ctx context.Context, username string) (*ent.User, error) {
	// 💡 直接跨包调用 db.Client
	return db.Client.User.Create().
		// SetUsername(username).
		Save(ctx)
}
