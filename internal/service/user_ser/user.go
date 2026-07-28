package user_ser

import (
	"EnGin/internal/db"
	"EnGin/internal/ent"
	"EnGin/internal/utils/pwd"
	"context"
	"errors"
)

var ErrHashPwd = errors.New("密码加密失败")

func CreateUser(ctx context.Context, username, password, email, ip string) (*ent.User, error) {
	// 直接跨包调用 db.Client
	hashedPwd, err := pwd.GenerateFromPassword(password)
	if err != nil {
		return nil, ErrHashPwd
	}

	return db.Client.User.Create().
		SetUsername(username).
		SetPassword(hashedPwd).
		SetEmail(email).
		SetIP(ip).
		Save(ctx)
}
